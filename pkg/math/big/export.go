// export by github.com/goplus/gossa/cmd/qexp

package big

import (
	"math/big"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("math/big", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*math/big.Float).Abs":            (*big.Float).Abs,
	"(*math/big.Float).Acc":            (*big.Float).Acc,
	"(*math/big.Float).Add":            (*big.Float).Add,
	"(*math/big.Float).Append":         (*big.Float).Append,
	"(*math/big.Float).Cmp":            (*big.Float).Cmp,
	"(*math/big.Float).Copy":           (*big.Float).Copy,
	"(*math/big.Float).Float32":        (*big.Float).Float32,
	"(*math/big.Float).Float64":        (*big.Float).Float64,
	"(*math/big.Float).Format":         (*big.Float).Format,
	"(*math/big.Float).GobDecode":      (*big.Float).GobDecode,
	"(*math/big.Float).GobEncode":      (*big.Float).GobEncode,
	"(*math/big.Float).Int":            (*big.Float).Int,
	"(*math/big.Float).Int64":          (*big.Float).Int64,
	"(*math/big.Float).IsInf":          (*big.Float).IsInf,
	"(*math/big.Float).IsInt":          (*big.Float).IsInt,
	"(*math/big.Float).MantExp":        (*big.Float).MantExp,
	"(*math/big.Float).MarshalText":    (*big.Float).MarshalText,
	"(*math/big.Float).MinPrec":        (*big.Float).MinPrec,
	"(*math/big.Float).Mode":           (*big.Float).Mode,
	"(*math/big.Float).Mul":            (*big.Float).Mul,
	"(*math/big.Float).Neg":            (*big.Float).Neg,
	"(*math/big.Float).Parse":          (*big.Float).Parse,
	"(*math/big.Float).Prec":           (*big.Float).Prec,
	"(*math/big.Float).Quo":            (*big.Float).Quo,
	"(*math/big.Float).Rat":            (*big.Float).Rat,
	"(*math/big.Float).Scan":           (*big.Float).Scan,
	"(*math/big.Float).Set":            (*big.Float).Set,
	"(*math/big.Float).SetFloat64":     (*big.Float).SetFloat64,
	"(*math/big.Float).SetInf":         (*big.Float).SetInf,
	"(*math/big.Float).SetInt":         (*big.Float).SetInt,
	"(*math/big.Float).SetInt64":       (*big.Float).SetInt64,
	"(*math/big.Float).SetMantExp":     (*big.Float).SetMantExp,
	"(*math/big.Float).SetMode":        (*big.Float).SetMode,
	"(*math/big.Float).SetPrec":        (*big.Float).SetPrec,
	"(*math/big.Float).SetRat":         (*big.Float).SetRat,
	"(*math/big.Float).SetString":      (*big.Float).SetString,
	"(*math/big.Float).SetUint64":      (*big.Float).SetUint64,
	"(*math/big.Float).Sign":           (*big.Float).Sign,
	"(*math/big.Float).Signbit":        (*big.Float).Signbit,
	"(*math/big.Float).Sqrt":           (*big.Float).Sqrt,
	"(*math/big.Float).String":         (*big.Float).String,
	"(*math/big.Float).Sub":            (*big.Float).Sub,
	"(*math/big.Float).Text":           (*big.Float).Text,
	"(*math/big.Float).Uint64":         (*big.Float).Uint64,
	"(*math/big.Float).UnmarshalText":  (*big.Float).UnmarshalText,
	"(*math/big.Int).Abs":              (*big.Int).Abs,
	"(*math/big.Int).Add":              (*big.Int).Add,
	"(*math/big.Int).And":              (*big.Int).And,
	"(*math/big.Int).AndNot":           (*big.Int).AndNot,
	"(*math/big.Int).Append":           (*big.Int).Append,
	"(*math/big.Int).Binomial":         (*big.Int).Binomial,
	"(*math/big.Int).Bit":              (*big.Int).Bit,
	"(*math/big.Int).BitLen":           (*big.Int).BitLen,
	"(*math/big.Int).Bits":             (*big.Int).Bits,
	"(*math/big.Int).Bytes":            (*big.Int).Bytes,
	"(*math/big.Int).Cmp":              (*big.Int).Cmp,
	"(*math/big.Int).CmpAbs":           (*big.Int).CmpAbs,
	"(*math/big.Int).Div":              (*big.Int).Div,
	"(*math/big.Int).DivMod":           (*big.Int).DivMod,
	"(*math/big.Int).Exp":              (*big.Int).Exp,
	"(*math/big.Int).Format":           (*big.Int).Format,
	"(*math/big.Int).GCD":              (*big.Int).GCD,
	"(*math/big.Int).GobDecode":        (*big.Int).GobDecode,
	"(*math/big.Int).GobEncode":        (*big.Int).GobEncode,
	"(*math/big.Int).Int64":            (*big.Int).Int64,
	"(*math/big.Int).IsInt64":          (*big.Int).IsInt64,
	"(*math/big.Int).IsUint64":         (*big.Int).IsUint64,
	"(*math/big.Int).Lsh":              (*big.Int).Lsh,
	"(*math/big.Int).MarshalJSON":      (*big.Int).MarshalJSON,
	"(*math/big.Int).MarshalText":      (*big.Int).MarshalText,
	"(*math/big.Int).Mod":              (*big.Int).Mod,
	"(*math/big.Int).ModInverse":       (*big.Int).ModInverse,
	"(*math/big.Int).ModSqrt":          (*big.Int).ModSqrt,
	"(*math/big.Int).Mul":              (*big.Int).Mul,
	"(*math/big.Int).MulRange":         (*big.Int).MulRange,
	"(*math/big.Int).Neg":              (*big.Int).Neg,
	"(*math/big.Int).Not":              (*big.Int).Not,
	"(*math/big.Int).Or":               (*big.Int).Or,
	"(*math/big.Int).ProbablyPrime":    (*big.Int).ProbablyPrime,
	"(*math/big.Int).Quo":              (*big.Int).Quo,
	"(*math/big.Int).QuoRem":           (*big.Int).QuoRem,
	"(*math/big.Int).Rand":             (*big.Int).Rand,
	"(*math/big.Int).Rem":              (*big.Int).Rem,
	"(*math/big.Int).Rsh":              (*big.Int).Rsh,
	"(*math/big.Int).Scan":             (*big.Int).Scan,
	"(*math/big.Int).Set":              (*big.Int).Set,
	"(*math/big.Int).SetBit":           (*big.Int).SetBit,
	"(*math/big.Int).SetBits":          (*big.Int).SetBits,
	"(*math/big.Int).SetBytes":         (*big.Int).SetBytes,
	"(*math/big.Int).SetInt64":         (*big.Int).SetInt64,
	"(*math/big.Int).SetString":        (*big.Int).SetString,
	"(*math/big.Int).SetUint64":        (*big.Int).SetUint64,
	"(*math/big.Int).Sign":             (*big.Int).Sign,
	"(*math/big.Int).Sqrt":             (*big.Int).Sqrt,
	"(*math/big.Int).String":           (*big.Int).String,
	"(*math/big.Int).Sub":              (*big.Int).Sub,
	"(*math/big.Int).Text":             (*big.Int).Text,
	"(*math/big.Int).TrailingZeroBits": (*big.Int).TrailingZeroBits,
	"(*math/big.Int).Uint64":           (*big.Int).Uint64,
	"(*math/big.Int).UnmarshalJSON":    (*big.Int).UnmarshalJSON,
	"(*math/big.Int).UnmarshalText":    (*big.Int).UnmarshalText,
	"(*math/big.Int).Xor":              (*big.Int).Xor,
	"(*math/big.Rat).Abs":              (*big.Rat).Abs,
	"(*math/big.Rat).Add":              (*big.Rat).Add,
	"(*math/big.Rat).Cmp":              (*big.Rat).Cmp,
	"(*math/big.Rat).Denom":            (*big.Rat).Denom,
	"(*math/big.Rat).Float32":          (*big.Rat).Float32,
	"(*math/big.Rat).Float64":          (*big.Rat).Float64,
	"(*math/big.Rat).FloatString":      (*big.Rat).FloatString,
	"(*math/big.Rat).GobDecode":        (*big.Rat).GobDecode,
	"(*math/big.Rat).GobEncode":        (*big.Rat).GobEncode,
	"(*math/big.Rat).Inv":              (*big.Rat).Inv,
	"(*math/big.Rat).IsInt":            (*big.Rat).IsInt,
	"(*math/big.Rat).MarshalText":      (*big.Rat).MarshalText,
	"(*math/big.Rat).Mul":              (*big.Rat).Mul,
	"(*math/big.Rat).Neg":              (*big.Rat).Neg,
	"(*math/big.Rat).Num":              (*big.Rat).Num,
	"(*math/big.Rat).Quo":              (*big.Rat).Quo,
	"(*math/big.Rat).RatString":        (*big.Rat).RatString,
	"(*math/big.Rat).Scan":             (*big.Rat).Scan,
	"(*math/big.Rat).Set":              (*big.Rat).Set,
	"(*math/big.Rat).SetFloat64":       (*big.Rat).SetFloat64,
	"(*math/big.Rat).SetFrac":          (*big.Rat).SetFrac,
	"(*math/big.Rat).SetFrac64":        (*big.Rat).SetFrac64,
	"(*math/big.Rat).SetInt":           (*big.Rat).SetInt,
	"(*math/big.Rat).SetInt64":         (*big.Rat).SetInt64,
	"(*math/big.Rat).SetString":        (*big.Rat).SetString,
	"(*math/big.Rat).SetUint64":        (*big.Rat).SetUint64,
	"(*math/big.Rat).Sign":             (*big.Rat).Sign,
	"(*math/big.Rat).String":           (*big.Rat).String,
	"(*math/big.Rat).Sub":              (*big.Rat).Sub,
	"(*math/big.Rat).UnmarshalText":    (*big.Rat).UnmarshalText,
	"(math/big.Accuracy).String":       (big.Accuracy).String,
	"(math/big.ErrNaN).Error":          (big.ErrNaN).Error,
	"(math/big.RoundingMode).String":   (big.RoundingMode).String,
	"math/big.Jacobi":                  big.Jacobi,
	"math/big.NewFloat":                big.NewFloat,
	"math/big.NewInt":                  big.NewInt,
	"math/big.NewRat":                  big.NewRat,
	"math/big.ParseFloat":              big.ParseFloat,
}

var typList = []interface{}{
	(*big.Accuracy)(nil),
	(*big.ErrNaN)(nil),
	(*big.Float)(nil),
	(*big.Int)(nil),
	(*big.Rat)(nil),
	(*big.RoundingMode)(nil),
	(*big.Word)(nil),
}
