package token

// TokenType は「整数」や「閉じ括弧」などを区別するもの
type TokenType string

// Token はTokenTypeとリテラルを保持する型
type Token struct {
	Type    TokenType
	Literal string
}
