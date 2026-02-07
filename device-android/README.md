# Todo App - Android

Android 版本建置自 `device_shared` Flutter 專案。

## 前置需求

- [Flutter SDK](https://flutter.dev/docs/get-started/install) 3.6+
- [Android Studio](https://developer.android.com/studio) 或 Android SDK
- Java 17+（建議使用 Temurin）

## 建置方式

### 方式一：使用建置腳本（推薦）

```bash
# 從 device-android 目錄執行
./build.sh run   # 在 Android 模擬器/實機執行
./build.sh apk   # 建置 APK（可直接安裝）
./build.sh aab   # 建置 App Bundle（Play Store 上架）
```

### 方式二：手動執行

```bash
cd ../device_shared

# 若 android/ 尚未完整建立
flutter create . --platforms=ios,android

flutter pub get
flutter run -d android
```

## 產物路徑

| 建置類型 | 產物路徑 |
|----------|----------|
| APK | `device_shared/build/app/outputs/flutter-apk/app-release.apk` |
| App Bundle | `device_shared/build/app/outputs/bundle/release/app-release.aab` |

可直接將 APK 安裝至 Android 實機或模擬器。

## App Bundle（Play Store）

```bash
cd ../device_shared
flutter build appbundle --release
```

產物 `app-release.aab` 可上傳至 Google Play Console。

## API 設定

開發時可指定後端 API URL：

```bash
flutter run --dart-define=API_URL=http://10.0.2.2:8080/api/v1
```

- **Android 模擬器連本地**：`http://10.0.2.2:8080/api/v1`
- **實機連本地**：使用電腦實際 IP，如 `http://192.168.1.100:8080/api/v1`

## CI/CD

GitHub Actions 使用 `build-device-android.yml`，於 `device_shared/` 或 `device-android/` 變更時觸發。會建置 APK 並上傳為 artifact。

## 相關檔案

- 共用程式碼：`../device_shared/lib/`
- 平台設定：`../device_shared/android/`
