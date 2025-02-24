package fiblogic

import (
	"math/big"
)

// TwoStepFibCalc は，2ステップずつ進めるフィボナッチ数計算を実装した構造体．
type TwoStepFibCalc struct{}

// FibLogic インターフェースを満たすことをチェック
var _ FibLogic = (*TwoStepFibCalc)(nil)

// コンストラクタ
func NewTwoStepFibCalc() *TwoStepFibCalc {
	return &TwoStepFibCalc{}
}

// CalcFib は，input として指定された n 番目に対するフィボナッチ数を返すメソッド．
func (ts *TwoStepFibCalc) CalcFib(input int) *big.Int {
	if input <= 2 {
		return big.NewInt(1)
	}

	first := big.NewInt(int64(input % 2))
	second := big.NewInt(1)

	for i := 0; i < input/2; i++ {
		first.Add(first, second)
		second.Add(second, first)
	}
	return first
}
