// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package binary

import (
	q "encoding/binary"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "binary",
		Path: "encoding/binary",
		Deps: map[string]string{
			"errors":  "errors",
			"io":      "io",
			"math":    "math",
			"reflect": "reflect",
			"sync":    "sync",
		},
		Interfaces: map[string]reflect.Type{
			"ByteOrder": reflect.TypeOf((*q.ByteOrder)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"BigEndian":    reflect.ValueOf(&q.BigEndian),
			"LittleEndian": reflect.ValueOf(&q.LittleEndian),
		},
		Funcs: map[string]reflect.Value{
			"PutUvarint":  reflect.ValueOf(q.PutUvarint),
			"PutVarint":   reflect.ValueOf(q.PutVarint),
			"Read":        reflect.ValueOf(q.Read),
			"ReadUvarint": reflect.ValueOf(q.ReadUvarint),
			"ReadVarint":  reflect.ValueOf(q.ReadVarint),
			"Size":        reflect.ValueOf(q.Size),
			"Uvarint":     reflect.ValueOf(q.Uvarint),
			"Varint":      reflect.ValueOf(q.Varint),
			"Write":       reflect.ValueOf(q.Write),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"MaxVarintLen16": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen16))},
			"MaxVarintLen32": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen32))},
			"MaxVarintLen64": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen64))},
		},
	})
}
