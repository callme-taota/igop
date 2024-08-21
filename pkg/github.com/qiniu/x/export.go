// export by github.com/goplus/igop/cmd/qexp

package x

import (
	_ "github.com/qiniu/x"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "x",
		Path: "github.com/qiniu/x",
		Deps: map[string]string{},
	})
}
