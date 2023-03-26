package main

import (
	"fmt"
)

func convertDigit(c rune) (uint64, error) {
	switch {
	case 'A' <= c && c <= 'Z':
		return 10 + uint64(c-'A'), nil
	case 'a' <= c && c <= 'z':
		return 10 + uint64(c-'a'), nil
	case '0' <= c && c <= '9':
		return uint64(c - '0'), nil
	}
	return 0, fmt.Errorf("failed to convert: %v", c)
}

func convertToDigit(v uint64) (byte, error) {
	switch {
	case v <= 9:
		return byte(v + '0'), nil
	case 10 <= v && v <= 35:
		return byte(v - 10 + 'A'), nil
	}
	return 0, fmt.Errorf("failed to convert to digit: %v", v)
}

func toUint64(s string, base uint64) (uint64, error) {
	if base <= 1 || base > 36 {
		return 0, fmt.Errorf("base must be in 2 <= base <= 36: %v", base)
	}
	b := uint64(1)
	res := uint64(0)
	rs := make([]rune, len(s))
	for i, c := range s {
		rs[len(s)-1-i] = c
	}
	for _, c := range rs {
		d, err := convertDigit(c)
		if err != nil {
			return 0, err
		}
		if d > base {
			return 0, fmt.Errorf("wrong digit for the base=%v: %v", base, d)
		}
		// TODO: check overflow
		res += d * b
		b *= base
	}
	return res, nil
}

func reverse(bs []byte) []byte {
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	return bs
}

func toBase(a uint64, base uint64) (string, error) {
	if base <= 1 || base > 36 {
		return "", fmt.Errorf("base must be in 2 <= base <= 36: %v", base)
	}
	bts := make([]byte, 0, 1)
	if a == 0 {
		return "0", nil
	}
	for a > 0 {
		c, err := convertToDigit(a % base)
		if err != nil {
			return "", nil
		}
		bts = append(bts, c)
		a /= base
	}
	return string(reverse(bts)), nil
}

func ConvertBase(s string, fromBase, to uint64) (string, error) {
	v, err := toUint64(s, fromBase)
	if err != nil {
		return "", err
	}
	return toBase(v, to)
}
