package v1

// handler および validation で使用するレスポンスを定義．

// フィボナッチ数計算結果のレスポンス (正常系)
type FibCalcResponse struct {
	Result string `json:"result"`
}

// エラー時のレスポンス (異常系)
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// 正常系レスポンスのコンストラクタ
func NewFibCalcResponse(result string) *FibCalcResponse {
	return &FibCalcResponse{Result: result}
}

// 異常系レスポンスのコンストラクタ
func NewErrorResponse(status int, message string) *ErrorResponse {
	return &ErrorResponse{Status: status, Message: message}
}
