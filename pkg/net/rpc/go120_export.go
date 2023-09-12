// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package rpc

import (
	q "net/rpc"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rpc",
		Path: "net/rpc",
		Deps: map[string]string{
			"bufio":         "bufio",
			"encoding/gob":  "gob",
			"errors":        "errors",
			"fmt":           "fmt",
			"go/token":      "token",
			"html/template": "template",
			"io":            "io",
			"log":           "log",
			"net":           "net",
			"net/http":      "http",
			"reflect":       "reflect",
			"sort":          "sort",
			"strings":       "strings",
			"sync":          "sync",
		},
		Interfaces: map[string]reflect.Type{
			"ClientCodec": reflect.TypeOf((*q.ClientCodec)(nil)).Elem(),
			"ServerCodec": reflect.TypeOf((*q.ServerCodec)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Call":        reflect.TypeOf((*q.Call)(nil)).Elem(),
			"Client":      reflect.TypeOf((*q.Client)(nil)).Elem(),
			"Request":     reflect.TypeOf((*q.Request)(nil)).Elem(),
			"Response":    reflect.TypeOf((*q.Response)(nil)).Elem(),
			"Server":      reflect.TypeOf((*q.Server)(nil)).Elem(),
			"ServerError": reflect.TypeOf((*q.ServerError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DefaultServer": reflect.ValueOf(&q.DefaultServer),
			"ErrShutdown":   reflect.ValueOf(&q.ErrShutdown),
		},
		Funcs: map[string]reflect.Value{
			"Accept":             reflect.ValueOf(q.Accept),
			"Dial":               reflect.ValueOf(q.Dial),
			"DialHTTP":           reflect.ValueOf(q.DialHTTP),
			"DialHTTPPath":       reflect.ValueOf(q.DialHTTPPath),
			"HandleHTTP":         reflect.ValueOf(q.HandleHTTP),
			"NewClient":          reflect.ValueOf(q.NewClient),
			"NewClientWithCodec": reflect.ValueOf(q.NewClientWithCodec),
			"NewServer":          reflect.ValueOf(q.NewServer),
			"Register":           reflect.ValueOf(q.Register),
			"RegisterName":       reflect.ValueOf(q.RegisterName),
			"ServeCodec":         reflect.ValueOf(q.ServeCodec),
			"ServeConn":          reflect.ValueOf(q.ServeConn),
			"ServeRequest":       reflect.ValueOf(q.ServeRequest),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"DefaultDebugPath": {"untyped string", constant.MakeString(string(q.DefaultDebugPath))},
			"DefaultRPCPath":   {"untyped string", constant.MakeString(string(q.DefaultRPCPath))},
		},
	})
}
