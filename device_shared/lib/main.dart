import 'package:flutter/material.dart';
import 'package:device_shared/init_ads.dart';
import 'package:device_shared/screens/login_screen.dart';
import 'package:device_shared/screens/home_screen.dart';
import 'package:device_shared/services/auth_service.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initAds();
  runApp(const TodoApp());
}

class TodoApp extends StatelessWidget {
  const TodoApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Todo App',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
        useMaterial3: true,
      ),
      home: const AuthWrapper(),
    );
  }
}

class AuthWrapper extends StatefulWidget {
  const AuthWrapper({super.key});

  @override
  State<AuthWrapper> createState() => _AuthWrapperState();
}

class _AuthWrapperState extends State<AuthWrapper> {
  final AuthService _auth = AuthService();
  bool _loading = true;
  bool _authenticated = false;

  @override
  void initState() {
    super.initState();
    _checkAuth();
  }

  Future<void> _checkAuth() async {
    final token = await _auth.getToken();
    setState(() {
      _authenticated = token != null && token.isNotEmpty;
      _loading = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_loading) {
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }
    if (_authenticated) {
      return HomeScreen(onLogout: () async {
        await _auth.logout();
        setState(() => _authenticated = false);
      });
    }
    return LoginScreen(onLogin: () async {
      await _checkAuth();
      setState(() => _authenticated = true);
    });
  }
}
