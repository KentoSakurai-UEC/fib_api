package v1

import (
	"encoding/json"
	"net/http"

	"github.com/KentoSakurai-UEC/fib_api/handler/consts"
	"github.com/KentoSakurai-UEC/fib_api/handler/validation"
	"github.com/KentoSakurai-UEC/fib_api/logic"
)

type Handler struct {
	businessLogic *logic.BusinessLogic
}

func NewHandler(b *logic.BusinessLogic) *Handler {
	return &Handler{businessLogic: b}
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
	result := h.businessLogic.CalcFib(input)
	json.NewEncoder(w).Encode(NewFibCalcResponse(result.String()))
}
