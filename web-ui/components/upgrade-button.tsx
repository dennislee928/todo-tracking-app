"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { useUser } from "@/lib/user-context";
import { api } from "@/lib/api";

export function UpgradeButton() {
  const { user, loading, refetch } = useUser();
  const [pending, setPending] = useState(false);

  if (loading || !user) return null;
  if (user.is_premium) {
    return (
      <span className="text-sm text-muted-foreground px-2">已升級</span>
    );
  }

  const handleUpgrade = async () => {
    setPending(true);
    try {
      const base =
        typeof window !== "undefined"
          ? window.location.origin
          : "http://localhost:3000";
      const { url } = await api.subscription.createCheckoutSession(
        `${base}?upgrade=success`,
        `${base}?upgrade=cancel`
      );
      if (url) window.location.href = url;
    } catch (err) {
      console.error(err);
      alert("無法建立付款連結，請稍後再試");
    } finally {
      setPending(false);
    }
  };

  return (
    <Button
      variant="outline"
      size="sm"
      onClick={handleUpgrade}
      disabled={pending}
    >
      {pending ? "處理中..." : "升級去廣告"}
    </Button>
  );
}
