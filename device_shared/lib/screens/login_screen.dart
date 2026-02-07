import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:device_shared/screens/register_screen.dart';
import 'package:http/http.dart' as http;
import 'package:device_shared/services/auth_service.dart';

// Android emulator: 10.0.2.2:8080 | iOS simulator: localhost:8080 | Real device: your server IP
const _baseUrl = String.fromEnvironment(
  'API_URL',
  defaultValue: 'http://localhost:8080/api/v1',
);

class LoginScreen extends StatelessWidget {
  final VoidCallback onLogin;

  const LoginScreen({super.key, required this.onLogin});

  Future<void> _login(BuildContext context, String email, String password) async {
    try {
      final res = await http.post(
        Uri.parse('$_baseUrl/auth/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'email': email, 'password': password}),
      );
      if (res.statusCode != 200) {
        final err = jsonDecode(res.body);
        if (context.mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text(err['error']?.toString() ?? 'Login failed')),
          );
        }
        return;
      }
      final data = jsonDecode(res.body) as Map<String, dynamic>;
      final token = data['token'] as String;
      await AuthService().setToken(token);
      if (context.mounted) onLogin();
    } catch (e) {
      if (context.mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Error: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final emailController = TextEditingController();
    final passwordController = TextEditingController();
    bool loading = false;

    return Scaffold(
      appBar: AppBar(title: const Text('Todo App - 登入')),
      body: StatefulBuilder(
        builder: (context, setState) {
          return Padding(
            padding: const EdgeInsets.all(24.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                TextField(
                  controller: emailController,
                  decoration: const InputDecoration(labelText: 'Email'),
                  keyboardType: TextInputType.emailAddress,
                ),
                const SizedBox(height: 16),
                TextField(
                  controller: passwordController,
                  decoration: const InputDecoration(labelText: '密碼'),
                  obscureText: true,
                ),
                const SizedBox(height: 24),
                TextButton(
                  onPressed: () => Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (_) => RegisterScreen(onRegistered: onLogin),
                    ),
                  ),
                  child: const Text('還沒有帳號？註冊'),
                ),
                const SizedBox(height: 16),
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: loading
                        ? null
                        : () async {
                            setState(() => loading = true);
                            await _login(
                              context,
                              emailController.text,
                              passwordController.text,
                            );
                            if (context.mounted) {
                              setState(() => loading = false);
                            }
                          },
                    child: Text(loading ? '登入中...' : '登入'),
                  ),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
