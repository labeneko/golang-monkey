package token

// TokenType は「整数」や「閉じ括弧」などを区別するもの
type TokenType string

// Token はTokenTypeとリテラルを保持する型
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL はトークンや文字が未知であることを示す
	ILLEGAL = "ILLEGAL"
	// EOF はファイル終端。構文解析器にここで停止してよいと伝える
	EOF = "EOF"

	// IDENT は識別子（add, foobar, x, y, ...)
	IDENT = "IDENT"
	// INT は整数
	INT = "INT"

	// ASSIGN は=演算子
	ASSIGN = "="
	// PLUS は+演算子
	PLUS = "+"

	// COMMA はデリミタ（,）
	COMMA = ","
	// SEMICOLON はデリミタ（;）
	SEMICOLON = ";"

	// LPAREN は左小括弧
	LPAREN = "("
	// RPAREN は右小括弧
	RPAREN = ")"
	// LBRACE は左中括弧
	LBRACE = "{"
	// RBRACE は右中括弧
	RBRACE = "}"

	// FUNCTION はFUNCTIONキーワード
	FUNCTION = "FUNCTION"
	// LET はLETキーワード
	LET = "LET"
)
