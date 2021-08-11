// export by github.com/goplus/interp/cmd/qexp

package macho

import (
	"debug/macho"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("debug/macho", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*debug/macho.FatFile).Close":            (*macho.FatFile).Close,
	"(*debug/macho.File).Close":               (*macho.File).Close,
	"(*debug/macho.File).DWARF":               (*macho.File).DWARF,
	"(*debug/macho.File).ImportedLibraries":   (*macho.File).ImportedLibraries,
	"(*debug/macho.File).ImportedSymbols":     (*macho.File).ImportedSymbols,
	"(*debug/macho.File).Section":             (*macho.File).Section,
	"(*debug/macho.File).Segment":             (*macho.File).Segment,
	"(*debug/macho.FormatError).Error":        (*macho.FormatError).Error,
	"(*debug/macho.Section).Data":             (*macho.Section).Data,
	"(*debug/macho.Section).Open":             (*macho.Section).Open,
	"(*debug/macho.Segment).Data":             (*macho.Segment).Data,
	"(*debug/macho.Segment).Open":             (*macho.Segment).Open,
	"(debug/macho.Cpu).GoString":              (macho.Cpu).GoString,
	"(debug/macho.Cpu).String":                (macho.Cpu).String,
	"(debug/macho.Dylib).Raw":                 (macho.Dylib).Raw,
	"(debug/macho.Dysymtab).Raw":              (macho.Dysymtab).Raw,
	"(debug/macho.FatArch).Close":             (macho.FatArch).Close,
	"(debug/macho.FatArch).DWARF":             (macho.FatArch).DWARF,
	"(debug/macho.FatArch).ImportedLibraries": (macho.FatArch).ImportedLibraries,
	"(debug/macho.FatArch).ImportedSymbols":   (macho.FatArch).ImportedSymbols,
	"(debug/macho.FatArch).Section":           (macho.FatArch).Section,
	"(debug/macho.FatArch).Segment":           (macho.FatArch).Segment,
	"(debug/macho.Load).Raw":                  (macho.Load).Raw,
	"(debug/macho.LoadBytes).Raw":             (macho.LoadBytes).Raw,
	"(debug/macho.LoadCmd).GoString":          (macho.LoadCmd).GoString,
	"(debug/macho.LoadCmd).String":            (macho.LoadCmd).String,
	"(debug/macho.RelocTypeARM).GoString":     (macho.RelocTypeARM).GoString,
	"(debug/macho.RelocTypeARM).String":       (macho.RelocTypeARM).String,
	"(debug/macho.RelocTypeARM64).GoString":   (macho.RelocTypeARM64).GoString,
	"(debug/macho.RelocTypeARM64).String":     (macho.RelocTypeARM64).String,
	"(debug/macho.RelocTypeGeneric).GoString": (macho.RelocTypeGeneric).GoString,
	"(debug/macho.RelocTypeGeneric).String":   (macho.RelocTypeGeneric).String,
	"(debug/macho.RelocTypeX86_64).GoString":  (macho.RelocTypeX86_64).GoString,
	"(debug/macho.RelocTypeX86_64).String":    (macho.RelocTypeX86_64).String,
	"(debug/macho.Rpath).Raw":                 (macho.Rpath).Raw,
	"(debug/macho.Section).ReadAt":            (macho.Section).ReadAt,
	"(debug/macho.Segment).Raw":               (macho.Segment).Raw,
	"(debug/macho.Segment).ReadAt":            (macho.Segment).ReadAt,
	"(debug/macho.Symtab).Raw":                (macho.Symtab).Raw,
	"(debug/macho.Type).GoString":             (macho.Type).GoString,
	"(debug/macho.Type).String":               (macho.Type).String,
	"debug/macho.ErrNotFat":                   &macho.ErrNotFat,
	"debug/macho.NewFatFile":                  macho.NewFatFile,
	"debug/macho.NewFile":                     macho.NewFile,
	"debug/macho.Open":                        macho.Open,
	"debug/macho.OpenFat":                     macho.OpenFat,
}

var typList = []interface{}{
	(*macho.Cpu)(nil),
	(*macho.Dylib)(nil),
	(*macho.DylibCmd)(nil),
	(*macho.Dysymtab)(nil),
	(*macho.DysymtabCmd)(nil),
	(*macho.FatArch)(nil),
	(*macho.FatArchHeader)(nil),
	(*macho.FatFile)(nil),
	(*macho.File)(nil),
	(*macho.FileHeader)(nil),
	(*macho.FormatError)(nil),
	(*macho.Load)(nil),
	(*macho.LoadBytes)(nil),
	(*macho.LoadCmd)(nil),
	(*macho.Nlist32)(nil),
	(*macho.Nlist64)(nil),
	(*macho.Regs386)(nil),
	(*macho.RegsAMD64)(nil),
	(*macho.Reloc)(nil),
	(*macho.RelocTypeARM)(nil),
	(*macho.RelocTypeARM64)(nil),
	(*macho.RelocTypeGeneric)(nil),
	(*macho.RelocTypeX86_64)(nil),
	(*macho.Rpath)(nil),
	(*macho.RpathCmd)(nil),
	(*macho.Section)(nil),
	(*macho.Section32)(nil),
	(*macho.Section64)(nil),
	(*macho.SectionHeader)(nil),
	(*macho.Segment)(nil),
	(*macho.Segment32)(nil),
	(*macho.Segment64)(nil),
	(*macho.SegmentHeader)(nil),
	(*macho.Symbol)(nil),
	(*macho.Symtab)(nil),
	(*macho.SymtabCmd)(nil),
	(*macho.Thread)(nil),
	(*macho.Type)(nil),
}
