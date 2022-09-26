package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/mateors/lexer/lexer"
	"github.com/mateors/lexer/parser"
	"github.com/mateors/lexer/repl"
	"github.com/mateors/lexer/token"
)

func lexerManualTest() {
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

func parserManualTest() {

	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		log.Fatal("Parse program returned nil")
	}

	fmt.Println(program.Statements, program.TokenLiteral())
	if len(program.Statements) != 3 {
		log.Fatalf("Program statement does not contains 3 statements, got %d", len(program.Statements))
	}
}

func parseExpressionManualTest() {

	input := `
	5+6;
	`
	l := lexer.New(input)
	p := parser.New(l)

	err := p.Errors()
	fmt.Println("err:", err)
	// fmt.Println("ParseProgram().String()", p.ParseProgram().String()) //do not execute this line before prog.Statements

	prog := p.ParseProgram()
	sts := prog.Statements
	fmt.Println(prog.String(), sts, len(sts), prog.TokenLiteral())

}

func ATest() {

	input := `
	if (1>0){

		let x=5;
	}
	`

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

	fmt.Printf("Hello %s, Welcome to the REPL:\n", user.Username) //RPPL = Read-Parse-Print-Loop
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

	//parserManualTest()
	//parseExpressionManualTest()

}
