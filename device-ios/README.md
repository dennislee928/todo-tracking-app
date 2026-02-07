# Todo App - iOS

iOS 版本建置自 `device_shared` Flutter 專案。

## 前置需求

- [Flutter SDK](https://flutter.dev/docs/get-started/install) 3.6+
- macOS
- [Xcode](https://developer.apple.com/xcode/)（Apple 開發者工具，含 iOS 模擬器）

## 建置方式

### 方式一：使用建置腳本（推薦）

```bash
# 從 device-ios 目錄執行
./build.sh run    # 在 iOS 模擬器執行
./build.sh build  # 建置 iOS（模擬器用，無需簽章）
./build.sh ipa    # 建置 IPA（需 Apple Developer 憑證）
```

### 方式二：手動執行

```bash
cd ../device_shared

# 若 ios/ 尚未完整建立
flutter create . --platforms=ios,android

flutter pub get
flutter run -d ios
```

## Xcode 專案

建置完成後，可用 Xcode 開啟：

```bash
cd ../device_shared
open ios/Runner.xcworkspace
```

在 Xcode 中選擇模擬器或實機，點擊 Run 即可執行。

## 實機安裝

需 [Apple Developer Program](https://developer.apple.com/programs/)（$99/年）。

```bash
cd ../device_shared
flutter build ipa
```

產物位於 `build/ios/ipa/`，可透過 TestFlight 或 Ad Hoc 分發。

## API 設定

開發時可指定後端 API URL：

```bash
flutter run --dart-define=API_URL=http://localhost:8080/api/v1
```

- **模擬器連本地**：`http://localhost:8080/api/v1`
- **實機連本地**：使用電腦實際 IP，如 `http://192.168.1.100:8080/api/v1`

## CI/CD

GitHub Actions 使用 `build-device-ios.yml`，於 `device_shared/` 或 `device-ios/` 變更時觸發。目前在 macOS runner 上執行 `flutter build ios --release --no-codesign`。若要產生簽章 IPA，需在 GitHub Secrets 設定憑證與 provisioning profile。

## 相關檔案

- 共用程式碼：`../device_shared/lib/`
- 平台設定：`../device_shared/ios/`
