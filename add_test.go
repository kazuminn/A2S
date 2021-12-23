package main

import (
	//"golang.org/x/tools/go/ssa"
	//"golang.org/x/tools/go/ssa/ssautil"
	"fmt"

	"github.com/k0kubun/pp"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	r, err := loadFile("add")
	if err != nil {
		fmt.Println(err)
	}
	ssa, _ := getSSAFromGo("add")
	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(ssa, r) {
		e := pp.Sprintf("%v", ssa)
		a := pp.Sprintf("%v", r)
		t.Errorf("expected %s, actual %s", e, a)
	}
}
