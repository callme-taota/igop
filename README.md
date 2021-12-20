# gossa - Golang SSA interpreter

[![Go1.16](https://github.com/goplus/gossa/workflows/Go1.16/badge.svg)](https://github.com/goplus/gossa/actions?query=workflow%3AGo1.16)

**Go1.17**

set env

`GOEXPERIMENT=noregabi`

### gossa command line
```
go get -u github.com/goplus/gossa/cmd/gossa
```

Commands
```
gossa run         # interpret package
gossa test        # test package
```

### gossa package

**run go source**
```
package main

import (
	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func main() {
	_, err := gossa.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}

```

**run gop source**
```
package main

import (
	"github.com/goplus/gossa"
	"github.com/goplus/gossa/gopbuild"
	_ "github.com/goplus/gossa/pkg"
)

var source = `
println "Hello, Go+"
`

func main() {
	ctx := gossa.NewContext(0)
	data, err := gopbuild.BuildFile(ctx, "main.gop", source)
	if err != nil {
		panic(err)
	}
	_, err = gossa.RunFile("main.go", data, nil, 0)
	if err != nil {
		panic(err)
	}
}
```