package fiblogic

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
)

// https://fibonnacci.aimary.com/fibo.txt よりダウンロードしたデータを用いてテスト．
func TestTwoStepFibCalclator(t *testing.T) {
	// テストデータ読み込み
	testdata, err := os.Open("testdata/fibo.txt")
	if err != nil {
		t.Fatalf("failed to open the test data: %v", err)
	}
	defer testdata.Close()

	// インスタンス生成
	twoStepFibCalculator := NewTwoStepFibCalc()

	// ループで使用
	line := 1
	scanner := bufio.NewScanner(testdata)

	for scanner.Scan() {
		lineData := scanner.Text()

		// フィボナッチ数列の index と値を取得
		lineDataSlice := strings.Fields(lineData)

		// 数列の index を数値に変換
		n, err := strconv.Atoi(lineDataSlice[0])
		if err != nil {
			t.Fatalf("failed to convert n: %v, line at: %d", err, line)
		}

		// フィボナッチ数を big.Int に変換
		expectedFib := new(big.Int)
		if _, ok := expectedFib.SetString(lineDataSlice[1], 10); !ok {
			t.Fatalf("failed to convert Fib, line at: %d", line)
		}

		// 計算結果とテストデータを比較
		result := twoStepFibCalculator.CalcFib(n)
		if result.Cmp(expectedFib) != 0 {
			t.Errorf("Failed on line %d , expected: %s, calclated: %s", line, expectedFib.String(), result.String())
		}

		line++
	}
}
