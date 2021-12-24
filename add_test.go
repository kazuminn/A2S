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
	fileName := "add"
	ssa, err := loadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	ssaOrg, err := getSSAFromGo(fileName)
	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(ssaOrg, ssa) {
		e := pp.Sprintf("%v", ssaOrg)
		a := pp.Sprintf("%v", ssa)
		t.Errorf("expected %s, actual %s", e, a)
	}
}
