# fib_api
フィボナッチ数列から n 番目の値を返すAPI

## 使用言語・インフラ
* Go >= 1.22
* Railway

## URI
~~[https://fibapi-production.up.railway.app/fib]()~~

## リクエストパラメータ
n: 自然数 (フィボナッチ数列における目的の index)

## リクエスト例
~~https://fibapi-production.up.railway.app/fib?n=99~~

## ステータスコード
| ステータス | 説明 |
| ---------------------| ---------- |
| 200 OK               | 正常にリクエストが成功 |
| 400 Bad Request      | 不正なリクエストにより失敗 |

## レスポンス例
* ステータス: 200 OK
```json
{
  "result":  218922995834555169026
}
```
* ステータス: 400 Bad request
```json
{
  "status":  400,
  "message": "Bad request: Specify parameter n."
}
```
```json
{
  "status":  400,
  "message": "Bad request: Specify a natural number for parameter n."
}
```

## 構成と概要
```
.
├── go.mod
├── handler
│   ├── consts
│   │   └── consts.go
│   ├── handler_v1
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   └── response_types.go
│   └── validation
│       ├── validation.go
│       └── validation_test.go
├── logic
│   ├── fib_calc_2step.go
│   ├── fib_calc_2step_test.go
│   ├── fib_interface.go
│   └── testdata
│       └── fibo.txt
└── main.go
```

* `go.mod`: モジュール定義
* `handler`: APIハンドラー
  * `consts`: ハンドラーで使用する定数群
  * `handler_v1`: ハンドラーの機能定義およびレスポンス定義
  * `validation`: 入力値検証
* `logic`: フィボナッチ数計算用ロジック
  * `testdata`: 第１項から第１００２項までのフィボナッチ数列を記載．テストに使用 ([https://fibonnacci.aimary.com/]())<br>
    ※第１００２項以降も計算可能
* `main.go`: エントリーポイント
