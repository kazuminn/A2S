package main

import (
	//"golang.org/x/tools/go/ssa"
	//"golang.org/x/tools/go/ssa/ssautil"
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	fileName := "add"
	CusSsa, err := getCustomSSA(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	OrgSsa, err := getOrgSSA(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if !reflect.DeepEqual(OrgSsa, CusSsa) {
		o := pp.Sprintf("%v", OrgSsa)
		c := pp.Sprintf("%v", CusSsa)
		t.Errorf("expected %s, actual %s", o, c)
	}
}
