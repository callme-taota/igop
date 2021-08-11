// export by github.com/goplus/interp/cmd/qexp

package cgi

import (
	"net/http/cgi"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/http/cgi", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/http/cgi.Handler).ServeHTTP": (*cgi.Handler).ServeHTTP,
	"net/http/cgi.Request":              cgi.Request,
	"net/http/cgi.RequestFromMap":       cgi.RequestFromMap,
	"net/http/cgi.Serve":                cgi.Serve,
}

var typList = []interface{}{
	(*cgi.Handler)(nil),
}
