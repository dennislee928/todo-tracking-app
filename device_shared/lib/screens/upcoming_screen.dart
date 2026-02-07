import 'package:flutter/material.dart';
import 'package:device_shared/services/api_service.dart';
import 'package:device_shared/widgets/task_tile.dart';

class UpcomingScreen extends StatefulWidget {
  const UpcomingScreen({super.key});

  @override
  State<UpcomingScreen> createState() => _UpcomingScreenState();
}

class _UpcomingScreenState extends State<UpcomingScreen> {
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
      final tasks = await _api.getUpcomingTasks(days: 14);
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
    final byDate = <String, List<Task>>{};
    for (final t in _tasks) {
      final key = t.dueDate != null
          ? DateTime.parse(t.dueDate!).toIso8601String().substring(0, 10)
          : '無日期';
      byDate.putIfAbsent(key, () => []).add(t);
    }
    final dates = byDate.keys.toList()..sort();
    return RefreshIndicator(
      onRefresh: _load,
      child: ListView.builder(
        padding: const EdgeInsets.all(16),
        itemCount: dates.length,
        itemBuilder: (context, i) {
          final date = dates[i];
          final tasks = byDate[date]!;
          return Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                date == '無日期' ? '無日期' : date,
                style: Theme.of(context).textTheme.titleMedium,
              ),
              const SizedBox(height: 8),
              ...tasks.map((t) => TaskTile(
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
              const SizedBox(height: 16),
            ],
          );
        },
      ),
    );
  }
}
