"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import Link from "next/link";
import { api } from "@/lib/api";
import type { Task, Project } from "@/lib/api";
import { TaskItem } from "@/components/task-item";
import { AddTask } from "@/components/add-task";
import { Button } from "@/components/ui/button";

export default function ProjectDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [project, setProject] = useState<Project | null>(null);
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchData = async () => {
    try {
      const [proj, taskList] = await Promise.all([
        api.projects.get(id),
        api.tasks.list(id),
      ]);
      setProject(proj);
      setTasks(taskList);
    } catch {
      setProject(null);
      setTasks([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [id]);

  const handleAdd = async (data: Partial<Task>) => {
    await api.tasks.create({ ...data, project_id: id });
    fetchData();
  };

  const handleUpdate = async (taskId: string, data: Partial<Task>) => {
    await api.tasks.update(taskId, data);
    fetchData();
  };

  const handleDelete = async (taskId: string) => {
    await api.tasks.delete(taskId);
    fetchData();
  };

  if (loading) {
    return <div className="text-center py-12">載入中...</div>;
  }

  if (!project) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground">找不到專案</p>
        <Button asChild className="mt-4">
          <Link href="/projects">返回專案列表</Link>
        </Button>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Button variant="ghost" size="sm" asChild>
          <Link href="/projects">← 返回</Link>
        </Button>
        <div>
          <h1 className="text-2xl font-bold">{project.name}</h1>
          <p className="text-muted-foreground">{tasks.length} 個任務</p>
        </div>
      </div>

      <AddTask projectId={id} onAdd={handleAdd} />

      <div className="space-y-2">
        {tasks.length === 0 ? (
          <p className="text-muted-foreground text-center py-8">
            此專案尚無任務
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
