import 'package:flutter/material.dart';
import 'package:device_shared/services/api_service.dart';
import 'package:device_shared/widgets/task_tile.dart';

class TodayScreen extends StatefulWidget {
  const TodayScreen({super.key});

  @override
  State<TodayScreen> createState() => _TodayScreenState();
}

class _TodayScreenState extends State<TodayScreen> {
  final ApiService _api = ApiService();
  List<Task> _tasks = [];
  bool _loading = true;
  String? _error;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() {
      _loading = true;
      _error = null;
    });
    try {
      final tasks = await _api.getTodayTasks();
      setState(() {
        _tasks = tasks;
        _loading = false;
      });
    } catch (e) {
      setState(() {
        _error = e.toString();
        _loading = false;
      });
    }
  }

  Future<void> _addTask(String title) async {
    try {
      final today = DateTime.now();
      final due = DateTime(today.year, today.month, today.day, 23, 59, 59);
      await _api.createTask(title: title, dueDate: due.toIso8601String());
      _load();
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_loading) {
      return const Center(child: CircularProgressIndicator());
    }
    if (_error != null) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(_error!, textAlign: TextAlign.center),
            ElevatedButton(onPressed: _load, child: const Text('重試')),
          ],
        ),
      );
    }
    return RefreshIndicator(
      onRefresh: _load,
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          Text(
            'Today - ${DateTime.now().day}/${DateTime.now().month}/${DateTime.now().year}',
            style: Theme.of(context).textTheme.titleLarge,
          ),
          const SizedBox(height: 16),
          OutlinedButton.icon(
            onPressed: () => _showAddDialog(context),
            icon: const Icon(Icons.add),
            label: const Text('新增任務'),
          ),
          const SizedBox(height: 16),
          ..._tasks.map((t) => TaskTile(
                task: t,
                onToggle: () async {
                  await _api.updateTask(
                    t.id,
                    {'status': t.isCompleted ? 'pending' : 'completed'},
                  );
                  _load();
                },
                onDelete: () async {
                  await _api.deleteTask(t.id);
                  _load();
                },
              )),
        ],
      ),
    );
  }

  void _showAddDialog(BuildContext context) {
    final controller = TextEditingController();
    showDialog(
      context: context,
      builder: (ctx) => AlertDialog(
        title: const Text('新增任務'),
        content: TextField(
          controller: controller,
          decoration: const InputDecoration(hintText: '任務標題'),
          autofocus: true,
        ),
        actions: [
          TextButton(onPressed: () => Navigator.pop(ctx), child: const Text('取消')),
          ElevatedButton(
            onPressed: () async {
              if (controller.text.trim().isEmpty) return;
              Navigator.pop(ctx);
              await _addTask(controller.text.trim());
            },
            child: const Text('新增'),
          ),
        ],
      ),
    );
  }
}
