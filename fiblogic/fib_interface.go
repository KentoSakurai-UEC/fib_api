package fiblogic

import (
	"math/big"
)

/*
1. FibLogic は，フィボナッチ数計算ロジックのインターフェースです．
2. CalclateFib メソッドを実装する際は，input 値に関して負値の考慮は不要です．(※HTTPリクエスト側の実装で対応)
*/
type FibLogic interface {
	CalcFib(input int) *big.Int
}
