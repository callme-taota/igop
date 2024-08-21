// export by github.com/goplus/igop/cmd/qexp

package rasterx

import (
	q "github.com/srwiley/rasterx"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rasterx",
		Path: "github.com/srwiley/rasterx",
		Deps: map[string]string{
			"fmt":                           "fmt",
			"golang.org/x/image/math/fixed": "fixed",
			"golang.org/x/image/vector":     "vector",
			"image":                         "image",
			"image/color":                   "color",
			"image/draw":                    "draw",
			"math":                          "math",
			"sort":                          "sort",
		},
		Interfaces: map[string]reflect.Type{
			"Adder":   reflect.TypeOf((*q.Adder)(nil)).Elem(),
			"Rasterx": reflect.TypeOf((*q.Rasterx)(nil)).Elem(),
			"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"C2Point":        reflect.TypeOf((*q.C2Point)(nil)).Elem(),
			"CapFunc":        reflect.TypeOf((*q.CapFunc)(nil)).Elem(),
			"ClipImage":      reflect.TypeOf((*q.ClipImage)(nil)).Elem(),
			"ColorFunc":      reflect.TypeOf((*q.ColorFunc)(nil)).Elem(),
			"ColorFuncImage": reflect.TypeOf((*q.ColorFuncImage)(nil)).Elem(),
			"Dasher":         reflect.TypeOf((*q.Dasher)(nil)).Elem(),
			"Filler":         reflect.TypeOf((*q.Filler)(nil)).Elem(),
			"GapFunc":        reflect.TypeOf((*q.GapFunc)(nil)).Elem(),
			"GradStop":       reflect.TypeOf((*q.GradStop)(nil)).Elem(),
			"Gradient":       reflect.TypeOf((*q.Gradient)(nil)).Elem(),
			"GradientUnits":  reflect.TypeOf((*q.GradientUnits)(nil)).Elem(),
			"JoinMode":       reflect.TypeOf((*q.JoinMode)(nil)).Elem(),
			"Matrix2D":       reflect.TypeOf((*q.Matrix2D)(nil)).Elem(),
			"MatrixAdder":    reflect.TypeOf((*q.MatrixAdder)(nil)).Elem(),
			"Path":           reflect.TypeOf((*q.Path)(nil)).Elem(),
			"PathCommand":    reflect.TypeOf((*q.PathCommand)(nil)).Elem(),
			"ScannerGV":      reflect.TypeOf((*q.ScannerGV)(nil)).Elem(),
			"SpreadMethod":   reflect.TypeOf((*q.SpreadMethod)(nil)).Elem(),
			"Stroker":        reflect.TypeOf((*q.Stroker)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ButtCap":      reflect.ValueOf(&q.ButtCap),
			"CubicCap":     reflect.ValueOf(&q.CubicCap),
			"CubicGap":     reflect.ValueOf(&q.CubicGap),
			"FlatGap":      reflect.ValueOf(&q.FlatGap),
			"Identity":     reflect.ValueOf(&q.Identity),
			"QuadraticCap": reflect.ValueOf(&q.QuadraticCap),
			"QuadraticGap": reflect.ValueOf(&q.QuadraticGap),
			"RoundCap":     reflect.ValueOf(&q.RoundCap),
			"RoundGap":     reflect.ValueOf(&q.RoundGap),
			"SquareCap":    reflect.ValueOf(&q.SquareCap),
		},
		Funcs: map[string]reflect.Value{
			"AddArc":                   reflect.ValueOf(q.AddArc),
			"AddCircle":                reflect.ValueOf(q.AddCircle),
			"AddEllipse":               reflect.ValueOf(q.AddEllipse),
			"AddRect":                  reflect.ValueOf(q.AddRect),
			"AddRoundRect":             reflect.ValueOf(q.AddRoundRect),
			"ApplyOpacity":             reflect.ValueOf(q.ApplyOpacity),
			"CalcIntersect":            reflect.ValueOf(q.CalcIntersect),
			"CircleCircleIntersection": reflect.ValueOf(q.CircleCircleIntersection),
			"ClosestPortside":          reflect.ValueOf(q.ClosestPortside),
			"CubeTo":                   reflect.ValueOf(q.CubeTo),
			"DotProd":                  reflect.ValueOf(q.DotProd),
			"FindEllipseCenter":        reflect.ValueOf(q.FindEllipseCenter),
			"GapToCap":                 reflect.ValueOf(q.GapToCap),
			"Invert":                   reflect.ValueOf(q.Invert),
			"Length":                   reflect.ValueOf(q.Length),
			"NewDasher":                reflect.ValueOf(q.NewDasher),
			"NewFiller":                reflect.ValueOf(q.NewFiller),
			"NewScannerGV":             reflect.ValueOf(q.NewScannerGV),
			"NewStroker":               reflect.ValueOf(q.NewStroker),
			"QuadTo":                   reflect.ValueOf(q.QuadTo),
			"RadCurvature":             reflect.ValueOf(q.RadCurvature),
			"RayCircleIntersection":    reflect.ValueOf(q.RayCircleIntersection),
			"RayCircleIntersectionF":   reflect.ValueOf(q.RayCircleIntersectionF),
			"ToFixedP":                 reflect.ValueOf(q.ToFixedP),
			"ToLength":                 reflect.ValueOf(q.ToLength),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Arc":               {reflect.TypeOf(q.Arc), constant.MakeInt64(int64(q.Arc))},
			"ArcClip":           {reflect.TypeOf(q.ArcClip), constant.MakeInt64(int64(q.ArcClip))},
			"Bevel":             {reflect.TypeOf(q.Bevel), constant.MakeInt64(int64(q.Bevel))},
			"MaxDx":             {reflect.TypeOf(q.MaxDx), constant.BinaryOp(constant.MakeFromLiteral("884279719003555", token.INT, 0), token.QUO, constant.MakeFromLiteral("2251799813685248", token.INT, 0))},
			"Miter":             {reflect.TypeOf(q.Miter), constant.MakeInt64(int64(q.Miter))},
			"MiterClip":         {reflect.TypeOf(q.MiterClip), constant.MakeInt64(int64(q.MiterClip))},
			"ObjectBoundingBox": {reflect.TypeOf(q.ObjectBoundingBox), constant.MakeInt64(int64(q.ObjectBoundingBox))},
			"PadSpread":         {reflect.TypeOf(q.PadSpread), constant.MakeInt64(int64(q.PadSpread))},
			"PathClose":         {reflect.TypeOf(q.PathClose), constant.MakeInt64(int64(q.PathClose))},
			"PathCubicTo":       {reflect.TypeOf(q.PathCubicTo), constant.MakeInt64(int64(q.PathCubicTo))},
			"PathLineTo":        {reflect.TypeOf(q.PathLineTo), constant.MakeInt64(int64(q.PathLineTo))},
			"PathMoveTo":        {reflect.TypeOf(q.PathMoveTo), constant.MakeInt64(int64(q.PathMoveTo))},
			"PathQuadTo":        {reflect.TypeOf(q.PathQuadTo), constant.MakeInt64(int64(q.PathQuadTo))},
			"ReflectSpread":     {reflect.TypeOf(q.ReflectSpread), constant.MakeInt64(int64(q.ReflectSpread))},
			"RepeatSpread":      {reflect.TypeOf(q.RepeatSpread), constant.MakeInt64(int64(q.RepeatSpread))},
			"Round":             {reflect.TypeOf(q.Round), constant.MakeInt64(int64(q.Round))},
			"UserSpaceOnUse":    {reflect.TypeOf(q.UserSpaceOnUse), constant.MakeInt64(int64(q.UserSpaceOnUse))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
