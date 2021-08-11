// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_go116_linux, typList_go116_linux)
}

var extMap_go116_linux = map[string]interface{}{
	"syscall.AllThreadsSyscall":  syscall.AllThreadsSyscall,
	"syscall.AllThreadsSyscall6": syscall.AllThreadsSyscall6,
}

var typList_go116_linux = []interface{}{}
