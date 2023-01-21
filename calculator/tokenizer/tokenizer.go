package tokenizer

import "fmt"

/*
Example:

	4+6*9/10
	3 + 9 * (20 + 1)
	-9
	+ 4 / 10

	tokens: + - * / ( ) [0-9]+
*/

func isDigit(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

// Tokenize split a given string into valid tokens
func Tokenize(s string) ([]string, error) {
	var res []string
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-', '*', '/', '(', ')':
			res = append(res, s[i:i+1])
			continue
		case ' ':
			continue
		}
		if isDigit(s[i]) {
			cur := i
			for ; i < len(s) && isDigit(s[i]); i++ {
			}
			res = append(res, s[cur:i])
			i--
		} else {
			return nil, fmt.Errorf("invalid token found: '%v' in '%v'", s[i], s)
		}
	}
	return res, nil
}
