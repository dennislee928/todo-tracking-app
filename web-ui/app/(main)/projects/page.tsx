"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { api } from "@/lib/api";
import type { Project } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

export default function ProjectsPage() {
  const [projects, setProjects] = useState<Project[]>([]);
  const [loading, setLoading] = useState(true);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [name, setName] = useState("");
  const [creating, setCreating] = useState(false);

  const fetchProjects = async () => {
    try {
      const data = await api.projects.list();
      setProjects(data);
    } catch {
      setProjects([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProjects();
  }, []);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;
    setCreating(true);
    try {
      await api.projects.create({ name: name.trim() });
      setName("");
      setDialogOpen(false);
      fetchProjects();
    } finally {
      setCreating(false);
    }
  };

  if (loading) {
    return <div className="text-center py-12">載入中...</div>;
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold">專案</h1>
          <p className="text-muted-foreground">管理您的專案與任務</p>
        </div>
        <Button onClick={() => setDialogOpen(true)}>新增專案</Button>
      </div>

      <div className="grid gap-4">
        {projects.length === 0 ? (
          <p className="text-muted-foreground text-center py-12">
            還沒有專案，點擊「新增專案」開始
          </p>
        ) : (
          projects.map((proj) => (
            <Link key={proj.id} href={`/projects/${proj.id}`}>
              <Card className="hover:bg-muted/50 transition-colors">
                <CardHeader className="pb-2">
                  <h3 className="font-semibold">{proj.name}</h3>
                </CardHeader>
                <CardContent>
                  <span className="text-sm text-muted-foreground">
                    點擊查看任務
                  </span>
                </CardContent>
              </Card>
            </Link>
          ))
        )}
      </div>

      <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>新增專案</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleCreate} className="space-y-4">
            <Input
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="專案名稱"
              required
            />
            <Button type="submit" disabled={creating}>
              {creating ? "建立中..." : "建立"}
            </Button>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  );
}
