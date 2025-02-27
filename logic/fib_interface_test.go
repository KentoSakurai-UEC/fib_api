package logic

import (
	"math/big"
	"testing"
)

// FibLogic インターフェースのモック
type mockFibLogic struct{}

// input 値を 10倍して返すメソッドとして実装
func (m *mockFibLogic) CalcFib(input int) *big.Int {
	return big.NewInt(int64(input * 10))
}

// モック実装の FibLogic を注入し，計算ロジックが機能するかどうかのテスト．
func TestBusinessLogic_CalcFib(t *testing.T) {
	// インスタンス生成 (モック注入)
	mock := &mockFibLogic{}
	businessLogic := NewBusinessLogic(mock)

	tests := []struct {
		input    int
		expected *big.Int
	}{
		{input: 1, expected: big.NewInt(10)},
		{input: 15, expected: big.NewInt(150)},
		{input: 101, expected: big.NewInt(1010)},
		{input: 1000, expected: big.NewInt(10000)},
	}

	for _, tt := range tests {
		result := businessLogic.CalcFib(tt.input)
		if result.Cmp(tt.expected) != 0 {
			t.Errorf("expected: %v, calculated: %v", tt.expected, result)
		}
	}
}
