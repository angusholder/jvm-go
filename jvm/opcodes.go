package jvm

type Opcode uint8

const (
	// arrayref, index → value
	// load onto the stack a reference from an array
	OpcodeAaload Opcode = 0x32

	// arrayref, index, value →
	// store into a reference in an array
	OpcodeAastore Opcode = 0x53

	// → null
	// push a null reference onto the stack
	OpcodeAconstNull Opcode = 0x01

	// 1: index
	// → objectref
	// load a reference onto the stack from a local variable #index
	OpcodeAload Opcode = 0x19

	// → objectref
	// load a reference onto the stack from local variable 0
	OpcodeAload0 Opcode = 0x2a

	// → objectref
	// load a reference onto the stack from local variable 1
	OpcodeAload1 Opcode = 0x2b

	// → objectref
	// load a reference onto the stack from local variable 2
	OpcodeAload2 Opcode = 0x2c

	// → objectref
	// load a reference onto the stack from local variable 3
	OpcodeAload3 Opcode = 0x2d

	// 2: indexbyte1, indexbyte2
	// count → arrayref
	// create a new array of references of length count and component type identified by the class reference index (.mw-parser-output .monospaced{font-family:monospace,monospace}indexbyte1 << 8 + indexbyte2) in the constant pool
	OpcodeAnewarray Opcode = 0xbd

	// objectref → [empty]
	// return a reference from a method
	OpcodeAreturn Opcode = 0xb0

	// arrayref → length
	// get the length of an array
	OpcodeArraylength Opcode = 0xbe

	// 1: index
	// objectref →
	// store a reference into a local variable #index
	OpcodeAstore Opcode = 0x3a

	// objectref →
	// store a reference into local variable 0
	OpcodeAstore0 Opcode = 0x4b

	// objectref →
	// store a reference into local variable 1
	OpcodeAstore1 Opcode = 0x4c

	// objectref →
	// store a reference into local variable 2
	OpcodeAstore2 Opcode = 0x4d

	// objectref →
	// store a reference into local variable 3
	OpcodeAstore3 Opcode = 0x4e

	// objectref → [empty], objectref
	// throws an error or exception (notice that the rest of the stack is cleared, leaving only a reference to the Throwable)
	OpcodeAthrow Opcode = 0xbf

	// arrayref, index → value
	// load a byte or Boolean value from an array
	OpcodeBaload Opcode = 0x33

	// arrayref, index, value →
	// store a byte or Boolean value into an array
	OpcodeBastore Opcode = 0x54

	// 1: byte
	// → value
	// push a byte onto the stack as an integer value
	OpcodeBipush Opcode = 0x10

	//
	// reserved for breakpoints in Java debuggers; should not appear in any class file
	OpcodeBreakpoint Opcode = 0xca

	// arrayref, index → value
	// load a char from an array
	OpcodeCaload Opcode = 0x34

	// arrayref, index, value →
	// store a char into an array
	OpcodeCastore Opcode = 0x55

	// 2: indexbyte1, indexbyte2
	// objectref → objectref
	// checks whether an objectref is of a certain type, the class reference of which is in the constant pool at index (indexbyte1 << 8 + indexbyte2)
	OpcodeCheckcast Opcode = 0xc0

	// value → result
	// convert a double to a float
	OpcodeD2F Opcode = 0x90

	// value → result
	// convert a double to an int
	OpcodeD2I Opcode = 0x8e

	// value → result
	// convert a double to a long
	OpcodeD2L Opcode = 0x8f

	// value1, value2 → result
	// add two doubles
	OpcodeDadd Opcode = 0x63

	// arrayref, index → value
	// load a double from an array
	OpcodeDaload Opcode = 0x31

	// arrayref, index, value →
	// store a double into an array
	OpcodeDastore Opcode = 0x52

	// value1, value2 → result
	// compare two doubles
	OpcodeDcmpg Opcode = 0x98

	// value1, value2 → result
	// compare two doubles
	OpcodeDcmpl Opcode = 0x97

	// → 0.0
	// push the constant 0.0 (a double) onto the stack
	OpcodeDconst0 Opcode = 0x0e

	// → 1.0
	// push the constant 1.0 (a double) onto the stack
	OpcodeDconst1 Opcode = 0x0f

	// value1, value2 → result
	// divide two doubles
	OpcodeDdiv Opcode = 0x6f

	// 1: index
	// → value
	// load a double value from a local variable #index
	OpcodeDload Opcode = 0x18

	// → value
	// load a double from local variable 0
	OpcodeDload0 Opcode = 0x26

	// → value
	// load a double from local variable 1
	OpcodeDload1 Opcode = 0x27

	// → value
	// load a double from local variable 2
	OpcodeDload2 Opcode = 0x28

	// → value
	// load a double from local variable 3
	OpcodeDload3 Opcode = 0x29

	// value1, value2 → result
	// multiply two doubles
	OpcodeDmul Opcode = 0x6b

	// value → result
	// negate a double
	OpcodeDneg Opcode = 0x77

	// value1, value2 → result
	// get the remainder from a division between two doubles
	OpcodeDrem Opcode = 0x73

	// value → [empty]
	// return a double from a method
	OpcodeDreturn Opcode = 0xaf

	// 1: index
	// value →
	// store a double value into a local variable #index
	OpcodeDstore Opcode = 0x39

	// value →
	// store a double into local variable 0
	OpcodeDstore0 Opcode = 0x47

	// value →
	// store a double into local variable 1
	OpcodeDstore1 Opcode = 0x48

	// value →
	// store a double into local variable 2
	OpcodeDstore2 Opcode = 0x49

	// value →
	// store a double into local variable 3
	OpcodeDstore3 Opcode = 0x4a

	// value1, value2 → result
	// subtract a double from another
	OpcodeDsub Opcode = 0x67

	// value → value, value
	// duplicate the value on top of the stack
	OpcodeDup Opcode = 0x59

	// value2, value1 → value1, value2, value1
	// insert a copy of the top value into the stack two values from the top. value1 and value2 must not be of the type double or long.
	OpcodeDupX1 Opcode = 0x5a

	// value3, value2, value1 → value1, value3, value2, value1
	// insert a copy of the top value into the stack two (if value2 is double or long it takes up the entry of value3, too) or three values (if value2 is neither double nor long) from the top
	OpcodeDupX2 Opcode = 0x5b

	// {value2, value1} → {value2, value1}, {value2, value1}
	// duplicate top two stack words (two values, if value1 is not double nor long; a single value, if value1 is double or long)
	OpcodeDup2 Opcode = 0x5c

	// value3, {value2, value1} → {value2, value1}, value3, {value2, value1}
	// duplicate two words and insert beneath third word (see explanation above)
	OpcodeDup2X1 Opcode = 0x5d

	// {value4, value3}, {value2, value1} → {value2, value1}, {value4, value3}, {value2, value1}
	// duplicate two words and insert beneath fourth word
	OpcodeDup2X2 Opcode = 0x5e

	// value → result
	// convert a float to a double
	OpcodeF2D Opcode = 0x8d

	// value → result
	// convert a float to an int
	OpcodeF2I Opcode = 0x8b

	// value → result
	// convert a float to a long
	OpcodeF2L Opcode = 0x8c

	// value1, value2 → result
	// add two floats
	OpcodeFadd Opcode = 0x62

	// arrayref, index → value
	// load a float from an array
	OpcodeFaload Opcode = 0x30

	// arrayref, index, value →
	// store a float in an array
	OpcodeFastore Opcode = 0x51

	// value1, value2 → result
	// compare two floats
	OpcodeFcmpg Opcode = 0x96

	// value1, value2 → result
	// compare two floats
	OpcodeFcmpl Opcode = 0x95

	// → 0.0f
	// push 0.0f on the stack
	OpcodeFconst0 Opcode = 0x0b

	// → 1.0f
	// push 1.0f on the stack
	OpcodeFconst1 Opcode = 0x0c

	// → 2.0f
	// push 2.0f on the stack
	OpcodeFconst2 Opcode = 0x0d

	// value1, value2 → result
	// divide two floats
	OpcodeFdiv Opcode = 0x6e

	// 1: index
	// → value
	// load a float value from a local variable #index
	OpcodeFload Opcode = 0x17

	// → value
	// load a float value from local variable 0
	OpcodeFload0 Opcode = 0x22

	// → value
	// load a float value from local variable 1
	OpcodeFload1 Opcode = 0x23

	// → value
	// load a float value from local variable 2
	OpcodeFload2 Opcode = 0x24

	// → value
	// load a float value from local variable 3
	OpcodeFload3 Opcode = 0x25

	// value1, value2 → result
	// multiply two floats
	OpcodeFmul Opcode = 0x6a

	// value → result
	// negate a float
	OpcodeFneg Opcode = 0x76

	// value1, value2 → result
	// get the remainder from a division between two floats
	OpcodeFrem Opcode = 0x72

	// value → [empty]
	// return a float
	OpcodeFreturn Opcode = 0xae

	// 1: index
	// value →
	// store a float value into a local variable #index
	OpcodeFstore Opcode = 0x38

	// value →
	// store a float value into local variable 0
	OpcodeFstore0 Opcode = 0x43

	// value →
	// store a float value into local variable 1
	OpcodeFstore1 Opcode = 0x44

	// value →
	// store a float value into local variable 2
	OpcodeFstore2 Opcode = 0x45

	// value →
	// store a float value into local variable 3
	OpcodeFstore3 Opcode = 0x46

	// value1, value2 → result
	// subtract two floats
	OpcodeFsub Opcode = 0x66

	// 2: indexbyte1, indexbyte2
	// objectref → value
	// get a field value of an object objectref, where the field is identified by field reference in the constant pool index (indexbyte1 << 8 + indexbyte2)
	OpcodeGetfield Opcode = 0xb4

	// 2: indexbyte1, indexbyte2
	// → value
	// get a static field value of a class, where the field is identified by field reference in the constant pool index (indexbyte1 << 8 + indexbyte2)
	OpcodeGetstatic Opcode = 0xb2

	// 2: branchbyte1, branchbyte2
	// [no change]
	// goes to another instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeGoto Opcode = 0xa7

	// 4: branchbyte1, branchbyte2, branchbyte3, branchbyte4
	// [no change]
	// goes to another instruction at branchoffset (signed int constructed from unsigned bytes branchbyte1 << 24 + branchbyte2 << 16 + branchbyte3 << 8 + branchbyte4)
	OpcodeGotoW Opcode = 0xc8

	// value → result
	// convert an int into a byte
	OpcodeI2B Opcode = 0x91

	// value → result
	// convert an int into a character
	OpcodeI2C Opcode = 0x92

	// value → result
	// convert an int into a double
	OpcodeI2D Opcode = 0x87

	// value → result
	// convert an int into a float
	OpcodeI2F Opcode = 0x86

	// value → result
	// convert an int into a long
	OpcodeI2L Opcode = 0x85

	// value → result
	// convert an int into a short
	OpcodeI2S Opcode = 0x93

	// value1, value2 → result
	// add two ints
	OpcodeIadd Opcode = 0x60

	// arrayref, index → value
	// load an int from an array
	OpcodeIaload Opcode = 0x2e

	// value1, value2 → result
	// perform a bitwise AND on two integers
	OpcodeIand Opcode = 0x7e

	// arrayref, index, value →
	// store an int into an array
	OpcodeIastore Opcode = 0x4f

	// → -1
	// load the int value −1 onto the stack
	OpcodeIconstM1 Opcode = 0x02

	// → 0
	// load the int value 0 onto the stack
	OpcodeIconst0 Opcode = 0x03

	// → 1
	// load the int value 1 onto the stack
	OpcodeIconst1 Opcode = 0x04

	// → 2
	// load the int value 2 onto the stack
	OpcodeIconst2 Opcode = 0x05

	// → 3
	// load the int value 3 onto the stack
	OpcodeIconst3 Opcode = 0x06

	// → 4
	// load the int value 4 onto the stack
	OpcodeIconst4 Opcode = 0x07

	// → 5
	// load the int value 5 onto the stack
	OpcodeIconst5 Opcode = 0x08

	// value1, value2 → result
	// divide two integers
	OpcodeIdiv Opcode = 0x6c

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if references are equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfAcmpeq Opcode = 0xa5

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if references are not equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfAcmpne Opcode = 0xa6

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if ints are equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmpeq Opcode = 0x9f

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is greater than or equal to value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmpge Opcode = 0xa2

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is greater than value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmpgt Opcode = 0xa3

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is less than or equal to value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmple Opcode = 0xa4

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is less than value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmplt Opcode = 0xa1

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if ints are not equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfIcmpne Opcode = 0xa0

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfeq Opcode = 0x99

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is greater than or equal to 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfge Opcode = 0x9c

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is greater than 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfgt Opcode = 0x9d

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is less than or equal to 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfle Opcode = 0x9e

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is less than 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIflt Opcode = 0x9b

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is not 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfne Opcode = 0x9a

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is not null, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfnonnull Opcode = 0xc7

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is null, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	OpcodeIfnull Opcode = 0xc6

	// 2: index, const
	// [No change]
	// increment local variable #index by signed byte const
	OpcodeIinc Opcode = 0x84

	// 1: index
	// → value
	// load an int value from a local variable #index
	OpcodeIload Opcode = 0x15

	// → value
	// load an int value from local variable 0
	OpcodeIload0 Opcode = 0x1a

	// → value
	// load an int value from local variable 1
	OpcodeIload1 Opcode = 0x1b

	// → value
	// load an int value from local variable 2
	OpcodeIload2 Opcode = 0x1c

	// → value
	// load an int value from local variable 3
	OpcodeIload3 Opcode = 0x1d

	//
	// reserved for implementation-dependent operations within debuggers; should not appear in any class file
	OpcodeImpdep1 Opcode = 0xfe

	//
	// reserved for implementation-dependent operations within debuggers; should not appear in any class file
	OpcodeImpdep2 Opcode = 0xff

	// value1, value2 → result
	// multiply two integers
	OpcodeImul Opcode = 0x68

	// value → result
	// negate int
	OpcodeIneg Opcode = 0x74

	// 2: indexbyte1, indexbyte2
	// objectref → result
	// determines if an object objectref is of a given type, identified by class reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInstanceof Opcode = 0xc1

	// 4: indexbyte1, indexbyte2, 0, 0
	// [arg1, [arg2 ...]] → result
	// invokes a dynamic method and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInvokedynamic Opcode = 0xba

	// 4: indexbyte1, indexbyte2, count, 0
	// objectref, [arg1, arg2, ...] → result
	// invokes an interface method on object objectref and puts the result on the stack (might be void); the interface method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInvokeinterface Opcode = 0xb9

	// 2: indexbyte1, indexbyte2
	// objectref, [arg1, arg2, ...] → result
	// invoke instance method on object objectref and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInvokespecial Opcode = 0xb7

	// 2: indexbyte1, indexbyte2
	// [arg1, arg2, ...] → result
	// invoke a static method and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInvokestatic Opcode = 0xb8

	// 2: indexbyte1, indexbyte2
	// objectref, [arg1, arg2, ...] → result
	// invoke virtual method on object objectref and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodeInvokevirtual Opcode = 0xb6

	// value1, value2 → result
	// bitwise int OR
	OpcodeIor Opcode = 0x80

	// value1, value2 → result
	// logical int remainder
	OpcodeIrem Opcode = 0x70

	// value → [empty]
	// return an integer from a method
	OpcodeIreturn Opcode = 0xac

	// value1, value2 → result
	// int shift left
	OpcodeIshl Opcode = 0x78

	// value1, value2 → result
	// int arithmetic shift right
	OpcodeIshr Opcode = 0x7a

	// 1: index
	// value →
	// store int value into variable #index
	OpcodeIstore Opcode = 0x36

	// value →
	// store int value into variable 0
	OpcodeIstore0 Opcode = 0x3b

	// value →
	// store int value into variable 1
	OpcodeIstore1 Opcode = 0x3c

	// value →
	// store int value into variable 2
	OpcodeIstore2 Opcode = 0x3d

	// value →
	// store int value into variable 3
	OpcodeIstore3 Opcode = 0x3e

	// value1, value2 → result
	// int subtract
	OpcodeIsub Opcode = 0x64

	// value1, value2 → result
	// int logical shift right
	OpcodeIushr Opcode = 0x7c

	// value1, value2 → result
	// int xor
	OpcodeIxor Opcode = 0x82

	// 2: branchbyte1, branchbyte2
	// → address
	// jump to subroutine at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2) and place the return address on the stack
	OpcodeJsr Opcode = 0xa8

	// 4: branchbyte1, branchbyte2, branchbyte3, branchbyte4
	// → address
	// jump to subroutine at branchoffset (signed int constructed from unsigned bytes branchbyte1 << 24 + branchbyte2 << 16 + branchbyte3 << 8 + branchbyte4) and place the return address on the stack
	OpcodeJsrW Opcode = 0xc9

	// value → result
	// convert a long to a double
	OpcodeL2D Opcode = 0x8a

	// value → result
	// convert a long to a float
	OpcodeL2F Opcode = 0x89

	// value → result
	// convert a long to a int
	OpcodeL2I Opcode = 0x88

	// value1, value2 → result
	// add two longs
	OpcodeLadd Opcode = 0x61

	// arrayref, index → value
	// load a long from an array
	OpcodeLaload Opcode = 0x2f

	// value1, value2 → result
	// bitwise AND of two longs
	OpcodeLand Opcode = 0x7f

	// arrayref, index, value →
	// store a long to an array
	OpcodeLastore Opcode = 0x50

	// value1, value2 → result
	// push 0 if the two longs are the same, 1 if value1 is greater than value2, -1 otherwise
	OpcodeLcmp Opcode = 0x94

	// → 0L
	// push 0L (the number zero with type long) onto the stack
	OpcodeLconst0 Opcode = 0x09

	// → 1L
	// push 1L (the number one with type long) onto the stack
	OpcodeLconst1 Opcode = 0x0a

	// 1: index
	// → value
	// push a constant #index from a constant pool (String, int, float, Class, java.lang.invoke.MethodType, or java.lang.invoke.MethodHandle) onto the stack
	OpcodeLdc Opcode = 0x12

	// 2: indexbyte1, indexbyte2
	// → value
	// push a constant #index from a constant pool (String, int, float, Class, java.lang.invoke.MethodType, or java.lang.invoke.MethodHandle) onto the stack (wide index is constructed as indexbyte1 << 8 + indexbyte2)
	OpcodeLdcW Opcode = 0x13

	// 2: indexbyte1, indexbyte2
	// → value
	// push a constant #index from a constant pool (double or long) onto the stack (wide index is constructed as indexbyte1 << 8 + indexbyte2)
	OpcodeLdc2W Opcode = 0x14

	// value1, value2 → result
	// divide two longs
	OpcodeLdiv Opcode = 0x6d

	// 1: index
	// → value
	// load a long value from a local variable #index
	OpcodeLload Opcode = 0x16

	// → value
	// load a long value from a local variable 0
	OpcodeLload0 Opcode = 0x1e

	// → value
	// load a long value from a local variable 1
	OpcodeLload1 Opcode = 0x1f

	// → value
	// load a long value from a local variable 2
	OpcodeLload2 Opcode = 0x20

	// → value
	// load a long value from a local variable 3
	OpcodeLload3 Opcode = 0x21

	// value1, value2 → result
	// multiply two longs
	OpcodeLmul Opcode = 0x69

	// value → result
	// negate a long
	OpcodeLneg Opcode = 0x75

	// 8+: <0–3 bytes padding>, defaultbyte1, defaultbyte2, defaultbyte3, defaultbyte4, npairs1, npairs2, npairs3, npairs4, match-offset pairs...
	// key →
	// a target address is looked up from a table using a key and execution continues from the instruction at that address
	OpcodeLookupswitch Opcode = 0xab

	// value1, value2 → result
	// bitwise OR of two longs
	OpcodeLor Opcode = 0x81

	// value1, value2 → result
	// remainder of division of two longs
	OpcodeLrem Opcode = 0x71

	// value → [empty]
	// return a long value
	OpcodeLreturn Opcode = 0xad

	// value1, value2 → result
	// bitwise shift left of a long value1 by int value2 positions
	OpcodeLshl Opcode = 0x79

	// value1, value2 → result
	// bitwise shift right of a long value1 by int value2 positions
	OpcodeLshr Opcode = 0x7b

	// 1: index
	// value →
	// store a long value in a local variable #index
	OpcodeLstore Opcode = 0x37

	// value →
	// store a long value in a local variable 0
	OpcodeLstore0 Opcode = 0x3f

	// value →
	// store a long value in a local variable 1
	OpcodeLstore1 Opcode = 0x40

	// value →
	// store a long value in a local variable 2
	OpcodeLstore2 Opcode = 0x41

	// value →
	// store a long value in a local variable 3
	OpcodeLstore3 Opcode = 0x42

	// value1, value2 → result
	// subtract two longs
	OpcodeLsub Opcode = 0x65

	// value1, value2 → result
	// bitwise shift right of a long value1 by int value2 positions, unsigned
	OpcodeLushr Opcode = 0x7d

	// value1, value2 → result
	// bitwise XOR of two longs
	OpcodeLxor Opcode = 0x83

	// objectref →
	// enter monitor for object (\"grab the lock\" – start of synchronized() section)
	OpcodeMonitorenter Opcode = 0xc2

	// objectref →
	// exit monitor for object (\"release the lock\" – end of synchronized() section)
	OpcodeMonitorexit Opcode = 0xc3

	// 3: indexbyte1, indexbyte2, dimensions
	// count1, [count2,...] → arrayref
	// create a new array of dimensions dimensions of type identified by class reference in constant pool index (indexbyte1 << 8 + indexbyte2); the sizes of each dimension is identified by count1, [count2, etc.]
	OpcodeMultianewarray Opcode = 0xc5

	// 2: indexbyte1, indexbyte2
	// → objectref
	// create new object of type identified by class reference in constant pool index (indexbyte1 << 8 + indexbyte2)
	OpcodeNew Opcode = 0xbb

	// 1: atype
	// count → arrayref
	// create new array with count elements of primitive type identified by atype
	OpcodeNewarray Opcode = 0xbc

	// [No change]
	// perform no operation
	OpcodeNop Opcode = 0x00

	// value →
	// discard the top value on the stack
	OpcodePop Opcode = 0x57

	// {value2, value1} →
	// discard the top two values on the stack (or one value, if it is a double or long)
	OpcodePop2 Opcode = 0x58

	// 2: indexbyte1, indexbyte2
	// objectref, value →
	// set field to value in an object objectref, where the field is identified by a field reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodePutfield Opcode = 0xb5

	// 2: indexbyte1, indexbyte2
	// value →
	// set static field to value in a class, where the field is identified by a field reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	OpcodePutstatic Opcode = 0xb3

	// 1: index
	// [No change]
	// continue execution from address taken from a local variable #index (the asymmetry with jsr is intentional)
	OpcodeRet Opcode = 0xa9

	// → [empty]
	// return void from method
	OpcodeReturn Opcode = 0xb1

	// arrayref, index → value
	// load short from array
	OpcodeSaload Opcode = 0x35

	// arrayref, index, value →
	// store short to array
	OpcodeSastore Opcode = 0x56

	// 2: byte1, byte2
	// → value
	// push a short onto the stack as an integer value
	OpcodeSipush Opcode = 0x11

	// value2, value1 → value1, value2
	// swaps two top words on the stack (note that value1 and value2 must not be double or long)
	OpcodeSwap Opcode = 0x5f

	// 16+: [0–3 bytes padding], defaultbyte1, defaultbyte2, defaultbyte3, defaultbyte4, lowbyte1, lowbyte2, lowbyte3, lowbyte4, highbyte1, highbyte2, highbyte3, highbyte4, jump offsets...
	// index →
	// continue execution from an address in the table at offset index
	OpcodeTableswitch Opcode = 0xaa

	// 3/5: opcode, indexbyte1, indexbyte2or iinc, indexbyte1, indexbyte2, countbyte1, countbyte2
	// [same as for corresponding instructions]
	// execute opcode, where opcode is either iload, fload, aload, lload, dload, istore, fstore, astore, lstore, dstore, or ret, but assume the index is 16 bit; or execute iinc, where the index is 16 bits and the constant to increment by is a signed 16 bit short
	OpcodeWide Opcode = 0xc4
)
