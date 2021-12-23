package main

import (
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
)

func main() {
    fset := token.NewFileSet()
    f, err := parser.ParseFile(fset, "C:\\Users\\warug\\a2s\\sample\\sample.go", nil, parser.Mode(0))
    if err != nil {
	fmt.Println("ast error")
    }

    for _, d := range f.Decls {
        ast.Print(fset, d)
    }
}
