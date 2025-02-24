package v1

import (
	"encoding/json"
	"net/http"

	"github.com/KentoSakurai-UEC/fib_api/fiblogic"
	"github.com/KentoSakurai-UEC/fib_api/handler/consts"
	"github.com/KentoSakurai-UEC/fib_api/handler/validation"
)

// Handler は，リクエストに応じた処理を行う．
type Handler struct {
	fibLogic fiblogic.FibLogic
}

// Handler のコンストラクタ
func NewHandler(fl fiblogic.FibLogic) *Handler {
	return &Handler{fibLogic: fl}
}

// handleメソッド
func (h *Handler) FibHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(consts.HeaderContentType, consts.ContentTypeJSON)

	// 入力パラメータの検証
	input, valMsg := validation.ParseInputValue(r.URL.Query().Get("n"))
	if valMsg != consts.ValOK {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewErrorResponse(http.StatusBadRequest, valMsg))
		return
	}

	// 検証通過のため，フィボナッチ数を計算
	result := h.fibLogic.CalcFib(input)
	json.NewEncoder(w).Encode(NewFibCalcResponse(result.String()))
}
