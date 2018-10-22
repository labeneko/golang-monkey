package repl

import (
	"bufio"
	"fmt"
	"github.com/miyohide/monkey/parser"
	"io"

	"github.com/miyohide/monkey/lexer"
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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

// printParserErrors は第一引数で取った io.Writer に対して
// 第二引数の []string の内容を整形して出力する
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
