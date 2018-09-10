package jvm

import (
	"math"
)

type Jint = int32
type Jlong = int64
type Jfloat = float32
type Jdouble = float64

type Jvm struct {
	code          []uint8
	codeOffset    uint
	locals        []Jlong
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

func (i *Jvm) Interpret(opcode Opcode) {
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

	}
}
