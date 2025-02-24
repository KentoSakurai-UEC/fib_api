package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KentoSakurai-UEC/fib_api/fiblogic"
	"github.com/KentoSakurai-UEC/fib_api/handler/consts"
)

func TestFibHandler(t *testing.T) {
	tests := []struct {
		name         string
		query        string
		wantStatus   int
		wantResponse string
	}{
		{
			name:         "valid",
			query:        "99",
			wantStatus:   http.StatusOK,
			wantResponse: "218922995834555169026",
		},
		{
			name:         "empty",
			query:        "",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgMissingParameter,
		},
		{
			name:         "negative",
			query:        "-10",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgInvalidParameter,
		},
		{
			name:         "zero",
			query:        "0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgInvalidParameter,
		},
		{
			name:         "non-numeric_A",
			query:        "A",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgInvalidParameter,
		},
		{
			name:         "decimal_0.0",
			query:        "0.0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgInvalidParameter,
		},
		{
			name:         "decimal_10.0",
			query:        "10.0",
			wantStatus:   http.StatusBadRequest,
			wantResponse: consts.ErrMsgInvalidParameter,
		},
	}

	// インスタンス生成
	twoStepFibCalc := fiblogic.NewTwoStepFibCalc()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/fib" + "?n=" + tt.query
			req := httptest.NewRequest("GET", url, nil)
			rr := httptest.NewRecorder()

			h := &Handler{
				fibLogic: twoStepFibCalc,
			}

			// ハンドラー呼び出し
			h.FibHandler(rr, req)

			// Status コード検証
			if rr.Code != tt.wantStatus {
				t.Errorf("gotStatus = %d, want = %d", rr.Code, tt.wantStatus)
			}

			// レスポンスBodyから、message を検証 (正常系、異常系別に検証)
			if rr.Code == 200 {
				var fibCalcResponse FibCalcResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &fibCalcResponse); err != nil {
					t.Fatalf("failed to unmarshall JSON")
				}
				// Result キーの値を取得
				responseMsg := fibCalcResponse.Result

				if responseMsg != tt.wantResponse {
					t.Errorf("gotResponse = %s, want = %s", responseMsg, tt.wantResponse)
				}

			} else {
				var errResponse ErrorResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &errResponse); err != nil {
					t.Fatalf("failed to unmarshall JSON")
				}
				// Message キーの値を取得
				responseMsg := errResponse.Message

				if responseMsg != tt.wantResponse {
					t.Errorf("gotResponse = %s, want = %s", responseMsg, tt.wantResponse)
				}
			}
		})
	}
}
