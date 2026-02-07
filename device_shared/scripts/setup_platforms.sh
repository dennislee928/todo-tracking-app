#!/bin/bash
# 為 device_shared 建立 iOS 與 Android 平台目錄
# 需已安裝 Flutter: https://flutter.dev/docs/get-started/install

set -e
cd "$(dirname "$0")/.."

if ! command -v flutter &> /dev/null; then
  echo "Error: Flutter not found. Install from https://flutter.dev/docs/get-started/install"
  exit 1
fi

if [ ! -d "android" ] || [ ! -d "ios" ]; then
  echo "Creating iOS and Android platforms..."
  flutter create . --platforms=ios,android
  echo "Done. Run 'flutter run' to start."
else
  echo "android/ and ios/ already exist."
fi
