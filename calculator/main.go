package main

import (
	"calculator/parser"
	"calculator/tokenizer"
	"fmt"
	"log"
)

func er(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dump(s string) {
	tokens, err := tokenizer.Tokenize(s)
	er(err)
	exp, err := parser.Parse(tokens)
	er(err)
	fmt.Println(exp)
}

func main() {
	dump("1+4*5*(9+10+((9)))")
}
