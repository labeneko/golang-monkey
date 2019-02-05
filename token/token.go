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
	// MINUS は-演算子
	MINUS = "-"
	// BANG は!演算子
	BANG = "!"
	// ASTERISK は*演算子
	ASTERISK = "*"
	// SLASH は/演算子
	SLASH = "/"
	// LT は<演算子
	LT = "<"
	// GT は>演算子
	GT = ">"
	// EQ は==演算子
	EQ = "=="
	// NOT_EQ は!=演算子
	NOT_EQ = "!="

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
	// LBRACKET は左大括弧
	LBRACKET = "["
	// RBRACKET は右大括弧
	RBRACKET = "]"
	// COLON はコロン（:）
	COLON = ":"

	// FUNCTION はFUNCTIONキーワード
	FUNCTION = "FUNCTION"
	// LET はLETキーワード
	LET = "LET"
	// TRUE はTRUEキーワード
	TRUE = "TRUE"
	// FALSE はFALSEキーワード
	FALSE = "FALSE"
	// IF はIFキーワード
	IF = "IF"
	// ELSE はELSEキーワード
	ELSE = "ELSE"
	// RETURN はRETURNキーワード
	RETURN = "RETURN"
	// STRING はSTRINGキーワード
	STRING = "STRING"
	// MACRO はMACROキーワード
	MACRO = "MACRO"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

// LookupIdent は渡された識別子がキーワードであればTokenType
// 定数を返す。そうでなければtoken.IDENTを返す。
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
