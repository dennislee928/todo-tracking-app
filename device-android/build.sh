#!/bin/bash
# Android 建置腳本 - 委派至 device_shared 執行
set -e
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT/device_shared"

if ! command -v flutter &> /dev/null; then
  echo "Error: Flutter not found. Install from https://flutter.dev/docs/get-started/install"
  exit 1
fi

# 若平台目錄不完整，先建立
if [ ! -d "android/app" ]; then
  echo "Creating Android platform..."
  flutter create . --platforms=ios,android
fi

flutter pub get

# Production API (Fly.io backend)
API_URL="${API_URL:-}"
[ -n "$API_URL" ] || API_URL="https://todo-web-be.fly.dev/api/v1"
DART_DEFINE="--dart-define=API_URL=$API_URL"

case "${1:-apk}" in
  run)
    flutter run -d android $DART_DEFINE
    ;;
  apk)
    flutter build apk --release $DART_DEFINE
    echo "APK: build/app/outputs/flutter-apk/app-release.apk"
    ;;
  aab)
    flutter build appbundle --release $DART_DEFINE
    echo "AAB: build/app/outputs/bundle/release/app-release.aab"
    ;;
  *)
    echo "Usage: $0 {run|apk|aab}"
    echo "  run  - 在 Android 模擬器/實機執行"
    echo "  apk  - 建置 APK（可直接安裝）"
    echo "  aab  - 建置 App Bundle（Play Store 上架）"
    echo ""
    echo "API 預設: $API_URL (Fly.io backend)"
    echo "自訂: API_URL=http://10.0.2.2:8080/api/v1 $0 run"
    exit 1
    ;;
esac
