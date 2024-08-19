// export by github.com/goplus/igop/cmd/qexp

package audio

import (
	q "github.com/qiniu/audio"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "audio",
		Path: "github.com/qiniu/audio",
		Deps: map[string]string{
			"errors":                    "errors",
			"github.com/qiniu/x/bufiox": "bufiox",
			"io":                        "io",
			"sync":                      "sync",
			"sync/atomic":               "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"Decoded": reflect.TypeOf((*q.Decoded)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"DecodeConfigFunc": reflect.TypeOf((*q.DecodeConfigFunc)(nil)).Elem(),
			"DecodeFunc":       reflect.TypeOf((*q.DecodeFunc)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrFormat": reflect.ValueOf(&q.ErrFormat),
		},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"DecodeConfig":   reflect.ValueOf(q.DecodeConfig),
			"RegisterFormat": reflect.ValueOf(q.RegisterFormat),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
