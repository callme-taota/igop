package igop_test

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/fmt"
)

func TestReplExpr(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`a+b`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[3 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplImports(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`import "fmt"`,
		`c := fmt.Sprintf("%v-%v",a,b)`,
		`c`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[1-2 string]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplClosure(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`fn := func() int { return a }`,
		`d := fn()+b`,
		`d`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[3 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplVar(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`var c = 100`,
		`fn := func() int { return a+c }`,
		`d := fn()+b`,
		`d`,
		`c`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[103 int]`,
		`[100 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplType(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`type T struct {
	X int
	Y int
}`,
		`v1 := &T{10,20}`,
		`import "fmt"`,
		`v1`,
		`func (t *T) String() string {
	return fmt.Sprintf("%v-%v",t.X,t.Y)
}`,
		`v2 := &T{10,20}`,
		`v2`,
	}
	result := []string{
		`-`,
		`-`,
		`-`,
		`[&{10 20} *main.T]`,
		`-`,
		`-`,
		`[10-20 *main.T]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFunc(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`a`,
		`fmt.Println(a)`,
		`s := fmt.Sprint(a)`,
		`s`,
		`fmt.Println(s)`,
	}
	result := []string{
		`[]`,
		`-`,
		`[hello string]`,
		`[6 int <nil> error]`,
		`-`,
		`[hello string]`,
		`[6 int <nil> error]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplTok(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`var c int`,
		`const d = 100`,
		`func test(v int) {}`,
		`fmt.Println(a)`,
		`type T int`,
	}
	toks := []token.Token{
		token.IDENT,
		token.IMPORT,
		token.VAR,
		token.CONST,
		token.FUNC,
		token.IDENT,
		token.TYPE,
	}
	for i, expr := range list {
		tok, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if tok != toks[i] {
			t.Fatalf("expr:%v token:%v", expr, tok)
		}
	}
}

func TestReplInit(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`var i int = 10`,
		`var j int = 20`,
		`func init() { i++ }`,
		`func init() { j++ }`,
		`var c int`,
		`i`,
		`j`,
	}
	result := []string{
		`-`,
		`-`,
		`-`,
		`-`,
		`-`,
		`[11 int]`,
		`[21 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFunLit(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`var a int`,
		`func(x int){ a = x }(100)`,
		`a`,
	}
	result := []string{
		`-`,
		`-`,
		`[100 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}
