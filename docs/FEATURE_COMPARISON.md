# Todo 追蹤 App 功能參考比較表

本文件依據 [Todoist](https://www.todoist.com/)、[ClickUp](https://clickup.com/)、[Microsoft To Do](https://www.microsoft.com/en-us/microsoft-365/microsoft-to-do-list-app)、[Clinked](https://www.clinked.com/features/task-management) 官方資訊整理，作為 MVP 功能設計參考。

---

## 一、核心功能比較

| 功能類別 | Todoist | ClickUp | Microsoft To Do | Clinked |
|----------|----------|---------|-----------------|---------|
| **任務擷取** | Quick Add 自然語言、Ramble 語音輸入 | AI 任務建立、自然語言輸入 | 基本輸入、Smart Due Date | 基本輸入、類別分組 |
| **專案/列表** | Projects、Sections | Spaces、Folders、Lists | 多列表、群組 | 任務類別（顏色標籤） |
| **子任務** | 支援 | 巢狀子任務、Checklist | 支援 Steps | 無明確子任務 |
| **優先級** | P1-P4 | 自訂優先級 | 重要性標記 | 無 |
| **標籤** | Labels | Tags、Custom Fields | Tags | Tags |
| **到期日** | 自然語言、遞迴 | 支援、自動排程 | 日期、遞迴 | 到期日、提醒 |
| **提醒** | 支援 | 支援 | 一次/遞迴提醒 | 自訂日期時間、Email |
| **檢視模式** | List、Calendar、Board | 15+ 種（List、Board、Gantt、Calendar、Mind Map 等） | 基本列表 | Kanban、Group Calendar |
| **智慧列表** | Today、Upcoming、Filters | Everything view、Filters | My Day、Assigned to Me | 依日期/負責人篩選 |
| **協作** | 共享專案、指派、評論 | 指派、評論、Chat、Docs | 共享列表、指派 | 指派、群組可見 |
| **進度追蹤** | Karma、生產力圖表 | 進度條、Sprint、Dashboards | 無 | 進度條、狀態（not started/in progress/completed） |
| **範本** | 50+ 範本 | 範本 | 無 | 無 |
| **整合** | 80+ 整合 | 50+ 整合、API | Outlook、Microsoft 365 | Zapier、Google Calendar |
| **平台** | 桌面、iOS、Android、Wear、瀏覽器 | 桌面、iOS、Android、瀏覽器 | 桌面、iOS、Android | Web、iOS、Android |
| **AI** | Ramble 語音 | Brain、Super Agents | 無 | 無 |
| **合規** | SOC2 | SOC2、ISO 27001、GDPR、HIPAA | Microsoft 365 合規 | GDPR、HIPAA、ISO 27001 |

---

## 二、細項功能對照

### 2.1 任務擷取與輸入

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| Quick Add | ✓ 自然語言 | ✓ | 基本輸入 | 基本輸入 |
| 語音輸入 | Ramble | 無 | 無 | 無 |
| 日期辨識 | 自然語言（如 "tomorrow"、"next Monday"） | 支援 | Smart Due Date | 日期選擇器 |
| 遞迴 | 支援 | 支援 | 支援 | 支援 daily/weekly/monthly/yearly |

### 2.2 組織結構

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| 階層 | Projects → Sections → Tasks | Spaces → Folders → Lists → Tasks | Lists → Tasks | Categories → Tasks |
| 顏色標籤 | 專案顏色 | 多種 | 列表主題 | 類別顏色 |
| Inbox | ✓ | ✓ | ✓ | 無專門 Inbox |

### 2.3 檢視模式

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| List | ✓ | ✓ | ✓ | ✓ |
| Board/Kanban | ✓ | ✓ | 無 | ✓ |
| Calendar | ✓ | ✓ | 無 | Group Calendar |
| Gantt | 無 | ✓ | 無 | 無 |
| Mind Map | 無 | ✓ | 無 | 無 |
| 自訂篩選 | Filters | Filters | 無 | 依日期/負責人 |

### 2.4 協作與權限

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| 專案/列表共享 | ✓ | ✓ | ✓ | 群組內可見 |
| 任務指派 | ✓ | ✓ | ✓ | ✓ |
| 評論 | ✓ | ✓ | 無 | Discussions |
| 角色權限 | Admin、Member | 多種角色 | 無 | 群組成員 |
| 連結分享 | ✓ | ✓ | 無 | 無 |

### 2.5 進度與生產力

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| 完成狀態 | ✓ | ✓ | ✓ | ✓ |
| 進度條 | 無 | ✓ | 無 | ✓ |
| 生產力統計 | Karma、圖表 | Dashboards、Reports | 無 | 無 |
| 活動紀錄 | ✓ | ✓ | 無 | Activity Stream |

### 2.6 檔案與附件

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| 附件 | ✓ | ✓ | ✓ 25MB | 無明確說明 |
| 文件編輯 | 無 | Docs | 無 | Online Document Editor |

### 2.7 認證與安全

| 細項 | Todoist | ClickUp | Microsoft To Do | Clinked |
|------|---------|---------|-----------------|---------|
| 登入方式 | Email、Google、Apple | Email、Google、SSO | Microsoft 帳號 | Email、SSO |
| 2FA | ✓ | ✓ | Microsoft 帳號 | ✓ |
| 合規認證 | SOC2 | SOC2、ISO 27001、GDPR、HIPAA | Microsoft 365 | GDPR、HIPAA、ISO 27001 |

---

## 三、MVP 建議優先實作功能

依參考 app 共性與計畫範圍，建議 MVP 包含：

| 優先級 | 功能 | 說明 |
|--------|------|------|
| P0 | 任務 CRUD | 建立、讀取、更新、刪除任務 |
| P0 | 專案/列表 | 建立專案、任務歸屬專案 |
| P0 | 子任務 | 任務下可有多個子任務 |
| P0 | 到期日 | 任務到期日、遞迴規則 |
| P0 | Today / Upcoming 檢視 | 今日到期、即將到期列表 |
| P0 | 認證 | Supabase Auth + 自建 JWT |
| P1 | 標籤 | Labels/Tags 分類 |
| P1 | 優先級 | P1–P4 或重要性 |
| P1 | 提醒 | 提醒時間 |
| P1 | 共享與指派 | 專案共享、任務指派 |
| P1 | List / Board 檢視 | 列表與看板切換 |
| P2 | 進度條 | 任務完成百分比 |
| P2 | 篩選 | 自訂篩選條件 |
| P2 | 範本 | 專案範本 |

---

## 四、參考連結

- [Todoist Features](https://www.todoist.com/features)
- [Todoist Task Management](https://www.todoist.com/task-management)
- [ClickUp Features](https://clickup.com/features)
- [Microsoft To Do](https://www.microsoft.com/en-us/microsoft-365/microsoft-to-do-list-app)
- [Clinked Task Management](https://www.clinked.com/features/task-management)
