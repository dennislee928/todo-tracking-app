"use client";

import Link from "next/link";

export function LogoutButton() {
  return (
    <Link
      href="/login"
      onClick={() => localStorage.removeItem("token")}
      className="text-sm text-muted-foreground hover:text-foreground"
    >
      登出
    </Link>
  );
}
