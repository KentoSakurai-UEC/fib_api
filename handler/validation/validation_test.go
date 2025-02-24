package validation

import (
	"testing"

	"github.com/KentoSakurai-UEC/fib_api/handler/consts"
)

func TestParseInputValue(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantInt int
		wantMsg string
	}{
		{
			name:    "valid1",
			input:   "1",
			wantInt: 1,
			wantMsg: consts.ValOK,
		},
		{
			name:    "valid1000",
			input:   "1000",
			wantInt: 1000,
			wantMsg: consts.ValOK,
		},
		{
			name:    "empty",
			input:   "",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgMissingParameter,
		},
		{
			name:    "zero",
			input:   "0",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "negative",
			input:   "-10",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "non-numeric_A",
			input:   "A",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "non-numeric_aaa",
			input:   "aaa",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "non-numeric_あ",
			input:   "あ",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "non-numeric_例",
			input:   "例",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "decimals_0.0",
			input:   "0.0",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "non-numeric_10.0",
			input:   "10.0",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_!",
			input:   "!",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_\"",
			input:   "\"",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_#",
			input:   "#",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_$",
			input:   "$",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_%",
			input:   "%",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_&",
			input:   "&",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_'",
			input:   "'",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_(",
			input:   "(",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_)",
			input:   ")",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_+",
			input:   "+",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_-",
			input:   "-",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_*",
			input:   "*",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_/",
			input:   "/",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_=",
			input:   "=",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_|",
			input:   "|",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_\\",
			input:   "\\",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_^",
			input:   "^",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_~",
			input:   "~",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_`",
			input:   "`",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_@",
			input:   "@",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_{",
			input:   "{",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_}",
			input:   "}",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_[",
			input:   "[",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_]",
			input:   "]",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_「",
			input:   "「",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_」",
			input:   "」",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_;",
			input:   ";",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_:",
			input:   ":",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_<",
			input:   "<",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_>",
			input:   ">",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_?",
			input:   "?",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_.",
			input:   ".",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
		{
			name:    "special_,",
			input:   ",",
			wantInt: consts.ErrValNum,
			wantMsg: consts.ErrMsgInvalidParameter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt, gotMsg := ParseInputValue(tt.input)
			if gotInt != tt.wantInt {
				t.Errorf("ParseInputValue() gotInt = %v, want = %v", gotInt, tt.wantInt)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("ParseInputValue() gotMsg = %v, want = %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
