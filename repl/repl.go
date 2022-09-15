package repl

import (
	"bufio"
	"fmt"
	"io"
	"mastercode/lexer"
	"mastercode/token"
)

//REPL = Read Eval Print Loop
//Similar to console
//Or interactive mode
//the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the interpreter and starts again
//this package takes input evaluate or process it and gives output and repeat it again and again

const PROMPT = ">>"

func Start(input io.Reader, output io.Writer) {

	scanner := bufio.NewScanner(input)

	for {

		fmt.Fprintf(output, PROMPT)
		canYouScan := scanner.Scan()
		if !canYouScan {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}

	}

}
