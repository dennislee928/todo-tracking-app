# Todo App - iOS

iOS 版本建置自 `device_shared` Flutter 專案。

## 建置方式

```bash
cd ../device_shared
flutter create . --platforms=ios,android  # 若 ios/ 尚未存在
flutter pub get
flutter build ios --release
```

或直接執行：

```bash
cd ../device_shared && flutter run
```

## 實機安裝

需 Apple Developer Program。產生 IPA：

```bash
cd ../device_shared
flutter build ipa
```

產物位於 `build/ios/ipa/`，可透過 TestFlight 或 Ad Hoc 分發。
