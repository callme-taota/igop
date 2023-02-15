// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package gosym

import (
	q "debug/gosym"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gosym",
		Path: "debug/gosym",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"fmt":             "fmt",
			"sort":            "sort",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"DecodingError":    reflect.TypeOf((*q.DecodingError)(nil)).Elem(),
			"Func":             reflect.TypeOf((*q.Func)(nil)).Elem(),
			"LineTable":        reflect.TypeOf((*q.LineTable)(nil)).Elem(),
			"Obj":              reflect.TypeOf((*q.Obj)(nil)).Elem(),
			"Sym":              reflect.TypeOf((*q.Sym)(nil)).Elem(),
			"Table":            reflect.TypeOf((*q.Table)(nil)).Elem(),
			"UnknownFileError": reflect.TypeOf((*q.UnknownFileError)(nil)).Elem(),
			"UnknownLineError": reflect.TypeOf((*q.UnknownLineError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewLineTable": reflect.ValueOf(q.NewLineTable),
			"NewTable":     reflect.ValueOf(q.NewTable),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
