package evaluator

import (
	"strings"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.VintObject {
	Package := &object.Package{
		Name:         node.Name,
		Env:          env,
		Scope:        object.NewEnclosedEnvironment(env),
		PrivateNames: make(map[string]bool),
	}

	Package.Scope.Define("@", Package)

	// Evaluate package contents and track private identifiers
	evalPackageContents(node.Block, Package)

	// Automatically run the init function if it exists
	if initFunc, ok := Package.Scope.Get("init"); ok {
		if fn, ok := initFunc.(*object.Function); ok {
			applyFunction(fn, []object.VintObject{}, 0)
		}
	}
	
	env.Define(node.Name.Value, Package)
	return Package
}

// evalPackageContents evaluates package contents and tracks private members
func evalPackageContents(block *ast.BlockStatement, pkg *object.Package) object.VintObject {
	var result object.VintObject
	
	for _, statement := range block.Statements {
		result = Eval(statement, pkg.Scope)
		if result != nil && result.Type() == object.ERROR_OBJ {
			return result
		}
		
		// Track private identifiers based on naming convention
		switch stmt := statement.(type) {
		case *ast.LetStatement:
			if strings.HasPrefix(stmt.Name.Value, "_") {
				pkg.DefinePrivate(stmt.Name.Value)
			}
		case *ast.ConstStatement:
			if strings.HasPrefix(stmt.Name.Value, "_") {
				pkg.DefinePrivate(stmt.Name.Value)
			}
		}
	}
	
	return result
}
