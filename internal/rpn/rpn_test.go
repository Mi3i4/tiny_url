package rpn

import (
	"errors"
	"testing"
)

func TestRpnCalc(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		result  float64
		wantErr error
	}{
		{
			name:   "single number",
			args:   "42",
			result: 42,
		},
		{
			name:   "simple",
			args:   "3 4 +",
			result: 7.0,
		},
		{
			name:   "chain",
			args:   "3 4 + 5 *",
			result: 35,
		},
		{
			name:   "subtract",
			args:   "6 2 -",
			result: 4,
		},
		{
			name:   "divide",
			args:   "10 2 /",
			result: 5,
		},
		{
			name:    "not enough",
			args:    "3 +",
			wantErr: ErrNotEnoughOperands,
		},
		{
			name:    "div zero",
			args:    "5 0 /",
			wantErr: ErrDivByZero,
		},
		{
			name:    "leftover",
			args:    "3 4",
			wantErr: ErrInvalidExpr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RPNCalc(tt.args)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got err %v, want err %v", err, tt.wantErr)
			}

			if err == nil && got != tt.result {
				t.Errorf("got result %v, want %v", got, tt.result)
			}
		})
	}
}
