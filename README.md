# SyncMate
SyncMateは共同生活を支援するプラットフォームです．
DDD，マイクロサービスアーキテクチャの学習を目的に実践的なアプリケーションとして開発しました．

## 機能(構想)

SyncMate共同生活に役立つ機能を提供します．各機能はマイクロサービスとして提供されREST APIを介して連携します．

- **ユーザーグループ管理**
  - ユーザーが自由にグループを作成でき，グループ単位で他のツールを利用できる．
  - グループメンバーにはカスタムロールを割り当て柔軟に権限を制御することができる．
  - グループへメンバーを追加する際，Emailアドレスを指定して招待を送る．
  - 招待を受けたメンターが承認/拒否を行う．
- **コスト管理**
  - 光熱費を月単位で管理できる．
  - 共同の家計簿をつけられる．
  - 引き落とし日を管理できる．
  - 指定した比率でグループメンバーのコストを算出できる．
  - コストを種別に年単位でグラフ化できる．
- **タスク管理**
  - 共同タスクを作成管理できる．
  - タスクには担当者を割り振ることができる．
  - タスクには期限を設定できる．
  - タスクには優先度を設定できる．
- **スケジュール管理**
  - 共同のカレンダー管理する．
- **グループチャット**
  - リアルタイムのチャットができる．

## 技術スタック

- バックエンド
  - Go, Echo, gorm
- フロントエンド
  - TypeScript, React, Next.js, Chakura UI
- インフラ
  - Terraform, Vercel, k8s? 未定
- CI/CD
  - GitHub Actions, PipeCD? 未定

## アーキテクチャ
### システムアーキテクチャ
BFFアーキテクチャを採用．バックエンドAPIはマイクロサービスとして分割しフロントエンドとの中間としてAPI Gatewayを配置する．
![システムアーキテクチャ](./image/architecture.png?raw=true)

### バックエンドアーキテクチャ
各APIで様々なアーキテクチャを採用する想定．（学習目的）
ユーザーグループ管理サービスではオニオンアーキテクチャを採用
```
├── cmd
│   ├── api
│   ├── initialization
│   └── migration
├── docs
└── internal
    ├── domain
    │   ├── model
    │   ├── repository
    │   └── service
    ├── infrastructure
    │   ├── database
    │   │   ├── initialize
    │   │   │   └── sql
    │   │   ├── migrate
    │   │   └── postgres
    │   └── persistence
    ├── interactor
    ├── presentation
    │   └── http
    │       ├── handler
    │       ├── middleware
    │       └── router
    ├── usecase
    │   └── mocks
    └── utils

```
