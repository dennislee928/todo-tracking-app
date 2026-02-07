# Todo Tracking App

參考 Todoist、ClickUp、Microsoft To Do、Clinked 的 todo 追蹤應用程式，具備任務管理、專案、標籤、Today/Upcoming 檢視等功能。

## 專案結構

```
├── web-be/           # Go 後端 (Gin, GORM, Swagger, gRPC)
├── web-ui/           # Next.js 15+ 前端
├── device_shared/    # Flutter 雙平台 (iOS + Android) 共用專案
├── device-ios/       # iOS 建置說明
├── device-android/   # Android 建置說明
├── docs/             # 功能比較表等文件
├── .github/workflows/# CI/CD pipelines
├── render.yaml       # Render Blueprint 部署
├── fly.toml          # Fly.io 後端部署
└── fly.web-ui.toml   # Fly.io 前端部署
```

## 快速開始

### 後端 (web-be)

```bash
cd web-be

# 設定環境變數
export DATABASE_URL="postgres://user:pass@host:5432/todo?sslmode=require"

# 執行 migration (需安裝 golang-migrate)
migrate -path database/migrations -database "$DATABASE_URL" up

# 執行
go run ./cmd/server
```

API 文件: http://localhost:8080/swagger/index.html

### 前端 (web-ui)

```bash
cd web-ui

# 設定 API URL
export NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1

npm install
npm run dev
```

### 行動裝置 (device_shared)

需先安裝 [Flutter](https://flutter.dev/docs/get-started/install)。

```bash
cd device_shared

# 若無 ios/android 目錄，先建立
flutter create . --platforms=ios,android

flutter pub get
flutter run  # 選擇模擬器或實機
```

Android 建置 APK: `flutter build apk --release`
iOS 建置: `flutter build ios --release` (需 Xcode 與 Apple Developer)

## CI/CD

- `build-web-be.yml` - Docker 建置後端
- `build-web-ui.yml` - Docker 建置前端
- `build-device-android.yml` - 建置 Android APK 並上傳 artifact
- `build-device-ios.yml` - 建置 iOS (macOS runner)
- `vulnerability-scan.yml` - Trivy 漏洞掃描

## 部署

- **Render**: 使用 `render.yaml` Blueprint，連接 GitHub 後一鍵部署
- **Fly.io**: `fly deploy` (web-be) 或 `fly deploy --config fly.web-ui.toml` (web-ui)
