package main

import (
	"log"
	"strconv"
	"testing"
)

func parseBinaryExpressionWithoutError(b string) int {
	n, err := strconv.ParseInt(b, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

func TestInsertBits(t *testing.T) {
	type args struct {
		n int
		m int
		i uint
		j uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: parseBinaryExpressionWithoutError("1010010111"),
				m: parseBinaryExpressionWithoutError("11011"),
				i: 3,
				j: 3 + 5,
			},
			want: parseBinaryExpressionWithoutError("1011011111"),
		},
		{
			args: args{
				n: parseBinaryExpressionWithoutError("10101010101"),
				m: parseBinaryExpressionWithoutError("1101001"),
				i: 3,
				j: 3 + 7,
			},
			want: parseBinaryExpressionWithoutError("11101001101"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertBits(tt.args.n, tt.args.m, tt.args.i, tt.args.j); got != tt.want {
				gb := strconv.FormatInt(int64(got), 2)
				wb := strconv.FormatInt(int64(tt.want), 2)
				t.Errorf("InsertBits() = %v, want %v", gb, wb)
			}
		})
	}
}

func TestBestReversedSize(t *testing.T) {
	tests := []struct {
		name         string
		binaryFormat string
		want         int
	}{
		{
			name:         "given sample",
			binaryFormat: "11011101111",
			want:         8,
		},
		{
			name:         "random",
			binaryFormat: "110101001011",
			want:         4,
		},
		{
			name:         "01",
			binaryFormat: "01",
			want:         2,
		},
		{
			name:         "0",
			binaryFormat: "0",
			want:         1,
		},
		{
			name:         "101",
			binaryFormat: "101",
			want:         3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parseBinaryExpressionWithoutError(tt.binaryFormat)
			if got := BestReversedSize(uint(n)); got != tt.want {
				t.Errorf("BestReversedSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
