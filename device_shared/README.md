# Todo App - Flutter (iOS & Android)

共用 Flutter 專案，同時支援 iOS 與 Android。此專案為 `device-ios` 與 `device-android` 的共用程式碼與建置來源。

## 專案結構

```
device_shared/
├── lib/
│   ├── main.dart              # 入口、AuthWrapper
│   ├── screens/               # 畫面
│   │   ├── login_screen.dart
│   │   ├── register_screen.dart
│   │   ├── home_screen.dart   # 底部導覽（Today / Upcoming / Projects）
│   │   ├── today_screen.dart
│   │   ├── upcoming_screen.dart
│   │   ├── projects_screen.dart
│   │   └── project_detail_screen.dart
│   ├── services/              # API 與認證
│   │   ├── api_service.dart   # REST API client
│   │   └── auth_service.dart  # JWT 登入/登出
│   └── widgets/
│       └── task_tile.dart     # 任務列表元件
├── android/                   # Android 平台設定
├── ios/                       # iOS 平台設定
├── scripts/
│   └── setup_platforms.sh     # 首次建立 ios/android 目錄
└── pubspec.yaml
```

## 功能

- **認證**：登入、註冊、JWT 儲存
- **Today**：今日到期任務
- **Upcoming**：未來 7 天任務
- **Projects**：專案列表與詳情
- **任務 CRUD**：建立、編輯、完成、刪除任務

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
flutter run --dart-define=API_URL=https://your-api.com/api/v1
flutter build apk --dart-define=API_URL=https://your-api.com/api/v1
```

- **Android 模擬器連本地**：`http://10.0.2.2:8080/api/v1`
- **iOS 模擬器連本地**：`http://localhost:8080/api/v1`
- **實機連本地**：使用電腦實際 IP，如 `http://192.168.1.100:8080/api/v1`

## 平台專用建置

- **iOS**：參見 `../device-ios/README.md`，可使用 `../device-ios/build.sh`
- **Android**：參見 `../device-android/README.md`，可使用 `../device-android/build.sh`
