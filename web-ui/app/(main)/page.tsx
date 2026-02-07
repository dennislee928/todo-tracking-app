"use client";

import { useEffect, useState } from "react";
import { api } from "@/lib/api";
import type { Task } from "@/lib/api";
import { TaskItem } from "@/components/task-item";
import { AddTask } from "@/components/add-task";

export default function TodayPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchTasks = async () => {
    try {
      const data = await api.tasks.today();
      setTasks(data);
    } catch {
      setTasks([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const handleAdd = async (data: Partial<Task>) => {
    const today = new Date();
    today.setHours(23, 59, 59, 999);
    await api.tasks.create({
      ...data,
      due_date: today.toISOString(),
    });
    fetchTasks();
  };

  const handleUpdate = async (id: string, data: Partial<Task>) => {
    await api.tasks.update(id, data);
    fetchTasks();
  };

  const handleDelete = async (id: string) => {
    await api.tasks.delete(id);
    fetchTasks();
  };

  if (loading) {
    return <div className="text-center py-12">載入中...</div>;
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold">Today</h1>
        <p className="text-muted-foreground">
          {new Date().toLocaleDateString("zh-TW", {
            weekday: "long",
            year: "numeric",
            month: "long",
            day: "numeric",
          })}
        </p>
      </div>

      <AddTask onAdd={handleAdd} />

      <div className="space-y-2">
        {tasks.length === 0 ? (
          <p className="text-muted-foreground text-center py-8">
            今天沒有任務，試試新增一個！
          </p>
        ) : (
          tasks.map((task) => (
            <TaskItem
              key={task.id}
              task={task}
              onUpdate={handleUpdate}
              onDelete={handleDelete}
            />
          ))
        )}
      </div>
    </div>
  );
}
