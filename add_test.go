package main

import (
	//"golang.org/x/tools/go/ssa"
	//"golang.org/x/tools/go/ssa/ssautil"
	"reflect"
	"testing"
	"github.com/k0kubun/pp"
)

func TestAdd(t *testing.T) {
	r := convert2SSA()

	if !reflect.DeepEqual(test[0].pass, r) {
		e := pp.Sprintf("%v", test)
		a := pp.Sprintf("%v", r)
		t.Errorf("expected %s, actual %s", e, a)
	}
}

var test = []struct {
	pass string
}{
	{
	"hoge",
	},
}

/*
	pass: &ssa.Function{
  name:   "add",
  object: &types.Func{
    object: types.object{
      parent: &types.Scope{
        parent: &types.Scope{
          parent:   (*types.Scope)(nil),
          children: []*types.Scope{},
          elems:    map[string]types.Object{
            "uintptr": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uintptr",
                typ:    &types.Basic{
                  kind: 12,
                  info: 6,
                  name: "uintptr",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "int16": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "int16",
                typ:    &types.Basic{
                  kind: 4,
                  info: 2,
                  name: "int16",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "panic": &types.Builtin{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "panic",
                typ:    &types.Basic{
                  kind: 0,
                  info: 0,
                  name: "invalid type",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 10,
            },
            "string": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "string",
                typ:    &types.Basic{
                  kind: 17,
                  info: 32,
                  name: "string",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "uint8": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uint8",
                typ:    &types.Basic{
                  kind: 8,
                  info: 6,
                  name: "uint8",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "iota": &types.Const{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "iota",
                typ:    &types.Basic{
                  kind: 20,
                  info: 66,
                  name: "untyped int",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              val: 0,
            },
            "cap": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "cap",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 1,
            },
            "complex": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "complex",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 3,
            },
            "int8": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "int8",
                typ:    &types.Basic{
                  kind: 3,
                  info: 2,
                  name: "int8",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "complex128": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "complex128",
                typ:    &types.Basic{
                  kind: 16,
                  info: 16,
                  name: "complex128",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "close": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "close",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 2,
            },
            "delete": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "delete",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 5,
            },
            "float64": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "float64",
                typ:    &types.Basic{
                  kind: 14,
                  info: 8,
                  name: "float64",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "uint": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uint",
                typ:    &types.Basic{
                  kind: 7,
                  info: 6,
                  name: "uint",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "uint16": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uint16",
                typ:    &types.Basic{
                  kind: 9,
                  info: 6,
                  name: "uint16",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "uint32": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uint32",
                typ:    &types.Basic{
                  kind: 10,
                  info: 6,
                  name: "uint32",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "uint64": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "uint64",
                typ:    &types.Basic{
                  kind: 11,
                  info: 6,
                  name: "uint64",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "true": &types.Const{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "true",
                typ:    &types.Basic{
                  kind: 19,
                  info: 65,
                  name: "untyped bool",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              val: true,
            },
            "false": &types.Const{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "false",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              val: false,
            },
            "copy": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "copy",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 4,
            },
            "int": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "int",
                typ:    &types.Basic{
                  kind: 2,
                  info: 2,
                  name: "int",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "real": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "real",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 13,
            },
            "len": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "len",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 7,
            },
            "error": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "error",
                typ:    &types.Named{
                  check:      (*types.Checker)(nil),
                  info:       0x0,
                  obj:        &types.TypeName{...},
                  orig:       nil,
                  underlying: &types.Interface{
                    methods: []*types.Func{
                      &types.Func{
                        object: types.object{
                          parent: (*types.Scope)(nil),
                          pos:    0,
                          pkg:    (*types.Package)(nil),
                          name:   "Error",
                          typ:    &types.Signature{
                            rparams: []*types.TypeName{},
                            tparams: []*types.TypeName{},
                            scope:   (*types.Scope)(nil),
                            recv:    &types.Var{
                              object: types.object{
                                parent:    (*types.Scope)(nil),
                                pos:       0,
                                pkg:       (*types.Package)(nil),
                                name:      "",
                                typ:       &types.Named{...},
                                order_:    0x00000000,
                                color_:    0x00000001,
                                scopePos_: 0,
                              },
                              embedded: false,
                              isField:  false,
                              used:     false,
                            },
                            params:  (*types.Tuple)(nil),
                            results: &types.Tuple{
                              vars: []*types.Var{
                                &types.Var{
                                  object: types.object{
                                    parent:    (*types.Scope)(nil),
                                    pos:       0,
                                    pkg:       (*types.Package)(nil),
                                    name:      "",
                                    typ:       &types.Basic{...},
                                    order_:    0x00000000,
                                    color_:    0x00000001,
                                    scopePos_: 0,
                                  },
                                  embedded: false,
                                  isField:  false,
                                  used:     false,
                                },
                              },
                            },
                            variadic: false,
                          },
                          order_:    0x00000000,
                          color_:    0x00000001,
                          scopePos_: 0,
                        },
                        hasPtrRecv: false,
                      },
                    },
                    types:      nil,
                    embeddeds:  []types.Type{},
                    allMethods: []*types.Func{
                      &types.Func{...},
                    },
                    allTypes: nil,
                    obj:      nil,
                  },
                  tparams: []*types.TypeName{},
                  targs:   []types.Type{},
                  methods: []*types.Func{},
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "imag": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "imag",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 6,
            },
            "new": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "new",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 9,
            },
            "println": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "println",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 12,
            },
            "rune": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "rune",
                typ:    &types.Basic{
                  kind: 5,
                  info: 2,
                  name: "rune",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "int64": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "int64",
                typ:    &types.Basic{
                  kind: 6,
                  info: 2,
                  name: "int64",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "complex64": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "complex64",
                typ:    &types.Basic{
                  kind: 15,
                  info: 16,
                  name: "complex64",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "append": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "append",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 0,
            },
            "make": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "make",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 8,
            },
            "print": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "print",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 11,
            },
            "bool": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "bool",
                typ:    &types.Basic{
                  kind: 1,
                  info: 1,
                  name: "bool",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "float32": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "float32",
                typ:    &types.Basic{
                  kind: 13,
                  info: 8,
                  name: "float32",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "byte": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "byte",
                typ:    &types.Basic{
                  kind: 8,
                  info: 6,
                  name: "byte",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "nil": &types.Nil{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "nil",
                typ:    &types.Basic{
                  kind: 25,
                  info: 64,
                  name: "untyped nil",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
            "recover": &types.Builtin{
              object: types.object{
                parent:    &types.Scope{...},
                pos:       0,
                pkg:       (*types.Package)(nil),
                name:      "recover",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              id: 14,
            },
            "int32": &types.TypeName{
              object: types.object{
                parent: &types.Scope{...},
                pos:    0,
                pkg:    (*types.Package)(nil),
                name:   "int32",
                typ:    &types.Basic{
                  kind: 5,
                  info: 2,
                  name: "int32",
                },
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
            },
          },
          pos:     0,
          end:     0,
          comment: "universe",
          isFunc:  false,
        },
        children: []*types.Scope{
          &types.Scope{
            parent:   &types.Scope{...},
            children: []*types.Scope{
              &types.Scope{
                parent:   &types.Scope{...},
                children: []*types.Scope{},
                elems:    map[string]types.Object{},
                pos:      32,
                end:      49,
                comment:  "function",
                isFunc:   true,
              },
            },
            elems:   map[string]types.Object{},
            pos:     1,
            end:     50,
            comment: "C:\\Users\\warug\\a2s\\sample\\sample.go",
            isFunc:  false,
          },
        },
        elems: map[string]types.Object{
          "add": &types.Func{...},
        },
        pos:     0,
        end:     0,
        comment: "package \"hello\"",
        isFunc:  false,
      },
      pos: 22,
      pkg: &types.Package{
        path:     "hello",
        name:     "sample",
        scope:    &types.Scope{...},
        complete: true,
        imports:  []*types.Package{},
        fake:     false,
        cgo:      false,
      },
      name: "add",
      typ:  &types.Signature{
        rparams: []*types.TypeName{},
        tparams: []*types.TypeName{},
        scope:   &types.Scope{...},
        recv:    (*types.Var)(nil),
        params:  (*types.Tuple)(nil),
        results: &types.Tuple{
          vars: []*types.Var{
            &types.Var{
              object: types.object{
                parent:    (*types.Scope)(nil),
                pos:       28,
                pkg:       &types.Package{...},
                name:      "",
                typ:       &types.Basic{...},
                order_:    0x00000000,
                color_:    0x00000001,
                scopePos_: 0,
              },
              embedded: false,
              isField:  false,
              used:     true,
            },
          },
        },
        variadic: false,
      },
      order_:    0x00000001,
      color_:    0x00000001,
      scopePos_: 0,
    },
    hasPtrRecv: false,
  },
  method:    (*types.Selection)(nil),
  Signature: &types.Signature{...},
  pos:       22,
  Synthetic: "",
  syntax:    ssa.extentNode{
    17,
    49,
  },
  parent: (*ssa.Function)(nil),
  Pkg:    &ssa.Package{
    Prog: &ssa.Program{
      Fset: &token.FileSet{
        mutex: sync.RWMutex{
          w: sync.Mutex{
            state: 0,
            sema:  0x00000000,
          },
          writerSem:   0x00000000,
          readerSem:   0x00000000,
          readerCount: 0,
          readerWait:  0,
        },
        base:  51,
        files: []*token.File{
          &token.File{
            set:   &token.FileSet{...},
            name:  "C:\\Users\\warug\\a2s\\sample\\sample.go",
            base:  1,
            size:  49,
            mutex: sync.Mutex{
              state: 0,
              sema:  0x00000000,
            },
            lines: []int{
              0,
              15,
              16,
              33,
              47,
            },
            infos: []token.lineInfo{},
          },
        },
        last: &token.File{...},
      },
      imported: map[string]*ssa.Package{},
      packages: map[*types.Package]*ssa.Package{
        &types.Package{...}: &ssa.Package{...},
      },
      mode:       0x8,
      MethodSets: typeutil.MethodSetCache{
        mu: sync.Mutex{
          state: 0,
          sema:  0x00000000,
        },
        named:  map[*types.Named]struct { value *types.MethodSet; pointer *types.MethodSet }{},
        others: map[types.Type]*types.MethodSet{},
      },
      methodsMu: sync.Mutex{
        state: 0,
        sema:  0x00000000,
      },
      methodSets: typeutil.Map{
        hasher: typeutil.Hasher{
          memo:       map[types.Type]uint32{},
          ptrMap:     map[interface {}]uint32{},
          sigTParams: (*typeparams.TypeParamList)(nil),
        },
        table:  map[uint32][]typeutil.entry{},
        length: 0,
      },
      runtimeTypes: typeutil.Map{
        hasher: typeutil.Hasher{
          memo:       map[types.Type]uint32{},
          ptrMap:     map[interface {}]uint32{},
          sigTParams: (*typeparams.TypeParamList)(nil),
        },
        table:  map[uint32][]typeutil.entry{},
        length: 0,
      },
      canon: typeutil.Map{
        hasher: typeutil.Hasher{
          memo:       map[types.Type]uint32{},
          ptrMap:     map[interface {}]uint32{},
          sigTParams: (*typeparams.TypeParamList)(nil),
        },
        table:  map[uint32][]typeutil.entry{},
        length: 0,
      },
      bounds: map[*types.Func]*ssa.Function{},
      thunks: map[ssa.selectionKey]*ssa.Function{},
    },
    Pkg:     &types.Package{...},
    Members: map[string]ssa.Member{
      "init": &ssa.Function{
        name:      "init",
        object:    nil,
        method:    (*types.Selection)(nil),
        Signature: &types.Signature{
          rparams:  []*types.TypeName{},
          tparams:  []*types.TypeName{},
          scope:    (*types.Scope)(nil),
          recv:     (*types.Var)(nil),
          params:   (*types.Tuple)(nil),
          results:  (*types.Tuple)(nil),
          variadic: false,
        },
        pos:       0,
        Synthetic: "package initializer",
        syntax:    nil,
        parent:    (*ssa.Function)(nil),
        Pkg:       &ssa.Package{...},
        Prog:      &ssa.Program{...},
        Params:    []*ssa.Parameter{},
        FreeVars:  []*ssa.FreeVar{},
        Locals:    []*ssa.Alloc{},
        Blocks:    []*ssa.BasicBlock{
          &ssa.BasicBlock{
            Index:   0,
            Comment: "entry",
            parent:  &ssa.Function{...},
            Instrs:  []ssa.Instruction{
              &ssa.UnOp{
                register: ssa.register{
                  anInstruction: ssa.anInstruction{
                    block: &ssa.BasicBlock{...},
                  },
                  num:       0,
                  typ:       &types.Basic{...},
                  pos:       0,
                  referrers: []ssa.Instruction{
                    &ssa.If{
                      anInstruction: ssa.anInstruction{
                        block: &ssa.BasicBlock{...},
                      },
                      Cond: &ssa.UnOp{...},
                    },
                  },
                },
                Op: 14,
                X:  &ssa.Global{
                  name:   "init$guard",
                  object: nil,
                  typ:    &types.Pointer{
                    base: &types.Basic{...},
                  },
                  pos: 0,
                  Pkg: &ssa.Package{...},
                },
                CommaOk: false,
              },
              &ssa.If{...},
            },
            Preds: []*ssa.BasicBlock{},
            Succs: []*ssa.BasicBlock{
              &ssa.BasicBlock{
                Index:   2,
                Comment: "init.done",
                parent:  &ssa.Function{...},
                Instrs:  []ssa.Instruction{
                  &ssa.Return{
                    anInstruction: ssa.anInstruction{
                      block: &ssa.BasicBlock{...},
                    },
                    Results: []ssa.Value{},
                    pos:     0,
                  },
                },
                Preds: []*ssa.BasicBlock{
                  &ssa.BasicBlock{...},
                  &ssa.BasicBlock{
                    Index:   1,
                    Comment: "init.start",
                    parent:  &ssa.Function{...},
                    Instrs:  []ssa.Instruction{
                      &ssa.Store{
                        anInstruction: ssa.anInstruction{
                          block: &ssa.BasicBlock{...},
                        },
                        Addr: &ssa.Global{...},
                        Val:  &ssa.Const{
                          typ:   &types.Basic{...},
                          Value: true,
                        },
                        pos: 0,
                      },
                      &ssa.Jump{
                        anInstruction: ssa.anInstruction{
                          block: &ssa.BasicBlock{...},
                        },
                      },
                    },
                    Preds: []*ssa.BasicBlock{
                      &ssa.BasicBlock{...},
                    },
                    Succs: []*ssa.BasicBlock{
                      &ssa.BasicBlock{...},
                    },
                    succs2: [2]*ssa.BasicBlock{
                      &ssa.BasicBlock{...},
                      (*ssa.BasicBlock)(nil),
                    },
                    dom: ssa.domInfo{
                      idom:     &ssa.BasicBlock{...},
                      children: []*ssa.BasicBlock{},
                      pre:      2,
                      post:     1,
                    },
                    gaps:      0,
                    rundefers: 0,
                  },
                },
                Succs:  []*ssa.BasicBlock{},
                succs2: [2]*ssa.BasicBlock{
                  (*ssa.BasicBlock)(nil),
                  (*ssa.BasicBlock)(nil),
                },
                dom: ssa.domInfo{
                  idom:     &ssa.BasicBlock{...},
                  children: []*ssa.BasicBlock{},
                  pre:      1,
                  post:     0,
                },
                gaps:      0,
                rundefers: 0,
              },
              &ssa.BasicBlock{...},
            },
            succs2: [2]*ssa.BasicBlock{
              &ssa.BasicBlock{...},
              &ssa.BasicBlock{...},
            },
            dom: ssa.domInfo{
              idom:     (*ssa.BasicBlock)(nil),
              children: []*ssa.BasicBlock{
                &ssa.BasicBlock{...},
                &ssa.BasicBlock{...},
              },
              pre:  0,
              post: 2,
            },
            gaps:      0,
            rundefers: 0,
          },
          &ssa.BasicBlock{...},
          &ssa.BasicBlock{...},
        },
        Recover:      (*ssa.BasicBlock)(nil),
        AnonFuncs:    []*ssa.Function{},
        referrers:    []ssa.Instruction{},
        currentBlock: (*ssa.BasicBlock)(nil),
        objects:      map[types.Object]ssa.Value{},
        namedResults: []*ssa.Alloc{},
        targets:      (*ssa.targets)(nil),
        lblocks:      map[*ast.Object]*ssa.lblock{},
      },
      "add":        &ssa.Function{...},
      "init$guard": &ssa.Global{...},
    },
    values: map[types.Object]ssa.Value{
      &types.Func{...}: &ssa.Function{...},
    },
    init:      &ssa.Function{...},
    debug:     false,
    buildOnce: sync.Once{
      done: 0x00000001,
      m:    sync.Mutex{
        state: 0,
        sema:  0x00000000,
      },
    },
    ninit: 0,
    info:  (*types.Info)(nil),
    files: []*ast.File{
      &ast.File{
        Doc:     (*ast.CommentGroup)(nil),
        Package: 1,
        Name:    &ast.Ident{
          NamePos: 9,
          Name:    "sample",
          Obj:     (*ast.Object)(nil),
        },
        Decls: []ast.Decl{
          &ast.FuncDecl{
            Doc:  (*ast.CommentGroup)(nil),
            Recv: (*ast.FieldList)(nil),
            Name: &ast.Ident{
              NamePos: 22,
              Name:    "add",
              Obj:     &ast.Object{
                Kind: 5,
                Name: "add",
                Decl: &ast.FuncDecl{...},
                Data: nil,
                Type: nil,
              },
            },
            Type: &ast.FuncType{
              Func:   17,
              Params: &ast.FieldList{
                Opening: 25,
                List:    []*ast.Field{},
                Closing: 26,
              },
              Results: &ast.FieldList{
                Opening: 0,
                List:    []*ast.Field{
                  &ast.Field{
                    Doc:   (*ast.CommentGroup)(nil),
                    Names: []*ast.Ident{},
                    Type:  &ast.Ident{
                      NamePos: 28,
                      Name:    "int",
                      Obj:     (*ast.Object)(nil),
                    },
                    Tag:     (*ast.BasicLit)(nil),
                    Comment: (*ast.CommentGroup)(nil),
                  },
                },
                Closing: 0,
              },
            },
            Body: &ast.BlockStmt{
              Lbrace: 32,
              List:   []ast.Stmt{
                &ast.ReturnStmt{
                  Return:  35,
                  Results: []ast.Expr{
                    &ast.BinaryExpr{
                      X: &ast.BasicLit{
                        ValuePos: 42,
                        Kind:     5,
                        Value:    "1",
                      },
                      OpPos: 44,
                      Op:    12,
                      Y:     &ast.BasicLit{
                        ValuePos: 46,
                        Kind:     5,
                        Value:    "2",
                      },
                    },
                  },
                },
              },
              Rbrace: 48,
            },
          },
        },
        Scope: &ast.Scope{
          Outer:   (*ast.Scope)(nil),
          Objects: map[string]*ast.Object{
            "add": &ast.Object{...},
          },
        },
        Imports:    []*ast.ImportSpec{},
        Unresolved: []*ast.Ident{
          &ast.Ident{...},
        },
        Comments: []*ast.CommentGroup{},
      },
    },
  },
  Prog:     &ssa.Program{...},
  Params:   []*ssa.Parameter{},
  FreeVars: []*ssa.FreeVar{},
  Locals:   []*ssa.Alloc{},
  Blocks:   []*ssa.BasicBlock{
    &ssa.BasicBlock{
      Index:   0,
      Comment: "entry",
      parent:  &ssa.Function{...},
      Instrs:  []ssa.Instruction{
        &ssa.Return{
          anInstruction: ssa.anInstruction{
            block: &ssa.BasicBlock{...},
          },
          Results: []ssa.Value{
            &ssa.Const{
              typ:   &types.Basic{...},
              Value: 3,
            },
          },
          pos: 35,
        },
      },
      Preds:  []*ssa.BasicBlock{},
      Succs:  []*ssa.BasicBlock{},
      succs2: [2]*ssa.BasicBlock{
        (*ssa.BasicBlock)(nil),
        (*ssa.BasicBlock)(nil),
      },
      dom: ssa.domInfo{
        idom:     (*ssa.BasicBlock)(nil),
        children: []*ssa.BasicBlock{},
        pre:      0,
        post:     0,
      },
      gaps:      0,
      rundefers: 1,
    },
  },
  Recover:      (*ssa.BasicBlock)(nil),
  AnonFuncs:    []*ssa.Function{},
  referrers:    []ssa.Instruction{},
  currentBlock: (*ssa.BasicBlock)(nil),
  objects:      map[types.Object]ssa.Value{},
  namedResults: []*ssa.Alloc{},
  targets:      (*ssa.targets)(nil),
  lblocks:      map[*ast.Object]*ssa.lblock{},
}
*/
