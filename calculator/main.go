package main

import (
	"calculator/calculator"
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
	res, err := calculator.Calculate(exp)
	er(err)
	fmt.Println(res)
}

func main() {
	dump("-1+4*5/(9+10+((9)))")
	dump("10/(9-10-1)")
	dump("34/125+99/98+34/13-2")
	dump("123*23/45*34/12-1")
	dump("(13+24/25)/18/7")
}
