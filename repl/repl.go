package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/chkda/monkey-lang-interpreter/token"

	"github.com/chkda/monkey-lang-interpreter/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		// l := lexer.New(line)
		m := lexer.New(line)

		for tok := m.NextToken(); tok.Type != token.EOF; tok = m.NextToken() {
			// fmt.Fprintf(out, "%v\n", tok)
			fmt.Println(tok)
		}

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%v\n", tok)
		// }
	}
}
