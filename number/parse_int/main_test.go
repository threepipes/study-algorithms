package main

import (
	"strconv"
	"testing"
)

func Test_convertDigit(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "should convert an uppercase",
			args: args{'A'},
			want: 10,
		},
		{
			name: "should convert an uppercase",
			args: args{'Z'},
			want: 35,
		},
		{
			name: "should convert a lowercase",
			args: args{'b'},
			want: 11,
		},
		{
			name: "should convert a lowercase",
			args: args{'z'},
			want: 35,
		},
		{
			name: "should convert a digit",
			args: args{'8'},
			want: 8,
		},
		{
			name: "should convert a digit",
			args: args{'0'},
			want: 0,
		},
		{
			name:    "should not convert a illegal case",
			args:    args{'-'},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertDigit(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toUint64(t *testing.T) {
	type args struct {
		s    string
		base uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "should convert base 10",
			args: args{
				s:    "123456",
				base: 10,
			},
			want: 123456,
		},
		{
			name: "should convert base 2",
			args: args{
				s:    "1001",
				base: 2,
			},
			want: 9,
		},
		{
			name: "should convert base 16",
			args: args{
				s:    "FAB3",
				base: 16,
			},
			want: 16*16*16*15 + 16*16*10 + 16*11 + 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toUint64(tt.args.s, tt.args.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("toUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToDigit(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr bool
	}{
		{
			name: "convert to number",
			args: args{1},
			want: '1',
		},
		{
			name: "convert to alphabet",
			args: args{10},
			want: 'A',
		},
		{
			name: "convert to alphabet(z)",
			args: args{35},
			want: 'Z',
		},
		{
			name:    "should not convert",
			args:    args{36},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToDigit(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBase(t *testing.T) {
	type args struct {
		a    uint64
		base uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should convert base 10 number",
			args: args{
				a:    1022,
				base: 10,
			},
			want: "1022",
		},
		{
			name: "should convert base 16 number",
			args: args{
				a:    16*16*16*15 + 16*16*10 + 16*11 + 3,
				base: 16,
			},
			want: "FAB3",
		},
		{
			name: "should convert base 2 number",
			args: args{
				a:    25,
				base: 2,
			},
			want: "11001",
		},
		{
			name: "should convert 0",
			args: args{
				a:    0,
				base: 35,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBase(tt.args.a, tt.args.base)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ignoreError[T any](v T, err error) T {
	return v
}

func TestConvertBase(t *testing.T) {
	type args struct {
		s        string
		fromBase uint64
		to       uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "convert base 2 number to base 16 number",
			args: args{
				s:        "1111" + "1010" + "0010" + "0110",
				fromBase: 2,
				to:       16,
			},
			want: "FA26",
		},
		{
			name: "convert base 10 number to base 2 number",
			args: args{
				s:        "197",
				fromBase: 10,
				to:       2,
			},
			want: strconv.FormatInt(197, 2),
		},
		{
			name: "convert base 16 number to base 10 number",
			args: args{
				s:        "abc174e",
				fromBase: 16,
				to:       10,
			},
			want: strconv.FormatUint(ignoreError(strconv.ParseUint("abc174e", 16, 64)), 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertBase(tt.args.s, tt.args.fromBase, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
