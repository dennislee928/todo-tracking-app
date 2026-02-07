import { Suspense } from "react";
import { AuthGuard } from "@/components/auth-guard";
import { LogoutButton } from "@/components/logout-button";
import { UpgradeButton } from "@/components/upgrade-button";
import { AdBanner } from "@/components/ad-banner";
import { AdSenseScript } from "@/components/adsense-script";
import { UpgradeCallback } from "@/components/upgrade-callback";
import { UserProvider } from "@/lib/user-context";
import Link from "next/link";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <AuthGuard>
      <UserProvider>
        <AdSenseScript />
        <Suspense fallback={null}>
          <UpgradeCallback />
        </Suspense>
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
            <div className="flex items-center gap-2">
              <UpgradeButton />
              <LogoutButton />
            </div>
          </nav>
          <main className="container max-w-2xl mx-auto py-8 px-4">
            {children}
            <AdBanner />
          </main>
        </div>
      </UserProvider>
    </AuthGuard>
  );
}
