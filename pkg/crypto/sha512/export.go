// export by github.com/goplus/gossa/cmd/qexp

package sha512

import (
	q "crypto/sha512"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "sha512",
		Path: "crypto/sha512",
		Deps: map[string]string{
			"crypto":          "crypto",
			"encoding/binary": "binary",
			"errors":          "errors",
			"hash":            "hash",
			"internal/cpu":    "cpu",
			"math/bits":       "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":        reflect.ValueOf(q.New),
			"New384":     reflect.ValueOf(q.New384),
			"New512_224": reflect.ValueOf(q.New512_224),
			"New512_256": reflect.ValueOf(q.New512_256),
			"Sum384":     reflect.ValueOf(q.Sum384),
			"Sum512":     reflect.ValueOf(q.Sum512),
			"Sum512_224": reflect.ValueOf(q.Sum512_224),
			"Sum512_256": reflect.ValueOf(q.Sum512_256),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
			"Size224":   {"untyped int", constant.MakeInt64(int64(q.Size224))},
			"Size256":   {"untyped int", constant.MakeInt64(int64(q.Size256))},
			"Size384":   {"untyped int", constant.MakeInt64(int64(q.Size384))},
		},
	})
}
