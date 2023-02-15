// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package macho

import (
	q "debug/macho"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "macho",
		Path: "debug/macho",
		Deps: map[string]string{
			"bytes":           "bytes",
			"compress/zlib":   "zlib",
			"debug/dwarf":     "dwarf",
			"encoding/binary": "binary",
			"fmt":             "fmt",
			"io":              "io",
			"os":              "os",
			"strconv":         "strconv",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Load": reflect.TypeOf((*q.Load)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Cpu":              reflect.TypeOf((*q.Cpu)(nil)).Elem(),
			"Dylib":            reflect.TypeOf((*q.Dylib)(nil)).Elem(),
			"DylibCmd":         reflect.TypeOf((*q.DylibCmd)(nil)).Elem(),
			"Dysymtab":         reflect.TypeOf((*q.Dysymtab)(nil)).Elem(),
			"DysymtabCmd":      reflect.TypeOf((*q.DysymtabCmd)(nil)).Elem(),
			"FatArch":          reflect.TypeOf((*q.FatArch)(nil)).Elem(),
			"FatArchHeader":    reflect.TypeOf((*q.FatArchHeader)(nil)).Elem(),
			"FatFile":          reflect.TypeOf((*q.FatFile)(nil)).Elem(),
			"File":             reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileHeader":       reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"FormatError":      reflect.TypeOf((*q.FormatError)(nil)).Elem(),
			"LoadBytes":        reflect.TypeOf((*q.LoadBytes)(nil)).Elem(),
			"LoadCmd":          reflect.TypeOf((*q.LoadCmd)(nil)).Elem(),
			"Nlist32":          reflect.TypeOf((*q.Nlist32)(nil)).Elem(),
			"Nlist64":          reflect.TypeOf((*q.Nlist64)(nil)).Elem(),
			"Regs386":          reflect.TypeOf((*q.Regs386)(nil)).Elem(),
			"RegsAMD64":        reflect.TypeOf((*q.RegsAMD64)(nil)).Elem(),
			"Reloc":            reflect.TypeOf((*q.Reloc)(nil)).Elem(),
			"RelocTypeARM":     reflect.TypeOf((*q.RelocTypeARM)(nil)).Elem(),
			"RelocTypeARM64":   reflect.TypeOf((*q.RelocTypeARM64)(nil)).Elem(),
			"RelocTypeGeneric": reflect.TypeOf((*q.RelocTypeGeneric)(nil)).Elem(),
			"RelocTypeX86_64":  reflect.TypeOf((*q.RelocTypeX86_64)(nil)).Elem(),
			"Rpath":            reflect.TypeOf((*q.Rpath)(nil)).Elem(),
			"RpathCmd":         reflect.TypeOf((*q.RpathCmd)(nil)).Elem(),
			"Section":          reflect.TypeOf((*q.Section)(nil)).Elem(),
			"Section32":        reflect.TypeOf((*q.Section32)(nil)).Elem(),
			"Section64":        reflect.TypeOf((*q.Section64)(nil)).Elem(),
			"SectionHeader":    reflect.TypeOf((*q.SectionHeader)(nil)).Elem(),
			"Segment":          reflect.TypeOf((*q.Segment)(nil)).Elem(),
			"Segment32":        reflect.TypeOf((*q.Segment32)(nil)).Elem(),
			"Segment64":        reflect.TypeOf((*q.Segment64)(nil)).Elem(),
			"SegmentHeader":    reflect.TypeOf((*q.SegmentHeader)(nil)).Elem(),
			"Symbol":           reflect.TypeOf((*q.Symbol)(nil)).Elem(),
			"Symtab":           reflect.TypeOf((*q.Symtab)(nil)).Elem(),
			"SymtabCmd":        reflect.TypeOf((*q.SymtabCmd)(nil)).Elem(),
			"Thread":           reflect.TypeOf((*q.Thread)(nil)).Elem(),
			"Type":             reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrNotFat": reflect.ValueOf(&q.ErrNotFat),
		},
		Funcs: map[string]reflect.Value{
			"NewFatFile": reflect.ValueOf(q.NewFatFile),
			"NewFile":    reflect.ValueOf(q.NewFile),
			"Open":       reflect.ValueOf(q.Open),
			"OpenFat":    reflect.ValueOf(q.OpenFat),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ARM64_RELOC_ADDEND":              {reflect.TypeOf(q.ARM64_RELOC_ADDEND), constant.MakeInt64(int64(q.ARM64_RELOC_ADDEND))},
			"ARM64_RELOC_BRANCH26":            {reflect.TypeOf(q.ARM64_RELOC_BRANCH26), constant.MakeInt64(int64(q.ARM64_RELOC_BRANCH26))},
			"ARM64_RELOC_GOT_LOAD_PAGE21":     {reflect.TypeOf(q.ARM64_RELOC_GOT_LOAD_PAGE21), constant.MakeInt64(int64(q.ARM64_RELOC_GOT_LOAD_PAGE21))},
			"ARM64_RELOC_GOT_LOAD_PAGEOFF12":  {reflect.TypeOf(q.ARM64_RELOC_GOT_LOAD_PAGEOFF12), constant.MakeInt64(int64(q.ARM64_RELOC_GOT_LOAD_PAGEOFF12))},
			"ARM64_RELOC_PAGE21":              {reflect.TypeOf(q.ARM64_RELOC_PAGE21), constant.MakeInt64(int64(q.ARM64_RELOC_PAGE21))},
			"ARM64_RELOC_PAGEOFF12":           {reflect.TypeOf(q.ARM64_RELOC_PAGEOFF12), constant.MakeInt64(int64(q.ARM64_RELOC_PAGEOFF12))},
			"ARM64_RELOC_POINTER_TO_GOT":      {reflect.TypeOf(q.ARM64_RELOC_POINTER_TO_GOT), constant.MakeInt64(int64(q.ARM64_RELOC_POINTER_TO_GOT))},
			"ARM64_RELOC_SUBTRACTOR":          {reflect.TypeOf(q.ARM64_RELOC_SUBTRACTOR), constant.MakeInt64(int64(q.ARM64_RELOC_SUBTRACTOR))},
			"ARM64_RELOC_TLVP_LOAD_PAGE21":    {reflect.TypeOf(q.ARM64_RELOC_TLVP_LOAD_PAGE21), constant.MakeInt64(int64(q.ARM64_RELOC_TLVP_LOAD_PAGE21))},
			"ARM64_RELOC_TLVP_LOAD_PAGEOFF12": {reflect.TypeOf(q.ARM64_RELOC_TLVP_LOAD_PAGEOFF12), constant.MakeInt64(int64(q.ARM64_RELOC_TLVP_LOAD_PAGEOFF12))},
			"ARM64_RELOC_UNSIGNED":            {reflect.TypeOf(q.ARM64_RELOC_UNSIGNED), constant.MakeInt64(int64(q.ARM64_RELOC_UNSIGNED))},
			"ARM_RELOC_BR24":                  {reflect.TypeOf(q.ARM_RELOC_BR24), constant.MakeInt64(int64(q.ARM_RELOC_BR24))},
			"ARM_RELOC_HALF":                  {reflect.TypeOf(q.ARM_RELOC_HALF), constant.MakeInt64(int64(q.ARM_RELOC_HALF))},
			"ARM_RELOC_HALF_SECTDIFF":         {reflect.TypeOf(q.ARM_RELOC_HALF_SECTDIFF), constant.MakeInt64(int64(q.ARM_RELOC_HALF_SECTDIFF))},
			"ARM_RELOC_LOCAL_SECTDIFF":        {reflect.TypeOf(q.ARM_RELOC_LOCAL_SECTDIFF), constant.MakeInt64(int64(q.ARM_RELOC_LOCAL_SECTDIFF))},
			"ARM_RELOC_PAIR":                  {reflect.TypeOf(q.ARM_RELOC_PAIR), constant.MakeInt64(int64(q.ARM_RELOC_PAIR))},
			"ARM_RELOC_PB_LA_PTR":             {reflect.TypeOf(q.ARM_RELOC_PB_LA_PTR), constant.MakeInt64(int64(q.ARM_RELOC_PB_LA_PTR))},
			"ARM_RELOC_SECTDIFF":              {reflect.TypeOf(q.ARM_RELOC_SECTDIFF), constant.MakeInt64(int64(q.ARM_RELOC_SECTDIFF))},
			"ARM_RELOC_VANILLA":               {reflect.TypeOf(q.ARM_RELOC_VANILLA), constant.MakeInt64(int64(q.ARM_RELOC_VANILLA))},
			"ARM_THUMB_32BIT_BRANCH":          {reflect.TypeOf(q.ARM_THUMB_32BIT_BRANCH), constant.MakeInt64(int64(q.ARM_THUMB_32BIT_BRANCH))},
			"ARM_THUMB_RELOC_BR22":            {reflect.TypeOf(q.ARM_THUMB_RELOC_BR22), constant.MakeInt64(int64(q.ARM_THUMB_RELOC_BR22))},
			"Cpu386":                          {reflect.TypeOf(q.Cpu386), constant.MakeInt64(int64(q.Cpu386))},
			"CpuAmd64":                        {reflect.TypeOf(q.CpuAmd64), constant.MakeInt64(int64(q.CpuAmd64))},
			"CpuArm":                          {reflect.TypeOf(q.CpuArm), constant.MakeInt64(int64(q.CpuArm))},
			"CpuArm64":                        {reflect.TypeOf(q.CpuArm64), constant.MakeInt64(int64(q.CpuArm64))},
			"CpuPpc":                          {reflect.TypeOf(q.CpuPpc), constant.MakeInt64(int64(q.CpuPpc))},
			"CpuPpc64":                        {reflect.TypeOf(q.CpuPpc64), constant.MakeInt64(int64(q.CpuPpc64))},
			"FlagAllModsBound":                {reflect.TypeOf(q.FlagAllModsBound), constant.MakeInt64(int64(q.FlagAllModsBound))},
			"FlagAllowStackExecution":         {reflect.TypeOf(q.FlagAllowStackExecution), constant.MakeInt64(int64(q.FlagAllowStackExecution))},
			"FlagAppExtensionSafe":            {reflect.TypeOf(q.FlagAppExtensionSafe), constant.MakeInt64(int64(q.FlagAppExtensionSafe))},
			"FlagBindAtLoad":                  {reflect.TypeOf(q.FlagBindAtLoad), constant.MakeInt64(int64(q.FlagBindAtLoad))},
			"FlagBindsToWeak":                 {reflect.TypeOf(q.FlagBindsToWeak), constant.MakeInt64(int64(q.FlagBindsToWeak))},
			"FlagCanonical":                   {reflect.TypeOf(q.FlagCanonical), constant.MakeInt64(int64(q.FlagCanonical))},
			"FlagDeadStrippableDylib":         {reflect.TypeOf(q.FlagDeadStrippableDylib), constant.MakeInt64(int64(q.FlagDeadStrippableDylib))},
			"FlagDyldLink":                    {reflect.TypeOf(q.FlagDyldLink), constant.MakeInt64(int64(q.FlagDyldLink))},
			"FlagForceFlat":                   {reflect.TypeOf(q.FlagForceFlat), constant.MakeInt64(int64(q.FlagForceFlat))},
			"FlagHasTLVDescriptors":           {reflect.TypeOf(q.FlagHasTLVDescriptors), constant.MakeInt64(int64(q.FlagHasTLVDescriptors))},
			"FlagIncrLink":                    {reflect.TypeOf(q.FlagIncrLink), constant.MakeInt64(int64(q.FlagIncrLink))},
			"FlagLazyInit":                    {reflect.TypeOf(q.FlagLazyInit), constant.MakeInt64(int64(q.FlagLazyInit))},
			"FlagNoFixPrebinding":             {reflect.TypeOf(q.FlagNoFixPrebinding), constant.MakeInt64(int64(q.FlagNoFixPrebinding))},
			"FlagNoHeapExecution":             {reflect.TypeOf(q.FlagNoHeapExecution), constant.MakeInt64(int64(q.FlagNoHeapExecution))},
			"FlagNoMultiDefs":                 {reflect.TypeOf(q.FlagNoMultiDefs), constant.MakeInt64(int64(q.FlagNoMultiDefs))},
			"FlagNoReexportedDylibs":          {reflect.TypeOf(q.FlagNoReexportedDylibs), constant.MakeInt64(int64(q.FlagNoReexportedDylibs))},
			"FlagNoUndefs":                    {reflect.TypeOf(q.FlagNoUndefs), constant.MakeInt64(int64(q.FlagNoUndefs))},
			"FlagPIE":                         {reflect.TypeOf(q.FlagPIE), constant.MakeInt64(int64(q.FlagPIE))},
			"FlagPrebindable":                 {reflect.TypeOf(q.FlagPrebindable), constant.MakeInt64(int64(q.FlagPrebindable))},
			"FlagPrebound":                    {reflect.TypeOf(q.FlagPrebound), constant.MakeInt64(int64(q.FlagPrebound))},
			"FlagRootSafe":                    {reflect.TypeOf(q.FlagRootSafe), constant.MakeInt64(int64(q.FlagRootSafe))},
			"FlagSetuidSafe":                  {reflect.TypeOf(q.FlagSetuidSafe), constant.MakeInt64(int64(q.FlagSetuidSafe))},
			"FlagSplitSegs":                   {reflect.TypeOf(q.FlagSplitSegs), constant.MakeInt64(int64(q.FlagSplitSegs))},
			"FlagSubsectionsViaSymbols":       {reflect.TypeOf(q.FlagSubsectionsViaSymbols), constant.MakeInt64(int64(q.FlagSubsectionsViaSymbols))},
			"FlagTwoLevel":                    {reflect.TypeOf(q.FlagTwoLevel), constant.MakeInt64(int64(q.FlagTwoLevel))},
			"FlagWeakDefines":                 {reflect.TypeOf(q.FlagWeakDefines), constant.MakeInt64(int64(q.FlagWeakDefines))},
			"GENERIC_RELOC_LOCAL_SECTDIFF":    {reflect.TypeOf(q.GENERIC_RELOC_LOCAL_SECTDIFF), constant.MakeInt64(int64(q.GENERIC_RELOC_LOCAL_SECTDIFF))},
			"GENERIC_RELOC_PAIR":              {reflect.TypeOf(q.GENERIC_RELOC_PAIR), constant.MakeInt64(int64(q.GENERIC_RELOC_PAIR))},
			"GENERIC_RELOC_PB_LA_PTR":         {reflect.TypeOf(q.GENERIC_RELOC_PB_LA_PTR), constant.MakeInt64(int64(q.GENERIC_RELOC_PB_LA_PTR))},
			"GENERIC_RELOC_SECTDIFF":          {reflect.TypeOf(q.GENERIC_RELOC_SECTDIFF), constant.MakeInt64(int64(q.GENERIC_RELOC_SECTDIFF))},
			"GENERIC_RELOC_TLV":               {reflect.TypeOf(q.GENERIC_RELOC_TLV), constant.MakeInt64(int64(q.GENERIC_RELOC_TLV))},
			"GENERIC_RELOC_VANILLA":           {reflect.TypeOf(q.GENERIC_RELOC_VANILLA), constant.MakeInt64(int64(q.GENERIC_RELOC_VANILLA))},
			"LoadCmdDylib":                    {reflect.TypeOf(q.LoadCmdDylib), constant.MakeInt64(int64(q.LoadCmdDylib))},
			"LoadCmdDylinker":                 {reflect.TypeOf(q.LoadCmdDylinker), constant.MakeInt64(int64(q.LoadCmdDylinker))},
			"LoadCmdDysymtab":                 {reflect.TypeOf(q.LoadCmdDysymtab), constant.MakeInt64(int64(q.LoadCmdDysymtab))},
			"LoadCmdRpath":                    {reflect.TypeOf(q.LoadCmdRpath), constant.MakeInt64(int64(q.LoadCmdRpath))},
			"LoadCmdSegment":                  {reflect.TypeOf(q.LoadCmdSegment), constant.MakeInt64(int64(q.LoadCmdSegment))},
			"LoadCmdSegment64":                {reflect.TypeOf(q.LoadCmdSegment64), constant.MakeInt64(int64(q.LoadCmdSegment64))},
			"LoadCmdSymtab":                   {reflect.TypeOf(q.LoadCmdSymtab), constant.MakeInt64(int64(q.LoadCmdSymtab))},
			"LoadCmdThread":                   {reflect.TypeOf(q.LoadCmdThread), constant.MakeInt64(int64(q.LoadCmdThread))},
			"LoadCmdUnixThread":               {reflect.TypeOf(q.LoadCmdUnixThread), constant.MakeInt64(int64(q.LoadCmdUnixThread))},
			"Magic32":                         {reflect.TypeOf(q.Magic32), constant.MakeInt64(int64(q.Magic32))},
			"Magic64":                         {reflect.TypeOf(q.Magic64), constant.MakeInt64(int64(q.Magic64))},
			"MagicFat":                        {reflect.TypeOf(q.MagicFat), constant.MakeInt64(int64(q.MagicFat))},
			"TypeBundle":                      {reflect.TypeOf(q.TypeBundle), constant.MakeInt64(int64(q.TypeBundle))},
			"TypeDylib":                       {reflect.TypeOf(q.TypeDylib), constant.MakeInt64(int64(q.TypeDylib))},
			"TypeExec":                        {reflect.TypeOf(q.TypeExec), constant.MakeInt64(int64(q.TypeExec))},
			"TypeObj":                         {reflect.TypeOf(q.TypeObj), constant.MakeInt64(int64(q.TypeObj))},
			"X86_64_RELOC_BRANCH":             {reflect.TypeOf(q.X86_64_RELOC_BRANCH), constant.MakeInt64(int64(q.X86_64_RELOC_BRANCH))},
			"X86_64_RELOC_GOT":                {reflect.TypeOf(q.X86_64_RELOC_GOT), constant.MakeInt64(int64(q.X86_64_RELOC_GOT))},
			"X86_64_RELOC_GOT_LOAD":           {reflect.TypeOf(q.X86_64_RELOC_GOT_LOAD), constant.MakeInt64(int64(q.X86_64_RELOC_GOT_LOAD))},
			"X86_64_RELOC_SIGNED":             {reflect.TypeOf(q.X86_64_RELOC_SIGNED), constant.MakeInt64(int64(q.X86_64_RELOC_SIGNED))},
			"X86_64_RELOC_SIGNED_1":           {reflect.TypeOf(q.X86_64_RELOC_SIGNED_1), constant.MakeInt64(int64(q.X86_64_RELOC_SIGNED_1))},
			"X86_64_RELOC_SIGNED_2":           {reflect.TypeOf(q.X86_64_RELOC_SIGNED_2), constant.MakeInt64(int64(q.X86_64_RELOC_SIGNED_2))},
			"X86_64_RELOC_SIGNED_4":           {reflect.TypeOf(q.X86_64_RELOC_SIGNED_4), constant.MakeInt64(int64(q.X86_64_RELOC_SIGNED_4))},
			"X86_64_RELOC_SUBTRACTOR":         {reflect.TypeOf(q.X86_64_RELOC_SUBTRACTOR), constant.MakeInt64(int64(q.X86_64_RELOC_SUBTRACTOR))},
			"X86_64_RELOC_TLV":                {reflect.TypeOf(q.X86_64_RELOC_TLV), constant.MakeInt64(int64(q.X86_64_RELOC_TLV))},
			"X86_64_RELOC_UNSIGNED":           {reflect.TypeOf(q.X86_64_RELOC_UNSIGNED), constant.MakeInt64(int64(q.X86_64_RELOC_UNSIGNED))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
