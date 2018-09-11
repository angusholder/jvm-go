package jvm

import (
	"encoding/binary"
	"math"
)

type Jint = int32
type Jlong = int64
type Jfloat = float32
type Jdouble = float64

type Jvm struct {
	code          []uint8
	codeOffset    uint
	locals        []uint64
	valueStack    []Jint
	valueStackTop uint
	wide          bool
}

func (i *Jvm) pushi(val Jint) {
	i.valueStack[i.valueStackTop] = val
	i.valueStackTop++
}
func (i *Jvm) popi() Jint {
	i.valueStackTop--
	return i.valueStack[i.valueStackTop]
}
func (i *Jvm) pushl(val Jlong) {
	i.pushi(Jint(val))
	i.pushi(Jint(val >> 32))
}
func (i *Jvm) popl() Jlong {
	var result Jlong
	result |= Jlong(i.popi()) << 32
	result |= Jlong(i.popi())
	return result
}
func (i *Jvm) pushf(val Jfloat) {
	i.pushi(Jint(math.Float32bits(val)))
}
func (i *Jvm) popf() Jfloat {
	return math.Float32frombits(uint32(i.popi()))
}
func (i *Jvm) pushd(val Jdouble) {
	i.pushl(Jlong(math.Float64bits(val)))
}
func (i *Jvm) popd() Jdouble {
	return math.Float64frombits(uint64(i.popl()))
}

func (i *Jvm) op2I(action func(Jint, Jint) Jint) {
	a := i.popi()
	b := i.popi()
	result := action(a, b)
	i.pushi(result)
}

func (i *Jvm) op2L(action func(Jlong, Jlong) Jlong) {
	a := i.popl()
	b := i.popl()
	result := action(a, b)
	i.pushl(result)
}

func (i *Jvm) op2F(action func(Jfloat, Jfloat) Jfloat) {
	a := i.popf()
	b := i.popf()
	result := action(a, b)
	i.pushf(result)
}

func (i *Jvm) op2D(action func(Jdouble, Jdouble) Jdouble) {
	a := i.popd()
	b := i.popd()
	result := action(a, b)
	i.pushd(result)
}

func (i *Jvm) loadLocalI(offset int) { i.pushi(Jint(i.locals[offset])) }
func (i *Jvm) loadLocalL(offset int) { i.pushl(Jlong(i.locals[offset])) }
func (i *Jvm) loadLocalF(offset int) { i.pushf(math.Float32frombits(uint32(i.locals[offset]))) }
func (i *Jvm) loadLocalD(offset int) { i.pushd(math.Float64frombits(i.locals[offset])) }

func (i *Jvm) storeLocalI(offset int) { i.locals[offset] = uint64(i.popi()) }
func (i *Jvm) storeLocalL(offset int) { i.locals[offset] = uint64(i.popl()) }
func (i *Jvm) storeLocalF(offset int) { i.locals[offset] = uint64(math.Float32bits(i.popf())) }
func (i *Jvm) storeLocalD(offset int) { i.locals[offset] = math.Float64bits(i.popd()) }

func (i *Jvm) getOpcodeOffset() int {
	var result int
	if i.wide {
		result = int(binary.BigEndian.Uint16(i.code[i.codeOffset:]))
		i.codeOffset += 2
	} else {
		result = int(i.code[i.codeOffset])
		i.codeOffset += 1
	}
	return result
}

