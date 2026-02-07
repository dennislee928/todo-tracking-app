#!/bin/bash
# iOS 建置腳本 - 委派至 device_shared 執行
set -e
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT/device_shared"

if ! command -v flutter &> /dev/null; then
  echo "Error: Flutter not found. Install from https://flutter.dev/docs/get-started/install"
  exit 1
fi

# 若平台目錄不完整，先建立
if [ ! -d "ios/Runner.xcodeproj" ]; then
  echo "Creating iOS platform..."
  flutter create . --platforms=ios,android
fi

flutter pub get

# Production API (Fly.io backend)
API_URL="${API_URL:-}"
[ -n "$API_URL" ] || API_URL="https://todo-web-be.fly.dev/api/v1"
DART_DEFINE="--dart-define=API_URL=$API_URL"

case "${1:-run}" in
  run)
    # 若無 iOS 裝置，自動改用 macOS
    if flutter devices 2>&1 | grep -qE '\bios\b|iPhone|iPad'; then
      flutter run -d ios $DART_DEFINE
    else
      echo "No iOS device found. Running on macOS instead..."
      flutter run -d macos $DART_DEFINE
    fi
    ;;
  macos)
    flutter run -d macos $DART_DEFINE
    ;;
  build)
    flutter build ios --release --no-codesign $DART_DEFINE
    echo "Build complete. Open ios/Runner.xcworkspace in Xcode to run on simulator."
    ;;
  ipa)
    flutter build ipa $DART_DEFINE
    echo "IPA: build/ios/ipa/*.ipa"
    ;;
  *)
    echo "Usage: $0 {run|macos|build|ipa}"
    echo "  run   - 在 iOS 模擬器/實機執行；若無則改用 macOS"
    echo "  macos - 直接在 macOS 桌面執行"
    echo "  build - 建置 iOS (no-codesign，模擬器用)"
    echo "  ipa   - 建置 IPA（需 Apple Developer 憑證）"
    echo ""
    echo "API 預設: $API_URL (Fly.io backend)"
    echo "自訂: API_URL=http://localhost:8080/api/v1 $0 run"
    exit 1
    ;;
esac
