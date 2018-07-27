package parser

import (
	"github.com/miyohide/monkey/ast"
	"github.com/miyohide/monkey/lexer"
	"github.com/miyohide/monkey/token"
)

// Parser は字句解析器インスタンスへのポインタと2つのトークンを持つ型
type Parser struct {
	// 字句解析器インスタンスへのポインタ
	l *lexer.Lexer

	// 現在のトークン
	curToken token.Token
	// 次のトークン
	peekToken token.Token
}

// New は新しく字句解析器を作成し、2つのトークンを読み込んだParserを返す
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込んでcurTokenとpeekTokenの両方をセット
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken はcurTokenとpeekTokenをセットするヘルパーメソッド
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram はプログラムをParseする関数
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
