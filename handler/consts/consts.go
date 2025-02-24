package consts

// handler および validation で使用する定数を定義．

// handler 使用分
const (
	// レスポンスヘッダー値
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json"
)

// validation 使用分
const (
	// 入力値検証関数が正常時に 第２戻り値 として返すメッセージ
	ValOK = "ValOK"

	// 入力値検証関数がエラー発生時に 第１戻り値 として返す数値
	ErrValNum = 0

	// 入力値検証関数がエラー発生時に 第２戻り値 として返すエラーメッセージ
	ErrMsgMissingParameter = "Bad request: Specify parameter n."
	ErrMsgInvalidParameter = "Bad request: Specify a natural number for parameter n."
)
