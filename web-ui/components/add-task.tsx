"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import type { Task } from "@/lib/api";

interface AddTaskProps {
  projectId?: string;
  onAdd: (data: Partial<Task>) => Promise<void>;
}

export function AddTask({ projectId, onAdd }: AddTaskProps) {
  const [title, setTitle] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;
    setLoading(true);
    try {
      await onAdd({
        title: title.trim(),
        project_id: projectId,
      });
      setTitle("");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex gap-2">
      <Input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="新增任務..."
        disabled={loading}
      />
      <Button type="submit" disabled={loading || !title.trim()}>
        新增
      </Button>
    </form>
  );
}
