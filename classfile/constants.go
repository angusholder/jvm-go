package classfile

import "github.com/pkg/errors"

const (
	constUtf8    = 1
	constInteger = 3
	constFloat   = 4
	constLong    = 5
	constDouble  = 6
	constClass   = 7
	constString  = 8
	//constFieldRef           = 9
	//constMethodRef          = 10
	//constInterfaceMethodRef = 11
	//constNameAndType        = 12
)

func parseConstant(scn *BinScanner) (cp CpInfo, err error) {
	switch scn.Uint8() {
	case constClass:
		cp = ConstantClassInfo{scn.Uint16()}
	case constString:
		cp = ConstantStringInfo{scn.Uint16()}
	case constInteger:
		cp = ConstantIntegerInfo{scn.Uint32()}
	case constFloat:
		cp = ConstantFloatInfo{scn.Uint32()}
	case constLong:
		cp = ConstantLongInfo{scn.Uint32(), scn.Uint32()}
	case constDouble:
		cp = ConstantDoubleInfo{scn.Uint32(), scn.Uint32()}
	case constUtf8:
		cp = ConstantUtf8Info{scn.Bytes(int(scn.Uint16()))}
	default:
		err = errors.New("invalid constant kind")
	}
	return
}

type CpInfo interface {
	constantMarkerFunc()
}

type ConstantClassInfo struct {
	NameIndex uint16 // Must be the index of a ConstantUtf8Info
}

func (c ConstantClassInfo) constantMarkerFunc() {}

type ConstantStringInfo struct {
	StringIndex uint16
}

func (c ConstantStringInfo) constantMarkerFunc() {}

type ConstantIntegerInfo struct {
	Bytes uint32
}

func (c ConstantIntegerInfo) constantMarkerFunc() {}

type ConstantFloatInfo struct {
	Bytes uint32
}

func (c ConstantFloatInfo) constantMarkerFunc() {}

type ConstantLongInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c ConstantLongInfo) constantMarkerFunc() {}

type ConstantDoubleInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c ConstantDoubleInfo) constantMarkerFunc() {}

type ConstantUtf8Info struct {
	Bytes []uint8
}

func (c ConstantUtf8Info) constantMarkerFunc() {}
