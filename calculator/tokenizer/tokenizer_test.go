package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal valid",
			args: args{"1 + 9"},
			want: []string{"1", "+", "9"},
		},
		{
			name: "contains two or more digits number",
			args: args{"12 - 1398"},
			want: []string{"12", "-", "1398"},
		},
		{
			name: "only one number",
			args: args{"121398"},
			want: []string{"121398"},
		},
		{
			name: "only one single digit number",
			args: args{"0"},
			want: []string{"0"},
		},
		{
			name: "no spaces expression",
			args: args{"-9+19*8"},
			want: []string{"-", "9", "+", "19", "*", "8"},
		},
		{
			name: "complex expression",
			args: args{"-9+19*(8+8)"},
			want: []string{"-", "9", "+", "19", "*", "(", "8", "+", "8", ")"},
		},
		{
			name:    "error: contains invalid token",
			args:    args{"3 + 9x"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
