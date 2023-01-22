# BNF

```
<expression> := [<addOp>] <terms> {<addOp> <terms>}
<addOp> := "-" | "+"
<term> := <factor> {<multOp> <factor>}
<multOp> := "*" | "/"
<factor> := <number> | "(" <expression> ")"
<number> := <digit> {<digit>}
<digit> := "0" ... "9"
```

# Tokenizer

Tokenizer splits a given string into valid tokens expressed in the BNF.

# Parser

Parser make a tree of valid BNF by given tokens tokenized.

# TODO

- [ ] test parser
- [ ] calculate the result
