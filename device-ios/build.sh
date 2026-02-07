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

case "${1:-ios}" in
  run)
    flutter run -d ios
    ;;
  build)
    flutter build ios --release --no-codesign
    echo "Build complete. Open ios/Runner.xcworkspace in Xcode to run on simulator."
    ;;
  ipa)
    flutter build ipa
    echo "IPA: build/ios/ipa/*.ipa"
    ;;
  *)
    echo "Usage: $0 {run|build|ipa}"
    echo "  run   - 在 iOS 模擬器/實機執行"
    echo "  build - 建置 iOS (no-codesign，模擬器用)"
    echo "  ipa   - 建置 IPA（需 Apple Developer 憑證）"
    exit 1
    ;;
esac
