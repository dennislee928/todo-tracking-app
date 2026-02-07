"use client";

import { useEffect } from "react";
import { useSearchParams } from "next/navigation";
import { useUser } from "@/lib/user-context";

export function UpgradeCallback() {
  const searchParams = useSearchParams();
  const { refetch } = useUser();

  useEffect(() => {
    const upgrade = searchParams.get("upgrade");
    if (upgrade === "success") {
      refetch();
      window.history.replaceState({}, "", window.location.pathname);
    }
  }, [searchParams, refetch]);

  return null;
}
