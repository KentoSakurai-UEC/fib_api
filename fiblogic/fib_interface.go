package fiblogic

import (
	"math/big"
)

/*
1. FibLogic は，フィボナッチ数計算ロジックのインターフェース．
2. アルゴリズムの変更や最適化を行う際，こちらのメソッドを満たすように実装する．
*/
type FibLogic interface {
	CalcFib(input int) *big.Int
}
