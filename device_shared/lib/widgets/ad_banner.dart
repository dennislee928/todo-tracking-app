import 'dart:io' show Platform;

import 'package:flutter/material.dart';
import 'package:google_mobile_ads/google_mobile_ads.dart';

/// AdMob banner - show only when user is not premium.
/// Use test IDs for development; replace with real IDs from AdMob console.
class AdBanner extends StatefulWidget {
  final bool isPremium;

  const AdBanner({super.key, required this.isPremium});

  @override
  State<AdBanner> createState() => _AdBannerState();
}

class _AdBannerState extends State<AdBanner> {
  BannerAd? _bannerAd;
  bool _loaded = false;

  // Test IDs - replace with real IDs from AdMob console
  static const _androidBannerId = 'ca-app-pub-3940256099942544/6300978111';
  static const _iosBannerId = 'ca-app-pub-3940256099942544/2934735716';

  static String get _bannerId =>
      Platform.isIOS ? _iosBannerId : _androidBannerId;

  @override
  void initState() {
    super.initState();
    if (!widget.isPremium) {
      _loadAd();
    }
  }

  @override
  void didUpdateWidget(AdBanner oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (widget.isPremium && _bannerAd != null) {
      _bannerAd = null;
      _loaded = false;
    } else if (!widget.isPremium && !_loaded) {
      _loadAd();
    }
  }

  void _loadAd() {
    _bannerAd = BannerAd(
      adUnitId: _bannerId,
      size: AdSize.banner,
      request: const AdRequest(),
      listener: BannerAdListener(
        onAdLoaded: (_) => setState(() => _loaded = true),
        onAdFailedToLoad: (_, err) => debugPrint('Ad load failed: $err'),
      ),
    );
    _bannerAd!.load();
  }

  @override
  void dispose() {
    _bannerAd?.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    if (widget.isPremium) return const SizedBox.shrink();
    if (_bannerAd == null || !_loaded) {
      return const SizedBox(
        height: 50,
        child: Center(
          child: Text('廣告載入中...', style: TextStyle(fontSize: 12)),
        ),
      );
    }
    return SizedBox(
      height: 50,
      child: AdWidget(ad: _bannerAd!),
    );
  }
}
