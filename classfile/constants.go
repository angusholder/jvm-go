package classfile

import (
	"fmt"
	"github.com/angusholder/jvm-go/jvm"
	"github.com/pkg/errors"
	"math"
)

const (
	constClass              = 7
	constString             = 8
	constInteger            = 3
	constFloat              = 4
	constLong               = 5
	constDouble             = 6
	constUtf8               = 1
	constFieldRef           = 9
	constMethodRef          = 10
	constInterfaceMethodRef = 11
	constNameAndType        = 12
)

func parseConstantPool(scn *BinScanner) ([]CpInfo, error) {
	constantCount := scn.Uint16()
	pool := make([]CpInfo, constantCount)
	fmt.Printf("Parsing %d (0x%x) constants\n", constantCount, constantCount)
	for i := 1; i < int(constantCount); i++ {
		con, err := parseConstant(scn)

		if err != nil {
			return nil, err
		}

		switch v := con.(type) {
		case ConstantIntegerInfo, ConstantLongInfo, ConstantFloatInfo, ConstantDoubleInfo:
			fmt.Printf("constant[%d] = %v (%T)\n", i, v, v)
		case ConstantUtf8Info:
			fmt.Printf("constant[%d] = \"%v\"\n", i, v)
		default:
			fmt.Printf("constant[%d] = %#v\n", i, v)
		}
		pool[i] = con
		if isDoubleLength(con) {
			pool[i+1] = ConstantUnusableInfo{} // 4.4.5 "The constant_pool index n+1 must be valid but is considered unusable"
			i++
		}
	}
	return pool, nil
}

func parseConstant(scn *BinScanner) (cp CpInfo, err error) {
	switch scn.Uint8() {
	case constClass:
		cp = ConstantClassInfo{scn.Uint16()}
	case constString:
		cp = ConstantStringInfo{scn.Uint16()}
	case constInteger:
		cp = scanConstantIntegerInfo(scn)
	case constFloat:
		cp = scanConstantFloatInfo(scn)
	case constLong:
		cp = scanConstantLongInfo(scn)
	case constDouble:
		cp = scanConstantDoubleInfo(scn)
	case constUtf8:
		cp = scanConstantUtf8Info(scn)
	case constFieldRef:
		cp = ConstantFieldRefInfo{scn.Uint16(), scn.Uint16()}
	case constMethodRef:
		cp = ConstantMethodRefInfo{scn.Uint16(), scn.Uint16()}
	case constInterfaceMethodRef:
		cp = ConstantInterfaceMethodRefInfo{scn.Uint16(), scn.Uint16()}
	case constNameAndType:
		cp = ConstantNameAndTypeInfo{scn.Uint16(), scn.Uint16()}
	default:
		err = errors.New("invalid constant kind")
	}
	return
}

type CpInfo interface {
	constantMarkerFunc()
}

func isDoubleLength(info CpInfo) bool {
	switch info.(type) {
	case ConstantLongInfo, ConstantDoubleInfo:
		return true
	}
	return false
}

type ConstantUnusableInfo struct{}

func (c ConstantUnusableInfo) constantMarkerFunc() {}

type ConstantClassInfo struct {
	NameIndex uint16 // Must be the index of a ConstantUtf8Info
}

func (c ConstantClassInfo) constantMarkerFunc() {}

type ConstantStringInfo struct {
	StringIndex uint16
}

func (c ConstantStringInfo) constantMarkerFunc() {}

type ConstantIntegerInfo jvm.Jint

func (c ConstantIntegerInfo) constantMarkerFunc() {}

func scanConstantIntegerInfo(scn *BinScanner) ConstantIntegerInfo {
	return ConstantIntegerInfo(scn.Uint32())
}

type ConstantFloatInfo jvm.Jfloat

func (c ConstantFloatInfo) constantMarkerFunc() {}

func scanConstantFloatInfo(scn *BinScanner) ConstantFloatInfo {
	return ConstantFloatInfo(math.Float32frombits(scn.Uint32()))
}

type ConstantLongInfo jvm.Jlong

func (c ConstantLongInfo) constantMarkerFunc() {}

func scanConstantLongInfo(scn *BinScanner) ConstantLongInfo {
	return ConstantLongInfo(scn.Uint64())
}

type ConstantDoubleInfo jvm.Jdouble

func (c ConstantDoubleInfo) constantMarkerFunc() {}

func scanConstantDoubleInfo(scn *BinScanner) ConstantDoubleInfo {
	return ConstantDoubleInfo(math.Float64frombits(scn.Uint64()))
}

type ConstantUtf8Info string

func (c ConstantUtf8Info) constantMarkerFunc() {}

func scanConstantUtf8Info(scn *BinScanner) ConstantUtf8Info {
	return ConstantUtf8Info(scn.Bytes(int(scn.Uint16())))
}

type ConstantFieldRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantFieldRefInfo) constantMarkerFunc() {}

type ConstantMethodRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantMethodRefInfo) constantMarkerFunc() {}

type ConstantInterfaceMethodRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantInterfaceMethodRefInfo) constantMarkerFunc() {}

type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c ConstantNameAndTypeInfo) constantMarkerFunc() {}
