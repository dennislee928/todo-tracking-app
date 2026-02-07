"use client";

import { useUser } from "@/lib/user-context";

const ADSENSE_ID = process.env.NEXT_PUBLIC_ADSENSE_ID;
const ADSENSE_SLOT = process.env.NEXT_PUBLIC_ADSENSE_SLOT || "1234567890";

export function AdBanner() {
  const { user, loading } = useUser();

  if (loading || !user) return null;
  if (user.is_premium) return null;
  if (!ADSENSE_ID) {
    // Placeholder when AdSense not configured
    return (
      <div className="w-full max-w-[728px] mx-auto my-4 p-4 border border-dashed border-muted-foreground/30 rounded-lg bg-muted/30 text-center text-sm text-muted-foreground">
        廣告區塊（設定 NEXT_PUBLIC_ADSENSE_ID 以顯示 AdSense）
      </div>
    );
  }

  return (
    <div className="w-full max-w-[728px] mx-auto my-4 text-center">
      <ins
        className="adsbygoogle"
        style={{ display: "block" }}
        data-ad-client={ADSENSE_ID}
        data-ad-slot={ADSENSE_SLOT}
        data-ad-format="auto"
        data-full-width-responsive="true"
      />
    </div>
  );
}
