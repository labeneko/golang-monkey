package ast

import "github.com/miyohide/monkey/token"

// Node はASTのすべてのノードが実装しなければならないインターフェイス
type Node interface {
	// ノードが関連付けられるトークンのリテラル値を返す
	TokenLiteral() string
}

// Statement はノードの一種で文を表す
type Statement interface {
	Node
	statementNode()
}

// Expression はノードの一種で式を表す
type Expression interface {
	Node
	expressionNode()
}

// Program は構文解析器が生成するすべてのASTのルートノード
type Program struct {
	Statements []Statement
}

// TokenLiteral はノードが関連付けられているトークンのリテラル値を返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement は束縛の識別子と値を保持するもの
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral はノードが関連付けられるトークンのリテラル値を返す
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier は束縛の識別子を保持する
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral はノードが関連付けられるトークンのリテラル値を返す
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement はreturn文の構造を保持するもの
type ReturnStatement struct {
	Token       token.Token // 'return'トークン
	ReturnValue Expression  // 戻りとなる式を保持するフィールド
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral はノードが関連付けられるトークンのリテラル値を返す
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// ExpressionStatement は「x + 10;」のような式だけからなる行を保持するもの
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral はノードが関連付けられるトークンのリテラル値を返す
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
