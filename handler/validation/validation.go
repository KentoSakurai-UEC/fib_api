package validation

// handler で使用するバリデーション関数を定義．

import (
	"strconv"

	"github.com/KentoSakurai-UEC/fib_api/handler/consts"
)

// ParseInputValue は入力値を検証します．
// 正常時はクエリパラメータnの数値と"ok"メッセージを返し，異常時は0値とエラーメッセージを返します．
func ParseInputValue(inputStr string) (int, string) {
	// パラメータ未指定の場合
	if inputStr == "" {
		return consts.ErrValNum, consts.ErrMsgMissingParameter
	}

	inputInt, err := strconv.Atoi(inputStr)
	// int値変換が不可 または 1未満の値が入力された場合
	if err != nil || inputInt < 1 {
		return consts.ErrValNum, consts.ErrMsgInvalidParameter
	}

	return inputInt, consts.ValOK
}
