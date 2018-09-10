package jvm

import (
	"github.com/angusholder/jvm-go/opcodes"
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

func (i *Jvm) Interpret(opcode opcodes.Opcode) {
	switch opcode {
	case opcodes.IconstM1:
		i.pushi(-1)
	case opcodes.Iconst0:
		i.pushi(0)
	case opcodes.Iconst1:
		i.pushi(1)
	case opcodes.Iconst2:
		i.pushi(2)
	case opcodes.Iconst3:
		i.pushi(3)
	case opcodes.Iconst4:
		i.pushi(4)
	case opcodes.Iconst5:
		i.pushi(5)

	case opcodes.Lconst0:
		i.pushl(0)
	case opcodes.Lconst1:
		i.pushl(1)

	case opcodes.Iadd:
		i.op2I(func(a, b Jint) Jint { return a + b })
	case opcodes.Isub:
		i.op2I(func(a, b Jint) Jint { return a - b })
	case opcodes.Imul:
		i.op2I(func(a, b Jint) Jint { return a * b })
	case opcodes.Idiv:
		i.op2I(func(a, b Jint) Jint { return a / b })
	case opcodes.Irem:
		i.op2I(func(a, b Jint) Jint { return a % b })
	case opcodes.Ineg:
		i.pushi(-i.popi())
	case opcodes.Iand:
		i.op2I(func(a, b Jint) Jint { return a & b })
	case opcodes.Ior:
		i.op2I(func(a, b Jint) Jint { return a | b })
	case opcodes.Ixor:
		i.op2I(func(a, b Jint) Jint { return a ^ b })
	case opcodes.Ishl:
		i.op2I(func(a, b Jint) Jint { return a << uint8(b) })
	case opcodes.Ishr:
		i.op2I(func(a, b Jint) Jint { return a >> uint8(b) })
	case opcodes.Iushr:
		i.op2I(func(a, b Jint) Jint { return Jint(uint32(a) >> uint32(b)) })

	case opcodes.Ladd:
		i.op2L(func(a, b Jlong) Jlong { return a + b })
	case opcodes.Lsub:
		i.op2L(func(a, b Jlong) Jlong { return a - b })
	case opcodes.Lmul:
		i.op2L(func(a, b Jlong) Jlong { return a * b })
	case opcodes.Ldiv:
		i.op2L(func(a, b Jlong) Jlong { return a / b })
	case opcodes.Lrem:
		i.op2L(func(a, b Jlong) Jlong { return a % b })
	case opcodes.Lneg:
		i.pushl(-i.popl())
	case opcodes.Land:
		i.op2L(func(a, b Jlong) Jlong { return a & b })
	case opcodes.Lor:
		i.op2L(func(a, b Jlong) Jlong { return a | b })
	case opcodes.Lxor:
		i.op2L(func(a, b Jlong) Jlong { return a ^ b })
	case opcodes.Lshl:
		i.op2L(func(a, b Jlong) Jlong { return a << uint8(b) })
	case opcodes.Lshr:
		i.op2L(func(a, b Jlong) Jlong { return a >> uint8(b) })
	case opcodes.Lushr:
		i.op2L(func(a, b Jlong) Jlong { return Jlong(uint64(a) >> uint64(b)) })

	}
}
