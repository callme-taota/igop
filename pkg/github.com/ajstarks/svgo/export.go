// export by github.com/goplus/igop/cmd/qexp

package svg

import (
	q "github.com/ajstarks/svgo"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "svg",
		Path: "github.com/ajstarks/svgo",
		Deps: map[string]string{
			"encoding/xml": "xml",
			"fmt":          "fmt",
			"io":           "io",
			"strings":      "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Filterspec": reflect.TypeOf((*q.Filterspec)(nil)).Elem(),
			"Offcolor":   reflect.TypeOf((*q.Offcolor)(nil)).Elem(),
			"SVG":        reflect.TypeOf((*q.SVG)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
