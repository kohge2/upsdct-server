## 概要
- コーディングテストで作成
- 請求書の登録と請求書の取得ができるサービスのAPI

## APIドキュメント
swaggerでもドキュメントを作成した
ローカルで立ち上げた後、
`http://localhost:8080/docs/swagger/index.html` で確認できる


### POST /api/invoices
新しい請求書データを作成する <br>

#### Request
**Header**
|パラメータ|値|
|----|----|
|Content-Type|application/json|

**パラメータ**
|パラメータ|型|説明|必須|備考|
|----|----|----|---|---|
|paidAmount|int|支払金額|yes||
|paidDueDate|string|支払期日|yes|RFC3339の文字列(例:日本時間22時の場合 "2025-05-31T22:00:00+09:00")|
|partnerCompanyID|string|取引先企業ID|yes||

#### Response 
**200** 
```json
{
  "ok": 1
}
```
**その他エラー**
```json
{
  "error": {
    "type": "エラータイプ(アプリケーション側で定義したもの)",
    "code": "ステータスコード",
    "message": "エラーメッセージ"
  }
}
```


### GET /api/invoices
指定期間内に支払いが発生する請求書データの一覧を取得する

ログイン中のユーザーが所属する企業が発行した請求書の一覧を支払い期日で絞り込んで取得することができる

#### Request
**クエリパラメータ**
|パラメータ|型|説明|必須|備考|
|----|----|----|---|---|
|startDate|string|支払い期日で絞り込む時の開始日時|no|RFC3339の文字列(例:日本時間22時の場合 "2025-05-31T22:00:00+09:00")|
|endDate|string|支払い期日で絞り込む時の終了日時|no|RFC3339の文字列(例:日本時間22時の場合 "2025-05-31T22:00:00+09:00")|

#### Response
**200**
```json
{
  "invoices": [
    {
      "billedAmount": 0,
      "commission": 0,
      "id": "string",
      "invoiceStatus": "string",
      "paidAmount": 0,
      "paidDueDate": "string",
      "partnerCompanyBankAccount": {
        "accountHolderName": "string",
        "accountNumber": "string",
        "accountType": "string",
        "bankName": "string",
        "branchName": "string"
      },
      "partnerCompanyID": "string",
      "partnerCompanyName": "string",
      "publishedDate": "string",
      "tax": 0
    }
  ]
}
```
※補足 <br>
「paidDueDate」: 支払い期日, 「publishedDate」: Dateのみ(時間は無し) <br>
「paidAmount」: 支払い金額, 　「billedAmount」: 請求金額(支払金額 に手数料4%を加えたものに更に手数料の消費税を加えたも の) <br>
「partnerCompanyBankAccount」: 取引先の口座情報 <br>

**その他エラー**
```json
{
  "error": {
    "type": "エラータイプ(アプリケーション側で定義したもの)",
    "code": "ステータスコード",
    "message": "エラーメッセージ"
  }
}
```

## テーブルについて
https://github.com/kohge2/upsdct-server/blob/main/sql/schema.sql 

## 立ち上げ方法
```
go run main.go
```

## テスト実行方法
```
make test
```

## ディレクトリの説明
```
.
├── adapter # データアクセス層。domain/repositoryの実装。データの取得や永続化に関する処理
├── config # 環境変数や定数の管理など、プロジェクト全体の設定を管理する。
├── database # データベース関連の処理
├── docs # swagger用のファイル。https://github.com/swaggo/swag を使用し、`swag init` コマンドでAPIドキュメントを自動生成。
├── domain 
|    ├── models　# エンティティ、データベースモデル
|    └── repository　# データの永続化に関するインタフェース。adapterと対応
├── sql # テーブル定義用のSQL
├── testmock # テスト用のmock
├── scripts # 手元で実行したスクリプトを保管している
├── usecases # ビジネスロジック層。handlerから呼び出されてdomainを操作するためのロジックを提供。
├── utils # プロジェクト全体で再利用可能な日付処理や暗号化などのユーティリティ関数を提供。
└── web # プレゼンテーション層。HTTPインターフェイスを提供する処理。
    ├── handler　# リクエストを受け取り、適切なusecaseを呼び出してレスポンスを返す。
    ├── middleware　# 認証やエラーハンドリングなどのmiddlewareを管理する。
    ├── request　# リクエストのデータ構造。
    └── response　# レスポンスのデータ構造。
```

## 補足
- ログインユーザーのID: どちらもログイン前提のAPIだが、ログイン機能がないので、userIDはtest1として処理している。contextにuserIDをsetしている。
  (参考: https://github.com/kohge2/upsdct-server/blob/main/web/middleware/auth.go )

## 実施メモ
**数日に分けて少しずつ進めて合計約2時間**
- 共通で使用する処理(utils,config(環境変数),エラー周りなど) https://github.com/kohge2/upsdct-server/pull/1
- モデル、テーブル定義など https://github.com/kohge2/upsdct-server/pull/2

**まとめて時間とって約3時間**
- 請求書登録API と swagger関連の設定 と domainのテスト https://github.com/kohge2/upsdct-server/pull/3
- 請求書取得API  https://github.com/kohge2/upsdct-server/pull/4

**調べながら約2時間**
- adapter,usecaseのテスト、その他修正 https://github.com/kohge2/upsdct-server/pull/6

**TODO 追加対応したい**
- 環境構築をDockerでできるように、初期データのセットアップ簡単に
- テスト増やす
