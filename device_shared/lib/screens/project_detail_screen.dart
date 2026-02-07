import 'package:flutter/material.dart';
import 'package:device_shared/services/api_service.dart';
import 'package:device_shared/widgets/task_tile.dart';

class ProjectDetailScreen extends StatefulWidget {
  final Project project;

  const ProjectDetailScreen({super.key, required this.project});

  @override
  State<ProjectDetailScreen> createState() => _ProjectDetailScreenState();
}

class _ProjectDetailScreenState extends State<ProjectDetailScreen> {
  final ApiService _api = ApiService();
  List<Task> _tasks = [];
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() => _loading = true);
    try {
      final tasks = await _api.getTasks(projectId: widget.project.id);
      setState(() {
        _tasks = tasks;
        _loading = false;
      });
    } catch (e) {
      setState(() => _loading = false);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
      }
    }
  }

  Future<void> _addTask(String title) async {
    try {
      await _api.createTask(title: title, projectId: widget.project.id);
      _load();
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(widget.project.name)),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : RefreshIndicator(
              onRefresh: _load,
              child: ListView(
                padding: const EdgeInsets.all(16),
                children: [
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
