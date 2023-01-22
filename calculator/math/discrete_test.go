package math

import (
	"fmt"
	"math/rand"
	"testing"
)

func slowLCD(a, b int64) int64 {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	for i := b; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func TestLCD(t *testing.T) {
	rand.Seed(1)
	for i := 0; i < 20; i++ {
		a, b := rand.Intn(300), rand.Intn(300)
		t.Run(fmt.Sprintf("[%d,%d]", a, b), func(t *testing.T) {
			expected := slowLCD(int64(a), int64(b))
			got := LCD(int64(a), int64(b))
			if expected != got {
				t.Errorf("LCD() = %v, want %v", got, expected)
			}
		})
	}
}
