package repl

import (
	"bufio"
	"fmt"
	"io"
	"mastercode/evaluator"
	"mastercode/lexer"
	"mastercode/object"
	"mastercode/parser"
)

//REPL = Read Eval Print Loop
//Similar to console
//Or interactive mode
//the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the interpreter and starts again
//this package takes input evaluate or process it and gives output and repeat it again and again

const PROMPT = ">>"

func Start(input io.Reader, output io.Writer) {

	scanner := bufio.NewScanner(input)
	env := object.NewEnvironment()

	for {

		fmt.Fprintf(output, PROMPT)
		canYouScan := scanner.Scan()
		if !canYouScan {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		p := parser.New(lex)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(output, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(output, evaluated.Inspect())
			io.WriteString(output, "\n")
		}

		// io.WriteString(output, program.String())
		// io.WriteString(output, "\n")

		// for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		// 	fmt.Fprintf(output, "%v %s\n", tok.Type, tok.Literal)
		// }

	}

}

// func printParserErrors(out io.Writer, errors []string) {
// 	for _, msg := range errors {
// 		io.WriteString(out, "\t"+msg+"\n")
// 	}
// }

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some issues here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
