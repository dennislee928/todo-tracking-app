import 'package:flutter/material.dart';
import 'package:device_shared/services/api_service.dart';

class TaskTile extends StatelessWidget {
  final Task task;
  final VoidCallback onToggle;
  final VoidCallback onDelete;

  const TaskTile({
    super.key,
    required this.task,
    required this.onToggle,
    required this.onDelete,
  });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: Checkbox(
        value: task.isCompleted,
        onChanged: (_) => onToggle(),
      ),
      title: Text(
        task.title,
        style: task.isCompleted
            ? const TextStyle(decoration: TextDecoration.lineThrough)
            : null,
      ),
      subtitle: task.dueDate != null
          ? Text(
              DateTime.parse(task.dueDate!).toString().substring(0, 10),
              style: const TextStyle(fontSize: 12),
            )
          : null,
      trailing: IconButton(
        icon: const Icon(Icons.delete),
        onPressed: () async {
          final ok = await showDialog<bool>(
            context: context,
            builder: (ctx) => AlertDialog(
              title: const Text('刪除任務'),
              content: const Text('確定要刪除此任務？'),
              actions: [
                TextButton(
                  onPressed: () => Navigator.pop(ctx, false),
                  child: const Text('取消'),
                ),
                TextButton(
                  onPressed: () => Navigator.pop(ctx, true),
                  child: const Text('刪除'),
                ),
              ],
            ),
          );
          if (ok == true) onDelete();
        },
      ),
    );
  }
}
