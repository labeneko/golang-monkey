package evaluator

import (
	"github.com/miyohide/monkey/ast"
	"github.com/miyohide/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
