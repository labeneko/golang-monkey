package ast

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