func (i *Jvm) Interpret() {
	opcode := Opcode(i.code[i.codeOffset])
	i.codeOffset++
	switch opcode {
	case OpcodeIconstM1:
		i.pushi(-1)
	case OpcodeIconst0:
		i.pushi(0)
	case OpcodeIconst1:
		i.pushi(1)
	case OpcodeIconst2:
		i.pushi(2)
	case OpcodeIconst3:
		i.pushi(3)
	case OpcodeIconst4:
		i.pushi(4)
	case OpcodeIconst5:
		i.pushi(5)

	case OpcodeLconst0:
		i.pushl(0)
	case OpcodeLconst1:
		i.pushl(1)

	case OpcodeIadd:
		i.op2I(func(a, b Jint) Jint { return a + b })
	case OpcodeIsub:
		i.op2I(func(a, b Jint) Jint { return a - b })
	case OpcodeImul:
		i.op2I(func(a, b Jint) Jint { return a * b })
	case OpcodeIdiv:
		i.op2I(func(a, b Jint) Jint { return a / b })
	case OpcodeIrem:
		i.op2I(func(a, b Jint) Jint { return a % b })
	case OpcodeIneg:
		i.pushi(-i.popi())
	case OpcodeIand:
		i.op2I(func(a, b Jint) Jint { return a & b })
	case OpcodeIor:
		i.op2I(func(a, b Jint) Jint { return a | b })
	case OpcodeIxor:
		i.op2I(func(a, b Jint) Jint { return a ^ b })
	case OpcodeIshl:
		i.op2I(func(a, b Jint) Jint { return a << uint8(b) })
	case OpcodeIshr:
		i.op2I(func(a, b Jint) Jint { return a >> uint8(b) })
	case OpcodeIushr:
		i.op2I(func(a, b Jint) Jint { return Jint(uint32(a) >> uint32(b)) })

	case OpcodeLadd:
		i.op2L(func(a, b Jlong) Jlong { return a + b })
	case OpcodeLsub:
		i.op2L(func(a, b Jlong) Jlong { return a - b })
	case OpcodeLmul:
		i.op2L(func(a, b Jlong) Jlong { return a * b })
	case OpcodeLdiv:
		i.op2L(func(a, b Jlong) Jlong { return a / b })
	case OpcodeLrem:
		i.op2L(func(a, b Jlong) Jlong { return a % b })
	case OpcodeLneg:
		i.pushl(-i.popl())
	case OpcodeLand:
		i.op2L(func(a, b Jlong) Jlong { return a & b })
	case OpcodeLor:
		i.op2L(func(a, b Jlong) Jlong { return a | b })
	case OpcodeLxor:
		i.op2L(func(a, b Jlong) Jlong { return a ^ b })
	case OpcodeLshl:
		i.op2L(func(a, b Jlong) Jlong { return a << uint8(b) })
	case OpcodeLshr:
		i.op2L(func(a, b Jlong) Jlong { return a >> uint8(b) })
	case OpcodeLushr:
		i.op2L(func(a, b Jlong) Jlong { return Jlong(uint64(a) >> uint64(b)) })

	case OpcodeFadd:
		i.op2F(func(a, b Jfloat) Jfloat { return a + b })
	case OpcodeFsub:
		i.op2F(func(a, b Jfloat) Jfloat { return a - b })
	case OpcodeFmul:
		i.op2F(func(a, b Jfloat) Jfloat { return a * b })
	case OpcodeFdiv:
		i.op2F(func(a, b Jfloat) Jfloat { return a / b })
	case OpcodeFrem:
		i.op2F(func(a, b Jfloat) Jfloat { return Jfloat(math.Remainder(float64(a), float64(b))) })
	case OpcodeFneg:
		i.pushf(-i.popf())

	case OpcodeDadd:
		i.op2D(func(a, b Jdouble) Jdouble { return a + b })
	case OpcodeDsub:
		i.op2D(func(a, b Jdouble) Jdouble { return a - b })
	case OpcodeDmul:
		i.op2D(func(a, b Jdouble) Jdouble { return a * b })
	case OpcodeDdiv:
		i.op2D(func(a, b Jdouble) Jdouble { return a / b })
	case OpcodeDrem:
		i.op2D(func(a, b Jdouble) Jdouble { return math.Remainder(a, b) })
	case OpcodeDneg:
		i.pushd(-i.popd())

	case OpcodeIload0, OpcodeIload1, OpcodeIload2, OpcodeIload3:
		i.loadLocalI(int(opcode - OpcodeIload0))
	case OpcodeLload0, OpcodeLload1, OpcodeLload2, OpcodeLload3:
		i.loadLocalL(int(opcode - OpcodeLload0))
	case OpcodeFload0, OpcodeFload1, OpcodeFload2, OpcodeFload3:
		i.loadLocalF(int(opcode - OpcodeFload0))
	case OpcodeDload0, OpcodeDload1, OpcodeDload2, OpcodeDload3:
		i.loadLocalD(int(opcode - OpcodeDload0))

	case OpcodeIstore0, OpcodeIstore1, OpcodeIstore2, OpcodeIstore3:
		i.storeLocalI(int(opcode - OpcodeIstore0))
	case OpcodeLstore0, OpcodeLstore1, OpcodeLstore2, OpcodeLstore3:
		i.storeLocalL(int(opcode - OpcodeLstore0))
	case OpcodeFstore0, OpcodeFstore1, OpcodeFstore2, OpcodeFstore3:
		i.storeLocalF(int(opcode - OpcodeFstore0))
	case OpcodeDstore0, OpcodeDstore1, OpcodeDstore2, OpcodeDstore3:
		i.storeLocalD(int(opcode - OpcodeDstore0))

	case OpcodeIload:
		i.loadLocalI(i.getOpcodeOffset())
	case OpcodeLload:
		i.loadLocalL(i.getOpcodeOffset())
	case OpcodeFload:
		i.loadLocalF(i.getOpcodeOffset())
	case OpcodeDload:
		i.loadLocalD(i.getOpcodeOffset())

	case OpcodeIstore:
		i.storeLocalI(i.getOpcodeOffset())
	case OpcodeLstore:
		i.storeLocalL(i.getOpcodeOffset())
	case OpcodeFstore:
		i.storeLocalF(i.getOpcodeOffset())
	case OpcodeDstore:
		i.storeLocalD(i.getOpcodeOffset())

	}
}
