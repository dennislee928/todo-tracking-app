const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

function getToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("token");
}

async function fetchAPI<T>(
  path: string,
  options: RequestInit = {}
): Promise<T> {
  const token = getToken();
  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...(options.headers as Record<string, string>),
  };
  if (token) {
    (headers as Record<string, string>)["Authorization"] = `Bearer ${token}`;
  }
  const res = await fetch(`${API_BASE}${path}`, { ...options, headers });
  if (!res.ok) {
    const err = await res.json().catch(() => ({ error: res.statusText }));
    throw new Error(err.error || res.statusText);
  }
  if (res.status === 204) return {} as T;
  return res.json();
}

export const api = {
  auth: {
    register: (email: string, password: string) =>
      fetchAPI<{ token: string; user: { id: string; email: string } }>(
        "/auth/register",
        {
          method: "POST",
          body: JSON.stringify({ email, password }),
        }
      ),
    login: (email: string, password: string) =>
      fetchAPI<{ token: string; user: { id: string; email: string } }>(
        "/auth/login",
        {
          method: "POST",
          body: JSON.stringify({ email, password }),
        }
      ),
  },
  tasks: {
    list: (projectId?: string) =>
      fetchAPI<Task[]>(
        projectId ? `/tasks?project_id=${projectId}` : "/tasks"
      ),
    today: () => fetchAPI<Task[]>("/tasks/today"),
    upcoming: (days = 7) =>
      fetchAPI<Task[]>(`/tasks/upcoming?days=${days}`),
    get: (id: string) => fetchAPI<Task>(`/tasks/${id}`),
    create: (data: Partial<Task>) =>
      fetchAPI<Task>("/tasks", {
        method: "POST",
        body: JSON.stringify(data),
      }),
    update: (id: string, data: Partial<Task>) =>
      fetchAPI<Task>(`/tasks/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),
    delete: (id: string) =>
      fetchAPI<void>(`/tasks/${id}`, { method: "DELETE" }),
  },
  projects: {
    list: () => fetchAPI<Project[]>("/projects"),
    get: (id: string) => fetchAPI<Project>(`/projects/${id}`),
    create: (data: { name: string; color?: string }) =>
      fetchAPI<Project>("/projects", {
        method: "POST",
        body: JSON.stringify(data),
      }),
    update: (id: string, data: { name?: string; color?: string }) =>
      fetchAPI<Project>(`/projects/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      }),
    delete: (id: string) =>
      fetchAPI<void>(`/projects/${id}`, { method: "DELETE" }),
  },
};

export interface Task {
  id: string;
  title: string;
  description?: string;
  project_id?: string;
  user_id: string;
  priority: number;
  status: string;
  progress: number;
  due_date?: string;
  reminder_at?: string;
  created_at: string;
  updated_at: string;
  labels?: { id: string; name: string; color: string }[];
}

export interface Project {
  id: string;
  name: string;
  color?: string;
  user_id: string;
  created_at: string;
}
