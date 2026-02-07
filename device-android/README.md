# Todo App - Android

Android 版本建置自 `device_shared` Flutter 專案。

## 建置方式

```bash
cd ../device_shared
flutter create . --platforms=ios,android  # 若 android/ 尚未存在
flutter pub get
flutter build apk --release
```

產物：`device_shared/build/app/outputs/flutter-apk/app-release.apk`

可直接安裝至 Android 實機或模擬器。

## App Bundle（Play Store）

```bash
cd ../device_shared
flutter build appbundle --release
```

產物：`build/app/outputs/bundle/release/app-release.aab`
