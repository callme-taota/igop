// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package dwarf

import (
	q "debug/dwarf"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "dwarf",
		Path: "debug/dwarf",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"path":            "path",
			"sort":            "sort",
			"strconv":         "strconv",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Type": reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"AddrType":        reflect.TypeOf((*q.AddrType)(nil)).Elem(),
			"ArrayType":       reflect.TypeOf((*q.ArrayType)(nil)).Elem(),
			"Attr":            reflect.TypeOf((*q.Attr)(nil)).Elem(),
			"BasicType":       reflect.TypeOf((*q.BasicType)(nil)).Elem(),
			"BoolType":        reflect.TypeOf((*q.BoolType)(nil)).Elem(),
			"CharType":        reflect.TypeOf((*q.CharType)(nil)).Elem(),
			"Class":           reflect.TypeOf((*q.Class)(nil)).Elem(),
			"CommonType":      reflect.TypeOf((*q.CommonType)(nil)).Elem(),
			"ComplexType":     reflect.TypeOf((*q.ComplexType)(nil)).Elem(),
			"Data":            reflect.TypeOf((*q.Data)(nil)).Elem(),
			"DecodeError":     reflect.TypeOf((*q.DecodeError)(nil)).Elem(),
			"DotDotDotType":   reflect.TypeOf((*q.DotDotDotType)(nil)).Elem(),
			"Entry":           reflect.TypeOf((*q.Entry)(nil)).Elem(),
			"EnumType":        reflect.TypeOf((*q.EnumType)(nil)).Elem(),
			"EnumValue":       reflect.TypeOf((*q.EnumValue)(nil)).Elem(),
			"Field":           reflect.TypeOf((*q.Field)(nil)).Elem(),
			"FloatType":       reflect.TypeOf((*q.FloatType)(nil)).Elem(),
			"FuncType":        reflect.TypeOf((*q.FuncType)(nil)).Elem(),
			"IntType":         reflect.TypeOf((*q.IntType)(nil)).Elem(),
			"LineEntry":       reflect.TypeOf((*q.LineEntry)(nil)).Elem(),
			"LineFile":        reflect.TypeOf((*q.LineFile)(nil)).Elem(),
			"LineReader":      reflect.TypeOf((*q.LineReader)(nil)).Elem(),
			"LineReaderPos":   reflect.TypeOf((*q.LineReaderPos)(nil)).Elem(),
			"Offset":          reflect.TypeOf((*q.Offset)(nil)).Elem(),
			"PtrType":         reflect.TypeOf((*q.PtrType)(nil)).Elem(),
			"QualType":        reflect.TypeOf((*q.QualType)(nil)).Elem(),
			"Reader":          reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"StructField":     reflect.TypeOf((*q.StructField)(nil)).Elem(),
			"StructType":      reflect.TypeOf((*q.StructType)(nil)).Elem(),
			"Tag":             reflect.TypeOf((*q.Tag)(nil)).Elem(),
			"TypedefType":     reflect.TypeOf((*q.TypedefType)(nil)).Elem(),
			"UcharType":       reflect.TypeOf((*q.UcharType)(nil)).Elem(),
			"UintType":        reflect.TypeOf((*q.UintType)(nil)).Elem(),
			"UnspecifiedType": reflect.TypeOf((*q.UnspecifiedType)(nil)).Elem(),
			"UnsupportedType": reflect.TypeOf((*q.UnsupportedType)(nil)).Elem(),
			"VoidType":        reflect.TypeOf((*q.VoidType)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrUnknownPC": reflect.ValueOf(&q.ErrUnknownPC),
		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AttrAbstractOrigin":        {reflect.TypeOf(q.AttrAbstractOrigin), constant.MakeInt64(int64(q.AttrAbstractOrigin))},
			"AttrAccessibility":         {reflect.TypeOf(q.AttrAccessibility), constant.MakeInt64(int64(q.AttrAccessibility))},
			"AttrAddrBase":              {reflect.TypeOf(q.AttrAddrBase), constant.MakeInt64(int64(q.AttrAddrBase))},
			"AttrAddrClass":             {reflect.TypeOf(q.AttrAddrClass), constant.MakeInt64(int64(q.AttrAddrClass))},
			"AttrAlignment":             {reflect.TypeOf(q.AttrAlignment), constant.MakeInt64(int64(q.AttrAlignment))},
			"AttrAllocated":             {reflect.TypeOf(q.AttrAllocated), constant.MakeInt64(int64(q.AttrAllocated))},
			"AttrArtificial":            {reflect.TypeOf(q.AttrArtificial), constant.MakeInt64(int64(q.AttrArtificial))},
			"AttrAssociated":            {reflect.TypeOf(q.AttrAssociated), constant.MakeInt64(int64(q.AttrAssociated))},
			"AttrBaseTypes":             {reflect.TypeOf(q.AttrBaseTypes), constant.MakeInt64(int64(q.AttrBaseTypes))},
			"AttrBinaryScale":           {reflect.TypeOf(q.AttrBinaryScale), constant.MakeInt64(int64(q.AttrBinaryScale))},
			"AttrBitOffset":             {reflect.TypeOf(q.AttrBitOffset), constant.MakeInt64(int64(q.AttrBitOffset))},
			"AttrBitSize":               {reflect.TypeOf(q.AttrBitSize), constant.MakeInt64(int64(q.AttrBitSize))},
			"AttrByteSize":              {reflect.TypeOf(q.AttrByteSize), constant.MakeInt64(int64(q.AttrByteSize))},
			"AttrCallAllCalls":          {reflect.TypeOf(q.AttrCallAllCalls), constant.MakeInt64(int64(q.AttrCallAllCalls))},
			"AttrCallAllSourceCalls":    {reflect.TypeOf(q.AttrCallAllSourceCalls), constant.MakeInt64(int64(q.AttrCallAllSourceCalls))},
			"AttrCallAllTailCalls":      {reflect.TypeOf(q.AttrCallAllTailCalls), constant.MakeInt64(int64(q.AttrCallAllTailCalls))},
			"AttrCallColumn":            {reflect.TypeOf(q.AttrCallColumn), constant.MakeInt64(int64(q.AttrCallColumn))},
			"AttrCallDataLocation":      {reflect.TypeOf(q.AttrCallDataLocation), constant.MakeInt64(int64(q.AttrCallDataLocation))},
			"AttrCallDataValue":         {reflect.TypeOf(q.AttrCallDataValue), constant.MakeInt64(int64(q.AttrCallDataValue))},
			"AttrCallFile":              {reflect.TypeOf(q.AttrCallFile), constant.MakeInt64(int64(q.AttrCallFile))},
			"AttrCallLine":              {reflect.TypeOf(q.AttrCallLine), constant.MakeInt64(int64(q.AttrCallLine))},
			"AttrCallOrigin":            {reflect.TypeOf(q.AttrCallOrigin), constant.MakeInt64(int64(q.AttrCallOrigin))},
			"AttrCallPC":                {reflect.TypeOf(q.AttrCallPC), constant.MakeInt64(int64(q.AttrCallPC))},
			"AttrCallParameter":         {reflect.TypeOf(q.AttrCallParameter), constant.MakeInt64(int64(q.AttrCallParameter))},
			"AttrCallReturnPC":          {reflect.TypeOf(q.AttrCallReturnPC), constant.MakeInt64(int64(q.AttrCallReturnPC))},
			"AttrCallTailCall":          {reflect.TypeOf(q.AttrCallTailCall), constant.MakeInt64(int64(q.AttrCallTailCall))},
			"AttrCallTarget":            {reflect.TypeOf(q.AttrCallTarget), constant.MakeInt64(int64(q.AttrCallTarget))},
			"AttrCallTargetClobbered":   {reflect.TypeOf(q.AttrCallTargetClobbered), constant.MakeInt64(int64(q.AttrCallTargetClobbered))},
			"AttrCallValue":             {reflect.TypeOf(q.AttrCallValue), constant.MakeInt64(int64(q.AttrCallValue))},
			"AttrCalling":               {reflect.TypeOf(q.AttrCalling), constant.MakeInt64(int64(q.AttrCalling))},
			"AttrCommonRef":             {reflect.TypeOf(q.AttrCommonRef), constant.MakeInt64(int64(q.AttrCommonRef))},
			"AttrCompDir":               {reflect.TypeOf(q.AttrCompDir), constant.MakeInt64(int64(q.AttrCompDir))},
			"AttrConstExpr":             {reflect.TypeOf(q.AttrConstExpr), constant.MakeInt64(int64(q.AttrConstExpr))},
			"AttrConstValue":            {reflect.TypeOf(q.AttrConstValue), constant.MakeInt64(int64(q.AttrConstValue))},
			"AttrContainingType":        {reflect.TypeOf(q.AttrContainingType), constant.MakeInt64(int64(q.AttrContainingType))},
			"AttrCount":                 {reflect.TypeOf(q.AttrCount), constant.MakeInt64(int64(q.AttrCount))},
			"AttrDataBitOffset":         {reflect.TypeOf(q.AttrDataBitOffset), constant.MakeInt64(int64(q.AttrDataBitOffset))},
			"AttrDataLocation":          {reflect.TypeOf(q.AttrDataLocation), constant.MakeInt64(int64(q.AttrDataLocation))},
			"AttrDataMemberLoc":         {reflect.TypeOf(q.AttrDataMemberLoc), constant.MakeInt64(int64(q.AttrDataMemberLoc))},
			"AttrDecimalScale":          {reflect.TypeOf(q.AttrDecimalScale), constant.MakeInt64(int64(q.AttrDecimalScale))},
			"AttrDecimalSign":           {reflect.TypeOf(q.AttrDecimalSign), constant.MakeInt64(int64(q.AttrDecimalSign))},
			"AttrDeclColumn":            {reflect.TypeOf(q.AttrDeclColumn), constant.MakeInt64(int64(q.AttrDeclColumn))},
			"AttrDeclFile":              {reflect.TypeOf(q.AttrDeclFile), constant.MakeInt64(int64(q.AttrDeclFile))},
			"AttrDeclLine":              {reflect.TypeOf(q.AttrDeclLine), constant.MakeInt64(int64(q.AttrDeclLine))},
			"AttrDeclaration":           {reflect.TypeOf(q.AttrDeclaration), constant.MakeInt64(int64(q.AttrDeclaration))},
			"AttrDefaultValue":          {reflect.TypeOf(q.AttrDefaultValue), constant.MakeInt64(int64(q.AttrDefaultValue))},
			"AttrDefaulted":             {reflect.TypeOf(q.AttrDefaulted), constant.MakeInt64(int64(q.AttrDefaulted))},
			"AttrDeleted":               {reflect.TypeOf(q.AttrDeleted), constant.MakeInt64(int64(q.AttrDeleted))},
			"AttrDescription":           {reflect.TypeOf(q.AttrDescription), constant.MakeInt64(int64(q.AttrDescription))},
			"AttrDigitCount":            {reflect.TypeOf(q.AttrDigitCount), constant.MakeInt64(int64(q.AttrDigitCount))},
			"AttrDiscr":                 {reflect.TypeOf(q.AttrDiscr), constant.MakeInt64(int64(q.AttrDiscr))},
			"AttrDiscrList":             {reflect.TypeOf(q.AttrDiscrList), constant.MakeInt64(int64(q.AttrDiscrList))},
			"AttrDiscrValue":            {reflect.TypeOf(q.AttrDiscrValue), constant.MakeInt64(int64(q.AttrDiscrValue))},
			"AttrDwoName":               {reflect.TypeOf(q.AttrDwoName), constant.MakeInt64(int64(q.AttrDwoName))},
			"AttrElemental":             {reflect.TypeOf(q.AttrElemental), constant.MakeInt64(int64(q.AttrElemental))},
			"AttrEncoding":              {reflect.TypeOf(q.AttrEncoding), constant.MakeInt64(int64(q.AttrEncoding))},
			"AttrEndianity":             {reflect.TypeOf(q.AttrEndianity), constant.MakeInt64(int64(q.AttrEndianity))},
			"AttrEntrypc":               {reflect.TypeOf(q.AttrEntrypc), constant.MakeInt64(int64(q.AttrEntrypc))},
			"AttrEnumClass":             {reflect.TypeOf(q.AttrEnumClass), constant.MakeInt64(int64(q.AttrEnumClass))},
			"AttrExplicit":              {reflect.TypeOf(q.AttrExplicit), constant.MakeInt64(int64(q.AttrExplicit))},
			"AttrExportSymbols":         {reflect.TypeOf(q.AttrExportSymbols), constant.MakeInt64(int64(q.AttrExportSymbols))},
			"AttrExtension":             {reflect.TypeOf(q.AttrExtension), constant.MakeInt64(int64(q.AttrExtension))},
			"AttrExternal":              {reflect.TypeOf(q.AttrExternal), constant.MakeInt64(int64(q.AttrExternal))},
			"AttrFrameBase":             {reflect.TypeOf(q.AttrFrameBase), constant.MakeInt64(int64(q.AttrFrameBase))},
			"AttrFriend":                {reflect.TypeOf(q.AttrFriend), constant.MakeInt64(int64(q.AttrFriend))},
			"AttrHighpc":                {reflect.TypeOf(q.AttrHighpc), constant.MakeInt64(int64(q.AttrHighpc))},
			"AttrIdentifierCase":        {reflect.TypeOf(q.AttrIdentifierCase), constant.MakeInt64(int64(q.AttrIdentifierCase))},
			"AttrImport":                {reflect.TypeOf(q.AttrImport), constant.MakeInt64(int64(q.AttrImport))},
			"AttrInline":                {reflect.TypeOf(q.AttrInline), constant.MakeInt64(int64(q.AttrInline))},
			"AttrIsOptional":            {reflect.TypeOf(q.AttrIsOptional), constant.MakeInt64(int64(q.AttrIsOptional))},
			"AttrLanguage":              {reflect.TypeOf(q.AttrLanguage), constant.MakeInt64(int64(q.AttrLanguage))},
			"AttrLinkageName":           {reflect.TypeOf(q.AttrLinkageName), constant.MakeInt64(int64(q.AttrLinkageName))},
			"AttrLocation":              {reflect.TypeOf(q.AttrLocation), constant.MakeInt64(int64(q.AttrLocation))},
			"AttrLoclistsBase":          {reflect.TypeOf(q.AttrLoclistsBase), constant.MakeInt64(int64(q.AttrLoclistsBase))},
			"AttrLowerBound":            {reflect.TypeOf(q.AttrLowerBound), constant.MakeInt64(int64(q.AttrLowerBound))},
			"AttrLowpc":                 {reflect.TypeOf(q.AttrLowpc), constant.MakeInt64(int64(q.AttrLowpc))},
			"AttrMacroInfo":             {reflect.TypeOf(q.AttrMacroInfo), constant.MakeInt64(int64(q.AttrMacroInfo))},
			"AttrMacros":                {reflect.TypeOf(q.AttrMacros), constant.MakeInt64(int64(q.AttrMacros))},
			"AttrMainSubprogram":        {reflect.TypeOf(q.AttrMainSubprogram), constant.MakeInt64(int64(q.AttrMainSubprogram))},
			"AttrMutable":               {reflect.TypeOf(q.AttrMutable), constant.MakeInt64(int64(q.AttrMutable))},
			"AttrName":                  {reflect.TypeOf(q.AttrName), constant.MakeInt64(int64(q.AttrName))},
			"AttrNamelistItem":          {reflect.TypeOf(q.AttrNamelistItem), constant.MakeInt64(int64(q.AttrNamelistItem))},
			"AttrNoreturn":              {reflect.TypeOf(q.AttrNoreturn), constant.MakeInt64(int64(q.AttrNoreturn))},
			"AttrObjectPointer":         {reflect.TypeOf(q.AttrObjectPointer), constant.MakeInt64(int64(q.AttrObjectPointer))},
			"AttrOrdering":              {reflect.TypeOf(q.AttrOrdering), constant.MakeInt64(int64(q.AttrOrdering))},
			"AttrPictureString":         {reflect.TypeOf(q.AttrPictureString), constant.MakeInt64(int64(q.AttrPictureString))},
			"AttrPriority":              {reflect.TypeOf(q.AttrPriority), constant.MakeInt64(int64(q.AttrPriority))},
			"AttrProducer":              {reflect.TypeOf(q.AttrProducer), constant.MakeInt64(int64(q.AttrProducer))},
			"AttrPrototyped":            {reflect.TypeOf(q.AttrPrototyped), constant.MakeInt64(int64(q.AttrPrototyped))},
			"AttrPure":                  {reflect.TypeOf(q.AttrPure), constant.MakeInt64(int64(q.AttrPure))},
			"AttrRanges":                {reflect.TypeOf(q.AttrRanges), constant.MakeInt64(int64(q.AttrRanges))},
			"AttrRank":                  {reflect.TypeOf(q.AttrRank), constant.MakeInt64(int64(q.AttrRank))},
			"AttrRecursive":             {reflect.TypeOf(q.AttrRecursive), constant.MakeInt64(int64(q.AttrRecursive))},
			"AttrReference":             {reflect.TypeOf(q.AttrReference), constant.MakeInt64(int64(q.AttrReference))},
			"AttrReturnAddr":            {reflect.TypeOf(q.AttrReturnAddr), constant.MakeInt64(int64(q.AttrReturnAddr))},
			"AttrRnglistsBase":          {reflect.TypeOf(q.AttrRnglistsBase), constant.MakeInt64(int64(q.AttrRnglistsBase))},
			"AttrRvalueReference":       {reflect.TypeOf(q.AttrRvalueReference), constant.MakeInt64(int64(q.AttrRvalueReference))},
			"AttrSegment":               {reflect.TypeOf(q.AttrSegment), constant.MakeInt64(int64(q.AttrSegment))},
			"AttrSibling":               {reflect.TypeOf(q.AttrSibling), constant.MakeInt64(int64(q.AttrSibling))},
			"AttrSignature":             {reflect.TypeOf(q.AttrSignature), constant.MakeInt64(int64(q.AttrSignature))},
			"AttrSmall":                 {reflect.TypeOf(q.AttrSmall), constant.MakeInt64(int64(q.AttrSmall))},
			"AttrSpecification":         {reflect.TypeOf(q.AttrSpecification), constant.MakeInt64(int64(q.AttrSpecification))},
			"AttrStartScope":            {reflect.TypeOf(q.AttrStartScope), constant.MakeInt64(int64(q.AttrStartScope))},
			"AttrStaticLink":            {reflect.TypeOf(q.AttrStaticLink), constant.MakeInt64(int64(q.AttrStaticLink))},
			"AttrStmtList":              {reflect.TypeOf(q.AttrStmtList), constant.MakeInt64(int64(q.AttrStmtList))},
			"AttrStrOffsetsBase":        {reflect.TypeOf(q.AttrStrOffsetsBase), constant.MakeInt64(int64(q.AttrStrOffsetsBase))},
			"AttrStride":                {reflect.TypeOf(q.AttrStride), constant.MakeInt64(int64(q.AttrStride))},
			"AttrStrideSize":            {reflect.TypeOf(q.AttrStrideSize), constant.MakeInt64(int64(q.AttrStrideSize))},
			"AttrStringLength":          {reflect.TypeOf(q.AttrStringLength), constant.MakeInt64(int64(q.AttrStringLength))},
			"AttrStringLengthBitSize":   {reflect.TypeOf(q.AttrStringLengthBitSize), constant.MakeInt64(int64(q.AttrStringLengthBitSize))},
			"AttrStringLengthByteSize":  {reflect.TypeOf(q.AttrStringLengthByteSize), constant.MakeInt64(int64(q.AttrStringLengthByteSize))},
			"AttrThreadsScaled":         {reflect.TypeOf(q.AttrThreadsScaled), constant.MakeInt64(int64(q.AttrThreadsScaled))},
			"AttrTrampoline":            {reflect.TypeOf(q.AttrTrampoline), constant.MakeInt64(int64(q.AttrTrampoline))},
			"AttrType":                  {reflect.TypeOf(q.AttrType), constant.MakeInt64(int64(q.AttrType))},
			"AttrUpperBound":            {reflect.TypeOf(q.AttrUpperBound), constant.MakeInt64(int64(q.AttrUpperBound))},
			"AttrUseLocation":           {reflect.TypeOf(q.AttrUseLocation), constant.MakeInt64(int64(q.AttrUseLocation))},
			"AttrUseUTF8":               {reflect.TypeOf(q.AttrUseUTF8), constant.MakeInt64(int64(q.AttrUseUTF8))},
			"AttrVarParam":              {reflect.TypeOf(q.AttrVarParam), constant.MakeInt64(int64(q.AttrVarParam))},
			"AttrVirtuality":            {reflect.TypeOf(q.AttrVirtuality), constant.MakeInt64(int64(q.AttrVirtuality))},
			"AttrVisibility":            {reflect.TypeOf(q.AttrVisibility), constant.MakeInt64(int64(q.AttrVisibility))},
			"AttrVtableElemLoc":         {reflect.TypeOf(q.AttrVtableElemLoc), constant.MakeInt64(int64(q.AttrVtableElemLoc))},
			"ClassAddrPtr":              {reflect.TypeOf(q.ClassAddrPtr), constant.MakeInt64(int64(q.ClassAddrPtr))},
			"ClassAddress":              {reflect.TypeOf(q.ClassAddress), constant.MakeInt64(int64(q.ClassAddress))},
			"ClassBlock":                {reflect.TypeOf(q.ClassBlock), constant.MakeInt64(int64(q.ClassBlock))},
			"ClassConstant":             {reflect.TypeOf(q.ClassConstant), constant.MakeInt64(int64(q.ClassConstant))},
			"ClassExprLoc":              {reflect.TypeOf(q.ClassExprLoc), constant.MakeInt64(int64(q.ClassExprLoc))},
			"ClassFlag":                 {reflect.TypeOf(q.ClassFlag), constant.MakeInt64(int64(q.ClassFlag))},
			"ClassLinePtr":              {reflect.TypeOf(q.ClassLinePtr), constant.MakeInt64(int64(q.ClassLinePtr))},
			"ClassLocList":              {reflect.TypeOf(q.ClassLocList), constant.MakeInt64(int64(q.ClassLocList))},
			"ClassLocListPtr":           {reflect.TypeOf(q.ClassLocListPtr), constant.MakeInt64(int64(q.ClassLocListPtr))},
			"ClassMacPtr":               {reflect.TypeOf(q.ClassMacPtr), constant.MakeInt64(int64(q.ClassMacPtr))},
			"ClassRangeListPtr":         {reflect.TypeOf(q.ClassRangeListPtr), constant.MakeInt64(int64(q.ClassRangeListPtr))},
			"ClassReference":            {reflect.TypeOf(q.ClassReference), constant.MakeInt64(int64(q.ClassReference))},
			"ClassReferenceAlt":         {reflect.TypeOf(q.ClassReferenceAlt), constant.MakeInt64(int64(q.ClassReferenceAlt))},
			"ClassReferenceSig":         {reflect.TypeOf(q.ClassReferenceSig), constant.MakeInt64(int64(q.ClassReferenceSig))},
			"ClassRngList":              {reflect.TypeOf(q.ClassRngList), constant.MakeInt64(int64(q.ClassRngList))},
			"ClassRngListsPtr":          {reflect.TypeOf(q.ClassRngListsPtr), constant.MakeInt64(int64(q.ClassRngListsPtr))},
			"ClassStrOffsetsPtr":        {reflect.TypeOf(q.ClassStrOffsetsPtr), constant.MakeInt64(int64(q.ClassStrOffsetsPtr))},
			"ClassString":               {reflect.TypeOf(q.ClassString), constant.MakeInt64(int64(q.ClassString))},
			"ClassStringAlt":            {reflect.TypeOf(q.ClassStringAlt), constant.MakeInt64(int64(q.ClassStringAlt))},
			"ClassUnknown":              {reflect.TypeOf(q.ClassUnknown), constant.MakeInt64(int64(q.ClassUnknown))},
			"TagAccessDeclaration":      {reflect.TypeOf(q.TagAccessDeclaration), constant.MakeInt64(int64(q.TagAccessDeclaration))},
			"TagArrayType":              {reflect.TypeOf(q.TagArrayType), constant.MakeInt64(int64(q.TagArrayType))},
			"TagAtomicType":             {reflect.TypeOf(q.TagAtomicType), constant.MakeInt64(int64(q.TagAtomicType))},
			"TagBaseType":               {reflect.TypeOf(q.TagBaseType), constant.MakeInt64(int64(q.TagBaseType))},
			"TagCallSite":               {reflect.TypeOf(q.TagCallSite), constant.MakeInt64(int64(q.TagCallSite))},
			"TagCallSiteParameter":      {reflect.TypeOf(q.TagCallSiteParameter), constant.MakeInt64(int64(q.TagCallSiteParameter))},
			"TagCatchDwarfBlock":        {reflect.TypeOf(q.TagCatchDwarfBlock), constant.MakeInt64(int64(q.TagCatchDwarfBlock))},
			"TagClassType":              {reflect.TypeOf(q.TagClassType), constant.MakeInt64(int64(q.TagClassType))},
			"TagCoarrayType":            {reflect.TypeOf(q.TagCoarrayType), constant.MakeInt64(int64(q.TagCoarrayType))},
			"TagCommonDwarfBlock":       {reflect.TypeOf(q.TagCommonDwarfBlock), constant.MakeInt64(int64(q.TagCommonDwarfBlock))},
			"TagCommonInclusion":        {reflect.TypeOf(q.TagCommonInclusion), constant.MakeInt64(int64(q.TagCommonInclusion))},
			"TagCompileUnit":            {reflect.TypeOf(q.TagCompileUnit), constant.MakeInt64(int64(q.TagCompileUnit))},
			"TagCondition":              {reflect.TypeOf(q.TagCondition), constant.MakeInt64(int64(q.TagCondition))},
			"TagConstType":              {reflect.TypeOf(q.TagConstType), constant.MakeInt64(int64(q.TagConstType))},
			"TagConstant":               {reflect.TypeOf(q.TagConstant), constant.MakeInt64(int64(q.TagConstant))},
			"TagDwarfProcedure":         {reflect.TypeOf(q.TagDwarfProcedure), constant.MakeInt64(int64(q.TagDwarfProcedure))},
			"TagDynamicType":            {reflect.TypeOf(q.TagDynamicType), constant.MakeInt64(int64(q.TagDynamicType))},
			"TagEntryPoint":             {reflect.TypeOf(q.TagEntryPoint), constant.MakeInt64(int64(q.TagEntryPoint))},
			"TagEnumerationType":        {reflect.TypeOf(q.TagEnumerationType), constant.MakeInt64(int64(q.TagEnumerationType))},
			"TagEnumerator":             {reflect.TypeOf(q.TagEnumerator), constant.MakeInt64(int64(q.TagEnumerator))},
			"TagFileType":               {reflect.TypeOf(q.TagFileType), constant.MakeInt64(int64(q.TagFileType))},
			"TagFormalParameter":        {reflect.TypeOf(q.TagFormalParameter), constant.MakeInt64(int64(q.TagFormalParameter))},
			"TagFriend":                 {reflect.TypeOf(q.TagFriend), constant.MakeInt64(int64(q.TagFriend))},
			"TagGenericSubrange":        {reflect.TypeOf(q.TagGenericSubrange), constant.MakeInt64(int64(q.TagGenericSubrange))},
			"TagImmutableType":          {reflect.TypeOf(q.TagImmutableType), constant.MakeInt64(int64(q.TagImmutableType))},
			"TagImportedDeclaration":    {reflect.TypeOf(q.TagImportedDeclaration), constant.MakeInt64(int64(q.TagImportedDeclaration))},
			"TagImportedModule":         {reflect.TypeOf(q.TagImportedModule), constant.MakeInt64(int64(q.TagImportedModule))},
			"TagImportedUnit":           {reflect.TypeOf(q.TagImportedUnit), constant.MakeInt64(int64(q.TagImportedUnit))},
			"TagInheritance":            {reflect.TypeOf(q.TagInheritance), constant.MakeInt64(int64(q.TagInheritance))},
			"TagInlinedSubroutine":      {reflect.TypeOf(q.TagInlinedSubroutine), constant.MakeInt64(int64(q.TagInlinedSubroutine))},
			"TagInterfaceType":          {reflect.TypeOf(q.TagInterfaceType), constant.MakeInt64(int64(q.TagInterfaceType))},
			"TagLabel":                  {reflect.TypeOf(q.TagLabel), constant.MakeInt64(int64(q.TagLabel))},
			"TagLexDwarfBlock":          {reflect.TypeOf(q.TagLexDwarfBlock), constant.MakeInt64(int64(q.TagLexDwarfBlock))},
			"TagMember":                 {reflect.TypeOf(q.TagMember), constant.MakeInt64(int64(q.TagMember))},
			"TagModule":                 {reflect.TypeOf(q.TagModule), constant.MakeInt64(int64(q.TagModule))},
			"TagMutableType":            {reflect.TypeOf(q.TagMutableType), constant.MakeInt64(int64(q.TagMutableType))},
			"TagNamelist":               {reflect.TypeOf(q.TagNamelist), constant.MakeInt64(int64(q.TagNamelist))},
			"TagNamelistItem":           {reflect.TypeOf(q.TagNamelistItem), constant.MakeInt64(int64(q.TagNamelistItem))},
			"TagNamespace":              {reflect.TypeOf(q.TagNamespace), constant.MakeInt64(int64(q.TagNamespace))},
			"TagPackedType":             {reflect.TypeOf(q.TagPackedType), constant.MakeInt64(int64(q.TagPackedType))},
			"TagPartialUnit":            {reflect.TypeOf(q.TagPartialUnit), constant.MakeInt64(int64(q.TagPartialUnit))},
			"TagPointerType":            {reflect.TypeOf(q.TagPointerType), constant.MakeInt64(int64(q.TagPointerType))},
			"TagPtrToMemberType":        {reflect.TypeOf(q.TagPtrToMemberType), constant.MakeInt64(int64(q.TagPtrToMemberType))},
			"TagReferenceType":          {reflect.TypeOf(q.TagReferenceType), constant.MakeInt64(int64(q.TagReferenceType))},
			"TagRestrictType":           {reflect.TypeOf(q.TagRestrictType), constant.MakeInt64(int64(q.TagRestrictType))},
			"TagRvalueReferenceType":    {reflect.TypeOf(q.TagRvalueReferenceType), constant.MakeInt64(int64(q.TagRvalueReferenceType))},
			"TagSetType":                {reflect.TypeOf(q.TagSetType), constant.MakeInt64(int64(q.TagSetType))},
			"TagSharedType":             {reflect.TypeOf(q.TagSharedType), constant.MakeInt64(int64(q.TagSharedType))},
			"TagSkeletonUnit":           {reflect.TypeOf(q.TagSkeletonUnit), constant.MakeInt64(int64(q.TagSkeletonUnit))},
			"TagStringType":             {reflect.TypeOf(q.TagStringType), constant.MakeInt64(int64(q.TagStringType))},
			"TagStructType":             {reflect.TypeOf(q.TagStructType), constant.MakeInt64(int64(q.TagStructType))},
			"TagSubprogram":             {reflect.TypeOf(q.TagSubprogram), constant.MakeInt64(int64(q.TagSubprogram))},
			"TagSubrangeType":           {reflect.TypeOf(q.TagSubrangeType), constant.MakeInt64(int64(q.TagSubrangeType))},
			"TagSubroutineType":         {reflect.TypeOf(q.TagSubroutineType), constant.MakeInt64(int64(q.TagSubroutineType))},
			"TagTemplateAlias":          {reflect.TypeOf(q.TagTemplateAlias), constant.MakeInt64(int64(q.TagTemplateAlias))},
			"TagTemplateTypeParameter":  {reflect.TypeOf(q.TagTemplateTypeParameter), constant.MakeInt64(int64(q.TagTemplateTypeParameter))},
			"TagTemplateValueParameter": {reflect.TypeOf(q.TagTemplateValueParameter), constant.MakeInt64(int64(q.TagTemplateValueParameter))},
			"TagThrownType":             {reflect.TypeOf(q.TagThrownType), constant.MakeInt64(int64(q.TagThrownType))},
			"TagTryDwarfBlock":          {reflect.TypeOf(q.TagTryDwarfBlock), constant.MakeInt64(int64(q.TagTryDwarfBlock))},
			"TagTypeUnit":               {reflect.TypeOf(q.TagTypeUnit), constant.MakeInt64(int64(q.TagTypeUnit))},
			"TagTypedef":                {reflect.TypeOf(q.TagTypedef), constant.MakeInt64(int64(q.TagTypedef))},
			"TagUnionType":              {reflect.TypeOf(q.TagUnionType), constant.MakeInt64(int64(q.TagUnionType))},
			"TagUnspecifiedParameters":  {reflect.TypeOf(q.TagUnspecifiedParameters), constant.MakeInt64(int64(q.TagUnspecifiedParameters))},
			"TagUnspecifiedType":        {reflect.TypeOf(q.TagUnspecifiedType), constant.MakeInt64(int64(q.TagUnspecifiedType))},
			"TagVariable":               {reflect.TypeOf(q.TagVariable), constant.MakeInt64(int64(q.TagVariable))},
			"TagVariant":                {reflect.TypeOf(q.TagVariant), constant.MakeInt64(int64(q.TagVariant))},
			"TagVariantPart":            {reflect.TypeOf(q.TagVariantPart), constant.MakeInt64(int64(q.TagVariantPart))},
			"TagVolatileType":           {reflect.TypeOf(q.TagVolatileType), constant.MakeInt64(int64(q.TagVolatileType))},
			"TagWithStmt":               {reflect.TypeOf(q.TagWithStmt), constant.MakeInt64(int64(q.TagWithStmt))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
