import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:device_shared/services/auth_service.dart';

const _baseUrl = String.fromEnvironment(
  'API_URL',
  defaultValue: 'http://localhost:8080/api/v1',
);

class ApiService {
  final AuthService _auth = AuthService();

  Future<Map<String, String>> _headers() async {
    final token = await _auth.getToken();
    return {
      'Content-Type': 'application/json',
      if (token != null) 'Authorization': 'Bearer $token',
    };
  }

  Future<List<Task>> getTodayTasks() async {
    final res = await http.get(
      Uri.parse('$_baseUrl/tasks/today'),
      headers: await _headers(),
    );
    if (res.statusCode != 200) throw Exception(res.body);
    final list = jsonDecode(res.body) as List;
    return list.map((e) => Task.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<List<Task>> getUpcomingTasks({int days = 7}) async {
    final res = await http.get(
      Uri.parse('$_baseUrl/tasks/upcoming?days=$days'),
      headers: await _headers(),
    );
    if (res.statusCode != 200) throw Exception(res.body);
    final list = jsonDecode(res.body) as List;
    return list.map((e) => Task.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<List<Task>> getTasks({String? projectId}) async {
    final uri = projectId != null
        ? Uri.parse('$_baseUrl/tasks?project_id=$projectId')
        : Uri.parse('$_baseUrl/tasks');
    final res = await http.get(uri, headers: await _headers());
    if (res.statusCode != 200) throw Exception(res.body);
    final list = jsonDecode(res.body) as List;
    return list.map((e) => Task.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<Task> createTask({
    required String title,
    String? projectId,
    String? dueDate,
  }) async {
    final body = <String, dynamic>{'title': title};
    if (projectId != null) body['project_id'] = projectId;
    if (dueDate != null) body['due_date'] = dueDate;
    final res = await http.post(
      Uri.parse('$_baseUrl/tasks'),
      headers: await _headers(),
      body: jsonEncode(body),
    );
    if (res.statusCode != 201) throw Exception(res.body);
    return Task.fromJson(jsonDecode(res.body) as Map<String, dynamic>);
  }

  Future<Task> updateTask(String id, Map<String, dynamic> data) async {
    final res = await http.put(
      Uri.parse('$_baseUrl/tasks/$id'),
      headers: await _headers(),
      body: jsonEncode(data),
    );
    if (res.statusCode != 200) throw Exception(res.body);
    return Task.fromJson(jsonDecode(res.body) as Map<String, dynamic>);
  }

  Future<void> deleteTask(String id) async {
    final res = await http.delete(
      Uri.parse('$_baseUrl/tasks/$id'),
      headers: await _headers(),
    );
    if (res.statusCode != 204) throw Exception(res.body);
  }

  Future<List<Project>> getProjects() async {
    final res = await http.get(
      Uri.parse('$_baseUrl/projects'),
      headers: await _headers(),
    );
    if (res.statusCode != 200) throw Exception(res.body);
    final list = jsonDecode(res.body) as List;
    return list.map((e) => Project.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<Project> createProject(String name) async {
    final res = await http.post(
      Uri.parse('$_baseUrl/projects'),
      headers: await _headers(),
      body: jsonEncode({'name': name}),
    );
    if (res.statusCode != 201) throw Exception(res.body);
    return Project.fromJson(jsonDecode(res.body) as Map<String, dynamic>);
  }

  Future<User> getMe() async {
    final res = await http.get(
      Uri.parse('$_baseUrl/me'),
      headers: await _headers(),
    );
    if (res.statusCode != 200) throw Exception(res.body);
    return User.fromJson(jsonDecode(res.body) as Map<String, dynamic>);
  }

  Future<void> verifyApplePurchase(String receiptData) async {
    final res = await http.post(
      Uri.parse('$_baseUrl/subscription/apple-verify'),
      headers: await _headers(),
      body: jsonEncode({'receipt_data': receiptData}),
    );
    if (res.statusCode != 200) throw Exception(res.body);
  }

  Future<void> verifyGooglePurchase(String purchaseToken, String productId) async {
    final res = await http.post(
      Uri.parse('$_baseUrl/subscription/google-verify'),
      headers: await _headers(),
      body: jsonEncode({
        'purchase_token': purchaseToken,
        'product_id': productId,
      }),
    );
    if (res.statusCode != 200) throw Exception(res.body);
  }
}

class User {
  final String id;
  final String email;
  final bool isPremium;

  User({required this.id, required this.email, required this.isPremium});

  factory User.fromJson(Map<String, dynamic> json) => User(
        id: json['id'] as String,
        email: json['email'] as String,
        isPremium: json['is_premium'] as bool? ?? false,
      );
}

class Task {
  final String id;
  final String title;
  final String? description;
  final String? projectId;
  final int priority;
  final String status;
  final String? dueDate;

  Task({
    required this.id,
    required this.title,
    this.description,
    this.projectId,
    this.priority = 0,
    this.status = 'pending',
    this.dueDate,
  });

  factory Task.fromJson(Map<String, dynamic> json) => Task(
        id: json['id'] as String,
        title: json['title'] as String,
        description: json['description'] as String?,
        projectId: json['project_id'] as String?,
        priority: json['priority'] as int? ?? 0,
        status: json['status'] as String? ?? 'pending',
        dueDate: json['due_date'] as String?,
      );

  bool get isCompleted => status == 'completed';
}

class Project {
  final String id;
  final String name;
  final String? color;

  Project({required this.id, required this.name, this.color});

  factory Project.fromJson(Map<String, dynamic> json) => Project(
        id: json['id'] as String,
        name: json['name'] as String,
        color: json['color'] as String?,
      );
}
