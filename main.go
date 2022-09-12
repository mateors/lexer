package main

import (
	"fmt"
	"mastercode/lexer"
	"mastercode/token"
)

func main() {

	input := "x=a+(b*c){,};"
	
	lex := lexer.New(input)
	for {
		tok := lex.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Println(tok.Type, tok.Literal)
	}
}
