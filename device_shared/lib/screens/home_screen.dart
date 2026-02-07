import 'package:flutter/material.dart';
import 'package:device_shared/screens/today_screen.dart';
import 'package:device_shared/screens/upcoming_screen.dart';
import 'package:device_shared/screens/projects_screen.dart';

class HomeScreen extends StatelessWidget {
  final VoidCallback onLogout;

  const HomeScreen({super.key, required this.onLogout});

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
            IconButton(
              icon: const Icon(Icons.logout),
              onPressed: onLogout,
            ),
          ],
        ),
        body: const TabBarView(
          children: [
            TodayScreen(),
            UpcomingScreen(),
            ProjectsScreen(),
          ],
        ),
      ),
    );
  }
}
