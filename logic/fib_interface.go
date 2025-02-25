package logic

import (
	"math/big"
)

// BusinessLogic は，FibLogic が注入される上位の構造体．
type BusinessLogic struct {
	algorithm FibLogic
}

// コンストラクタ
func NewBusinessLogic(fl FibLogic) *BusinessLogic {
	return &BusinessLogic{algorithm: fl}
}

// FibLogic は，フィボナッチ数計算ロジックのインターフェース．
// アルゴリズムの変更や最適化を行う際，最低限こちらのメソッドを満たすように実装する．
type FibLogic interface {
	CalcFib(input int) *big.Int
}

// インターフェース定義のメソッドを実装．
// 計算に使用するのは注入されたロジック側の実装．
func (b *BusinessLogic) CalcFib(input int) *big.Int {
	return b.algorithm.CalcFib(input)
}
