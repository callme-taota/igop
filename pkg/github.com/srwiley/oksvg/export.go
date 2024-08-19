// export by github.com/goplus/igop/cmd/qexp

package oksvg

import (
	q "github.com/srwiley/oksvg"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "oksvg",
		Path: "github.com/srwiley/oksvg",
		Deps: map[string]string{
			"encoding/xml":                  "xml",
			"errors":                        "errors",
			"fmt":                           "fmt",
			"github.com/srwiley/rasterx":    "rasterx",
			"golang.org/x/image/colornames": "colornames",
			"golang.org/x/image/math/fixed": "fixed",
			"golang.org/x/net/html/charset": "charset",
			"image/color":                   "color",
			"io":                            "io",
			"log":                           "log",
			"math":                          "math",
			"os":                            "os",
			"strconv":                       "strconv",
			"strings":                       "strings",
			"unicode":                       "unicode",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ErrorMode":  reflect.TypeOf((*q.ErrorMode)(nil)).Elem(),
			"IconCursor": reflect.TypeOf((*q.IconCursor)(nil)).Elem(),
			"PathCursor": reflect.TypeOf((*q.PathCursor)(nil)).Elem(),
			"PathStyle":  reflect.TypeOf((*q.PathStyle)(nil)).Elem(),
			"SvgIcon":    reflect.TypeOf((*q.SvgIcon)(nil)).Elem(),
			"SvgPath":    reflect.TypeOf((*q.SvgPath)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DefaultStyle": reflect.ValueOf(&q.DefaultStyle),
		},
		Funcs: map[string]reflect.Value{
			"ParseSVGColor":    reflect.ValueOf(q.ParseSVGColor),
			"ParseSVGColorNum": reflect.ValueOf(q.ParseSVGColorNum),
			"ReadIcon":         reflect.ValueOf(q.ReadIcon),
			"ReadIconStream":   reflect.ValueOf(q.ReadIconStream),
		},
		TypedConsts: map[string]igop.TypedConst{
			"IgnoreErrorMode": {reflect.TypeOf(q.IgnoreErrorMode), constant.MakeInt64(int64(q.IgnoreErrorMode))},
			"StrictErrorMode": {reflect.TypeOf(q.StrictErrorMode), constant.MakeInt64(int64(q.StrictErrorMode))},
			"WarnErrorMode":   {reflect.TypeOf(q.WarnErrorMode), constant.MakeInt64(int64(q.WarnErrorMode))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
