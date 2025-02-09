## 概要
- コーディングテストで作成
- 請求書の登録と請求書の取得ができるサービスのAPI

## APIドキュメント
swaggerでドキュメントを作成した

ローカルで立ち上げた後、
`http://localhost:8080/docs/swagger/index.html` で確認できる

## 立ち上げ方法
```
go run main.go
```

## テスト実行方法
```
make test
```

## 補足
- ログインユーザーのID: どちらもログイン前提のAPIだが、ログイン機能がないので、userIDはtest1として処理している。contextにuserIDをsetしている。
  (参考: https://github.com/kohge2/upsdct-server/blob/main/web/middleware/auth.go )

## 実施メモ
数日に分けて少しずつ進めて合計約2時間
- feat: 共通で使用する処理(utils,config(環境変数),エラー周りなど) https://github.com/kohge2/upsdct-server/pull/1
- feat: モデルとDDL https://github.com/kohge2/upsdct-server/pull/2

まとめて時間とって約3時間
- feat: 請求書登録API と swagger関連の設定 と domainのテスト https://github.com/kohge2/upsdct-server/pull/3
- feat: 請求書取得API  https://github.com/kohge2/upsdct-server/pull/4

調べながら約2時間
- adapter,usecaseのテスト、その他修正

追加対応したい
- TODO 環境構築をDockerでできるように
