# Todo App - Flutter (iOS & Android)

共用 Flutter 專案，同時支援 iOS 與 Android。

## 前置需求

- [Flutter SDK](https://flutter.dev/docs/get-started/install) 3.6+
- Xcode（macOS，僅 iOS 建置需要）
- Android Studio 或 Android SDK（Android 建置需要）

## 首次設定

若 `android/` 或 `ios/` 目錄不存在，請執行：

```bash
flutter create . --platforms=ios,android
```

此指令會產生平台所需的設定檔與專案結構。

## 執行

```bash
flutter pub get
flutter run  # 選擇模擬器或實機
```

## 建置

### Android APK（可直接安裝至實機）

```bash
flutter build apk --release
```

產物：`build/app/outputs/flutter-apk/app-release.apk`

### Android App Bundle（Play Store 上架）

```bash
flutter build appbundle --release
```

### iOS（需 macOS + Xcode）

```bash
flutter build ios --release
```

實機安裝需 Apple Developer 帳號，使用 `flutter build ipa` 產生 IPA。

## API 設定

建置時可指定後端 API URL：

```bash
flutter build apk --dart-define=API_URL=https://your-api.com/api/v1
```

- Android 模擬器連本地：`http://10.0.2.2:8080/api/v1`
- iOS 模擬器連本地：`http://localhost:8080/api/v1`
- 實機連本地：使用電腦實際 IP，如 `http://192.168.1.100:8080/api/v1`
