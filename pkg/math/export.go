// export by github.com/goplus/gossa/cmd/qexp

package math

import (
	q "math"

	"go/constant"
	"go/token"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "math",
		Path: "math",
		Deps: map[string]string{
			"internal/cpu": "cpu",
			"math/bits":    "bits",
			"unsafe":       "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Abs":             reflect.ValueOf(q.Abs),
			"Acos":            reflect.ValueOf(q.Acos),
			"Acosh":           reflect.ValueOf(q.Acosh),
			"Asin":            reflect.ValueOf(q.Asin),
			"Asinh":           reflect.ValueOf(q.Asinh),
			"Atan":            reflect.ValueOf(q.Atan),
			"Atan2":           reflect.ValueOf(q.Atan2),
			"Atanh":           reflect.ValueOf(q.Atanh),
			"Cbrt":            reflect.ValueOf(q.Cbrt),
			"Ceil":            reflect.ValueOf(q.Ceil),
			"Copysign":        reflect.ValueOf(q.Copysign),
			"Cos":             reflect.ValueOf(q.Cos),
			"Cosh":            reflect.ValueOf(q.Cosh),
			"Dim":             reflect.ValueOf(q.Dim),
			"Erf":             reflect.ValueOf(q.Erf),
			"Erfc":            reflect.ValueOf(q.Erfc),
			"Erfcinv":         reflect.ValueOf(q.Erfcinv),
			"Erfinv":          reflect.ValueOf(q.Erfinv),
			"Exp":             reflect.ValueOf(q.Exp),
			"Exp2":            reflect.ValueOf(q.Exp2),
			"Expm1":           reflect.ValueOf(q.Expm1),
			"FMA":             reflect.ValueOf(q.FMA),
			"Float32bits":     reflect.ValueOf(q.Float32bits),
			"Float32frombits": reflect.ValueOf(q.Float32frombits),
			"Float64bits":     reflect.ValueOf(q.Float64bits),
			"Float64frombits": reflect.ValueOf(q.Float64frombits),
			"Floor":           reflect.ValueOf(q.Floor),
			"Frexp":           reflect.ValueOf(q.Frexp),
			"Gamma":           reflect.ValueOf(q.Gamma),
			"Hypot":           reflect.ValueOf(q.Hypot),
			"Ilogb":           reflect.ValueOf(q.Ilogb),
			"Inf":             reflect.ValueOf(q.Inf),
			"IsInf":           reflect.ValueOf(q.IsInf),
			"IsNaN":           reflect.ValueOf(q.IsNaN),
			"J0":              reflect.ValueOf(q.J0),
			"J1":              reflect.ValueOf(q.J1),
			"Jn":              reflect.ValueOf(q.Jn),
			"Ldexp":           reflect.ValueOf(q.Ldexp),
			"Lgamma":          reflect.ValueOf(q.Lgamma),
			"Log":             reflect.ValueOf(q.Log),
			"Log10":           reflect.ValueOf(q.Log10),
			"Log1p":           reflect.ValueOf(q.Log1p),
			"Log2":            reflect.ValueOf(q.Log2),
			"Logb":            reflect.ValueOf(q.Logb),
			"Max":             reflect.ValueOf(q.Max),
			"Min":             reflect.ValueOf(q.Min),
			"Mod":             reflect.ValueOf(q.Mod),
			"Modf":            reflect.ValueOf(q.Modf),
			"NaN":             reflect.ValueOf(q.NaN),
			"Nextafter":       reflect.ValueOf(q.Nextafter),
			"Nextafter32":     reflect.ValueOf(q.Nextafter32),
			"Pow":             reflect.ValueOf(q.Pow),
			"Pow10":           reflect.ValueOf(q.Pow10),
			"Remainder":       reflect.ValueOf(q.Remainder),
			"Round":           reflect.ValueOf(q.Round),
			"RoundToEven":     reflect.ValueOf(q.RoundToEven),
			"Signbit":         reflect.ValueOf(q.Signbit),
			"Sin":             reflect.ValueOf(q.Sin),
			"Sincos":          reflect.ValueOf(q.Sincos),
			"Sinh":            reflect.ValueOf(q.Sinh),
			"Sqrt":            reflect.ValueOf(q.Sqrt),
			"Tan":             reflect.ValueOf(q.Tan),
			"Tanh":            reflect.ValueOf(q.Tanh),
			"Trunc":           reflect.ValueOf(q.Trunc),
			"Y0":              reflect.ValueOf(q.Y0),
			"Y1":              reflect.ValueOf(q.Y1),
			"Yn":              reflect.ValueOf(q.Yn),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"E":                      {"untyped float", constant.MakeFromLiteral("2.71828182845904523536028747135266249775724709369995957496696763", token.FLOAT, 0)},
			"Ln10":                   {"untyped float", constant.MakeFromLiteral("2.3025850929940456840179914546843642076011014886287729760333279", token.FLOAT, 0)},
			"Ln2":                    {"untyped float", constant.MakeFromLiteral("0.693147180559945309417232121458176568075500134360255254120680009", token.FLOAT, 0)},
			"Log10E":                 {"untyped float", constant.BinaryOp(constant.MakeFromLiteral("10000000000000000000000000000000000000000000000000000000000000", token.INT, 0), token.QUO, constant.MakeFromLiteral("23025850929940456840179914546843642076011014886287729760333279", token.INT, 0))},
			"Log2E":                  {"untyped float", constant.BinaryOp(constant.MakeFromLiteral("1000000000000000000000000000000000000000000000000000000000000000", token.INT, 0), token.QUO, constant.MakeFromLiteral("693147180559945309417232121458176568075500134360255254120680009", token.INT, 0))},
			"MaxFloat32":             {"untyped float", constant.MakeFromLiteral("3.4028234663852885981170418348451692544e+38", token.FLOAT, 0)},
			"MaxFloat64":             {"untyped float", constant.MakeFromLiteral("1.797693134862315708145274237317043567981e+308", token.FLOAT, 0)},
			"MaxInt16":               {"untyped int", constant.MakeInt64(int64(q.MaxInt16))},
			"MaxInt32":               {"untyped int", constant.MakeInt64(int64(q.MaxInt32))},
			"MaxInt64":               {"untyped int", constant.MakeInt64(int64(q.MaxInt64))},
			"MaxInt8":                {"untyped int", constant.MakeInt64(int64(q.MaxInt8))},
			"MaxUint16":              {"untyped int", constant.MakeInt64(int64(q.MaxUint16))},
			"MaxUint32":              {"untyped int", constant.MakeInt64(int64(q.MaxUint32))},
			"MaxUint64":              {"untyped int", constant.MakeUint64(uint64(q.MaxUint64))},
			"MaxUint8":               {"untyped int", constant.MakeInt64(int64(q.MaxUint8))},
			"MinInt16":               {"untyped int", constant.MakeInt64(int64(q.MinInt16))},
			"MinInt32":               {"untyped int", constant.MakeInt64(int64(q.MinInt32))},
			"MinInt64":               {"untyped int", constant.MakeInt64(int64(q.MinInt64))},
			"MinInt8":                {"untyped int", constant.MakeInt64(int64(q.MinInt8))},
			"Phi":                    {"untyped float", constant.MakeFromLiteral("1.61803398874989484820458683436563811772030917980576286213544862", token.FLOAT, 0)},
			"Pi":                     {"untyped float", constant.MakeFromLiteral("3.14159265358979323846264338327950288419716939937510582097494459", token.FLOAT, 0)},
			"SmallestNonzeroFloat32": {"untyped float", constant.BinaryOp(constant.MakeFromLiteral("17516230804060213386546619791123951641", token.INT, 0), token.QUO, constant.MakeFromLiteral("12500000000000000000000000000000000000000000000000000000000000000000000000000000000", token.INT, 0))},
			"SmallestNonzeroFloat64": {"untyped float", constant.MakeFromLiteral("4.940656458412465441765687928682213723651e-324", token.FLOAT, 0)},
			"Sqrt2":                  {"untyped float", constant.MakeFromLiteral("1.41421356237309504880168872420969807856967187537694807317667974", token.FLOAT, 0)},
			"SqrtE":                  {"untyped float", constant.MakeFromLiteral("1.64872127070012814684865078781416357165377610071014801157507931", token.FLOAT, 0)},
			"SqrtPhi":                {"untyped float", constant.MakeFromLiteral("1.27201964951406896425242246173749149171560804184009624861664038", token.FLOAT, 0)},
			"SqrtPi":                 {"untyped float", constant.MakeFromLiteral("1.77245385090551602729816748334114518279754945612238712821380779", token.FLOAT, 0)},
		},
	})
}
