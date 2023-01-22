package calculator

import (
	"reflect"
	"testing"
)

func TestResult_AddTo(t *testing.T) {
	tests := []struct {
		name     string
		v1       *Result
		v2       *Result
		expected *Result
	}{
		{
			name: "same denom",
			v1: &Result{
				Num:   1,
				Denom: 5,
				Sign:  1,
			},
			v2: &Result{
				Num:   2,
				Denom: 5,
				Sign:  1,
			},
			expected: &Result{
				Num:   3,
				Denom: 5,
				Sign:  1,
			},
		},
		{
			name: "different denom",
			v1: &Result{
				Num:   1,
				Denom: 4,
				Sign:  1,
			},
			v2: &Result{
				Num:   1,
				Denom: 2,
				Sign:  1,
			},
			expected: &Result{
				Num:   3,
				Denom: 4,
				Sign:  1,
			},
		},
		{
			name: "different denom",
			v1: &Result{
				Num:   1,
				Denom: 4,
				Sign:  1,
			},
			v2: &Result{
				Num:   1,
				Denom: 2,
				Sign:  1,
			},
			expected: &Result{
				Num:   3,
				Denom: 4,
				Sign:  1,
			},
		},
		{
			name: "1/5 + 1/4",
			v1: &Result{
				Num:   1,
				Denom: 5,
				Sign:  1,
			},
			v2: &Result{
				Num:   1,
				Denom: 4,
				Sign:  1,
			},
			expected: &Result{
				Num:   9,
				Denom: 20,
				Sign:  1,
			},
		},
		{
			name: "subtraction",
			v1: &Result{
				Num:   1,
				Denom: 6,
				Sign:  1,
			},
			v2: &Result{
				Num:   1,
				Denom: 2,
				Sign:  -1,
			},
			expected: &Result{
				Num:   1,
				Denom: 3,
				Sign:  -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v1.AddTo(tt.v2)
			tt.v1.normalize()
			if !reflect.DeepEqual(tt.v1, tt.expected) {
				t.Errorf("wrong result: got=%v expected=%v", tt.v1, tt.expected)
			}
		})
	}
}
