package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
}

func loadFile(fileName string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "C:\\Users\\warug\\a2s\\sample\\"+fileName+".go", nil, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
	}

	return convert2SSA(f), nil
}

func convert2SSA(file *ast.File) string {
	return "hoge"
}

func getSSAFromGo(fileName string) (ssa.Function, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "C:\\Users\\warug\\a2s\\sample\\"+fileName+".go", nil, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
	}
	files := []*ast.File{f}
	pkg := types.NewPackage("hello", "")
	hello, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		return ssa.Function{}, err
	}

	return *hello.Func("f"), nil
}
