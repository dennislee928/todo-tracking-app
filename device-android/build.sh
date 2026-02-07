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

case "${1:-apk}" in
  run)
    flutter run -d android
    ;;
  apk)
    flutter build apk --release
    echo "APK: build/app/outputs/flutter-apk/app-release.apk"
    ;;
  aab)
    flutter build appbundle --release
    echo "AAB: build/app/outputs/bundle/release/app-release.aab"
    ;;
  *)
    echo "Usage: $0 {run|apk|aab}"
    echo "  run  - 在 Android 模擬器/實機執行"
    echo "  apk  - 建置 APK（可直接安裝）"
    echo "  aab  - 建置 App Bundle（Play Store 上架）"
    exit 1
    ;;
esac
