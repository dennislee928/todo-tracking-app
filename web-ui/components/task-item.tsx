"use client";

import { useState } from "react";
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import type { Task } from "@/lib/api";

interface TaskItemProps {
  task: Task;
  onUpdate: (id: string, data: Partial<Task>) => Promise<void>;
  onDelete: (id: string) => Promise<void>;
}

export function TaskItem({ task, onUpdate, onDelete }: TaskItemProps) {
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState(task.title);
  const [saving, setSaving] = useState(false);

  const isCompleted = task.status === "completed";

  const handleToggle = async () => {
    const newStatus = isCompleted ? "pending" : "completed";
    await onUpdate(task.id, { status: newStatus });
  };

  const handleSave = async () => {
    setSaving(true);
    try {
      await onUpdate(task.id, { title });
      setOpen(false);
    } finally {
      setSaving(false);
    }
  };

  const handleDelete = async () => {
    if (confirm("確定要刪除此任務？")) {
      await onDelete(task.id);
      setOpen(false);
    }
  };

  const dueDate = task.due_date
    ? new Date(task.due_date).toLocaleDateString("zh-TW", {
        month: "short",
        day: "numeric",
      })
    : null;

  return (
    <>
      <div
        className="flex items-center gap-3 p-3 rounded-lg border hover:bg-muted/50 cursor-pointer group"
        onClick={() => setOpen(true)}
      >
        <Checkbox
          checked={isCompleted}
          onCheckedChange={handleToggle}
          onClick={(e) => e.stopPropagation()}
        />
        <div className="flex-1 min-w-0">
          <span
            className={isCompleted ? "line-through text-muted-foreground" : ""}
          >
            {task.title}
          </span>
          {dueDate && (
            <Badge variant="outline" className="ml-2 text-xs">
              {dueDate}
            </Badge>
          )}
        </div>
        {task.labels?.map((l) => (
          <Badge key={l.id} variant="secondary" className="text-xs">
            {l.name}
          </Badge>
        ))}
      </div>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>編輯任務</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <Input
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="任務標題"
            />
            <div className="flex gap-2">
              <Button onClick={handleSave} disabled={saving}>
                {saving ? "儲存中..." : "儲存"}
              </Button>
              <Button variant="destructive" onClick={handleDelete}>
                刪除
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </>
  );
}
