import { AuthGuard } from "@/components/auth-guard";
import { LogoutButton } from "@/components/logout-button";
import Link from "next/link";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <AuthGuard>
      <div className="min-h-screen">
        <nav className="border-b px-4 py-3 flex items-center justify-between">
          <div className="flex gap-6">
            <Link href="/" className="font-semibold">
              Today
            </Link>
            <Link href="/upcoming" className="text-muted-foreground hover:text-foreground">
              Upcoming
            </Link>
            <Link href="/projects" className="text-muted-foreground hover:text-foreground">
              專案
            </Link>
          </div>
          <LogoutButton />
        </nav>
        <main className="container max-w-2xl mx-auto py-8 px-4">
          {children}
        </main>
      </div>
    </AuthGuard>
  );
}
