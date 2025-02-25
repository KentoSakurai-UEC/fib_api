# fib_api
フィボナッチ数列から n 番目の値を返すAPI

## 使用言語・インフラ
* Go >= 1.22
* Railway

## URI
[https://fibapi-production.up.railway.app/fib]()

## リクエストパラメータ
n: 自然数 (フィボナッチ数列における目的の index)

## リクエスト例
https://fibapi-production.up.railway.app/fib?n=99

## curlコマンド実行例
```
$ curl -X GET -H "Content-Type: application/json" "https://fibapi-production.up.railway.app/fib?n=99" 
```
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

## 一部補足説明
### handler/consts/consts.go
レスポンスヘッダ値や入力値検証関数の戻り値で使用する定数群を定義．<br>
handler の新バージョンが実装された場合にも共通使用できると考え，パッケージを分離．

### handler/handler_v1/handler.go
v1として実装．コンストラクタでは，`FibLogic` を満たしている `BusinessLogic` を注入する．<br>
入力値検証関数は validation パッケージから利用し，正常なリクエスト時のみフィボナッチ数計算を実行する．

### handler/handler_v1/response_types.go
v1として実装．正常系であるフィボナッチ数計算結果のレスポンスと，異常系であるエラー時のレスポンスを用意．

### handler/validation/validation.go
入力されたクエリパラメータ値を検証する．<br>
正常時はクエリパラメータ n の数値と "ok" メッセージを返し，異常時は 0値 とエラーメッセージを返す．<br>
handler の新バージョンが実装された場合にも共通使用できると考え，パッケージを分離．

### logic/fib_calc_2step.go
一度のフィボナッチ数計算で2ステップ進める機能を実装として採用．

### logic/fib_interface.go
フィボナッチ数計算ロジックのインターフェース `FibLogic` を定義．<br>
フィボナッチ数計算のアルゴリズム変更や新規実装の際は，こちらのインターフェースを最低限満たすように実装する．
