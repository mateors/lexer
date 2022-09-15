package main

import (
	"fmt"
	"log"
	"mastercode/lexer"
	"mastercode/repl"
	"mastercode/token"
	"os"
	"os/user"
)

func manualTest() {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
	  return true;
	} else {
		return false;
	}
	
	10 == 10;
	10 != 9`

	lex := lexer.New(input)
	for {
		tok := lex.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Println(tok.Type, tok.Literal)
	}
}

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s, Welcome to the REPL:\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
