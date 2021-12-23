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
	"os"

	"github.com/k0kubun/pp"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "C:\\Users\\warug\\a2s\\sample\\sample.go", nil, parser.Mode(0))
	if err != nil {
		fmt.Println(err)
	}

	files := []*ast.File{f}

	// Create the type-checker's package.
	pkg := types.NewPackage("hello", "")

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	hello, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Print(err) // type error in some package
		return
	}

	// Print out the package.
	pp.Println("%#v", hello.Func("add"))

	fmt.Printf("-----------------------------------------")

	hello.WriteTo(os.Stdout)

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}
}

func convert2SSA() string {
    return "hoge"
}