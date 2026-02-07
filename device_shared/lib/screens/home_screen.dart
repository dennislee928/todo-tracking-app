import 'dart:io';

import 'package:flutter/material.dart';
import 'package:device_shared/screens/today_screen.dart';
import 'package:device_shared/screens/upcoming_screen.dart';
import 'package:device_shared/screens/projects_screen.dart';
import 'package:device_shared/services/api_service.dart';
import 'package:device_shared/services/subscription_service.dart';
import 'package:device_shared/widgets/ad_banner.dart';

class HomeScreen extends StatefulWidget {
  final VoidCallback onLogout;

  const HomeScreen({super.key, required this.onLogout});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final ApiService _api = ApiService();
  final SubscriptionService _subscription = SubscriptionService();

  bool _isPremium = false;
  bool _loading = true;
  bool _purchasing = false;

  @override
  void initState() {
    super.initState();
    _subscription.onPurchaseVerified = _fetchUser;
    _subscription.init();
    _fetchUser();
  }

  @override
  void dispose() {
    _subscription.dispose();
    super.dispose();
  }

  Future<void> _fetchUser() async {
    try {
      final user = await _api.getMe();
      if (mounted) {
        setState(() {
          _isPremium = user.isPremium;
          _loading = false;
        });
      }
    } catch (_) {
      if (mounted) {
        setState(() {
          _isPremium = false;
          _loading = false;
        });
      }
    }
  }

  Future<void> _handleUpgrade() async {
    if (!Platform.isIOS && !Platform.isAndroid) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('IAP 僅支援 iOS 與 Android')),
      );
      return;
    }

    setState(() => _purchasing = true);
    try {
      final ok = await _subscription.purchaseRemoveAds();
      if (!ok && mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('無法啟動購買，請稍後再試')),
        );
      }
      // Purchase result comes via stream; onPurchaseVerified triggers _fetchUser
    } finally {
      if (mounted) setState(() => _purchasing = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Todo App'),
          bottom: const TabBar(
            tabs: [
              Tab(text: 'Today'),
              Tab(text: 'Upcoming'),
              Tab(text: '專案'),
            ],
          ),
          actions: [
            if (!_loading && !_isPremium)
              TextButton(
                onPressed: _purchasing ? null : _handleUpgrade,
                child: Text(_purchasing ? '處理中...' : '升級去廣告'),
              )
            else if (!_loading && _isPremium)
              const Padding(
                padding: EdgeInsets.symmetric(horizontal: 8),
                child: Text('已升級', style: TextStyle(fontSize: 12)),
              ),
            IconButton(
              icon: const Icon(Icons.logout),
              onPressed: widget.onLogout,
            ),
          ],
        ),
        body: Column(
          children: [
            const Expanded(
              child: TabBarView(
                children: [
                  TodayScreen(),
                  UpcomingScreen(),
                  ProjectsScreen(),
                ],
              ),
            ),
            AdBanner(isPremium: _isPremium),
          ],
        ),
      ),
    );
  }
}
