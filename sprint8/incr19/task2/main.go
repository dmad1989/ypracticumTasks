package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
    
func main() {
     ids := 77
     id := ids + 1
     fmt.Println("id равно:", id/2 )
}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if c, ok := n.(*ast.Ident); ok && c.Name == `id` {
			c.Name = `Ident`
		}
		return true
	})

	printer.Fprint(os.Stdout, fset, f)
}
