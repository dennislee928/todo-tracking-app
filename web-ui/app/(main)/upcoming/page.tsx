"use client";

import { useEffect, useState } from "react";
import { api } from "@/lib/api";
import type { Task } from "@/lib/api";
import { TaskItem } from "@/components/task-item";

export default function UpcomingPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchTasks = async () => {
    try {
      const data = await api.tasks.upcoming(14);
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

  // Group by date
  const byDate = tasks.reduce<Record<string, Task[]>>((acc, task) => {
    const key = task.due_date
      ? new Date(task.due_date).toDateString()
      : "無日期";
    if (!acc[key]) acc[key] = [];
    acc[key].push(task);
    return acc;
  }, {});

  const sortedDates = Object.keys(byDate).sort(
    (a, b) =>
      (a === "無日期" ? 0 : new Date(a).getTime()) -
      (b === "無日期" ? 0 : new Date(b).getTime())
  );

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-2xl font-bold">Upcoming</h1>
        <p className="text-muted-foreground">未來 14 天的任務</p>
      </div>

      {tasks.length === 0 ? (
        <p className="text-muted-foreground text-center py-12">
          沒有即將到來的任務
        </p>
      ) : (
        sortedDates.map((date) => (
          <div key={date}>
            <h2 className="text-lg font-semibold mb-3">
              {date === "無日期"
                ? "無日期"
                : new Date(date).toLocaleDateString("zh-TW", {
                    weekday: "long",
                    month: "long",
                    day: "numeric",
                  })}
            </h2>
            <div className="space-y-2">
              {byDate[date].map((task) => (
                <TaskItem
                  key={task.id}
                  task={task}
                  onUpdate={handleUpdate}
                  onDelete={handleDelete}
                />
              ))}
            </div>
          </div>
        ))
      )}
    </div>
  );
}
