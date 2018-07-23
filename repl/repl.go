package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/miyohide/monkey/lexer"
	"github.com/miyohide/monkey/token"
)

// PROMPT はREPLで入力を促すプロンプト
const PROMPT = ">> "

// Start は入力を待ち受け、読み込めなくなるまで無限ループする
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
