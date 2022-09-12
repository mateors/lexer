package main

import (
	"fmt"
	"mastercode/lexer"
	"mastercode/token"
)

func main() {

	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);`

	lex := lexer.New(input)
	for {
		tok := lex.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Println(tok.Type, tok.Literal)
	}
}
