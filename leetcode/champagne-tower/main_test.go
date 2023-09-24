package main

import "testing"

func Test_champagneTower(t *testing.T) {
	type args struct {
		poured      int
		query_row   int
		query_glass int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case 1",
			args: args{
				poured:      1,
				query_row:   1,
				query_glass: 1,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				poured:      2,
				query_row:   1,
				query_glass: 1,
			},
			want: 0.5,
		},
		{
			name: "case 3",
			args: args{
				poured:      100000009,
				query_row:   33,
				query_glass: 17,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				poured:      100000009,
				query_row:   33,
				query_glass: 17,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := champagneTower(tt.args.poured, tt.args.query_row, tt.args.query_glass); got != tt.want {
				t.Errorf("champagneTower() = %v, want %v", got, tt.want)
			}
		})
	}
}
