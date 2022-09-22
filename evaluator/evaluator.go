package evaluator

import (
	"mastercode/ast"
	"mastercode/object"
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	//Statements
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	//Expression
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {

	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
