package evaluator

import (
	"github.com/miyohide/monkey/ast"
	"github.com/miyohide/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
