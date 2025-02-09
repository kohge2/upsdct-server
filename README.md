## 概要
- コーディングテストで作成
- 請求書の登録と請求書の取得ができるサービスのAPI

## APIドキュメント
ローカルで立ち上げ後、
`http://localhost:8080/docs/swagger/index.html` で確認可能

## 立ち上げ方法
```
go run main.go
```
## 補足
- ログインユーザーのID: ログイン前提の機能なのでuserIDはtest1として処理される

## 実施メモ

数日に分けて少しずつ進めて合計約2時間
- https://github.com/kohge2/upsdct-server/pull/1
- https://github.com/kohge2/upsdct-server/pull/2

まとめて時間とって約3時間
- https://github.com/kohge2/upsdct-server/pull/3
- https://github.com/kohge2/upsdct-server/pull/4

提出後にやった内容
- TODO テスト増やすなど
- TODO 環境構築をDockerでできるように(現状デバッグしづらい)
