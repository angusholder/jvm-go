package opcodes

type Opcode uint8

const (
	// arrayref, index → value
	// load onto the stack a reference from an array
	Aaload Opcode = 0x32

	// arrayref, index, value →
	// store into a reference in an array
	Aastore Opcode = 0x53

	// → null
	// push a null reference onto the stack
	AconstNull Opcode = 0x01

	// 1: index
	// → objectref
	// load a reference onto the stack from a local variable #index
	Aload Opcode = 0x19

	// → objectref
	// load a reference onto the stack from local variable 0
	Aload0 Opcode = 0x2a

	// → objectref
	// load a reference onto the stack from local variable 1
	Aload1 Opcode = 0x2b

	// → objectref
	// load a reference onto the stack from local variable 2
	Aload2 Opcode = 0x2c

	// → objectref
	// load a reference onto the stack from local variable 3
	Aload3 Opcode = 0x2d

	// 2: indexbyte1, indexbyte2
	// count → arrayref
	// create a new array of references of length count and component type identified by the class reference index (.mw-parser-output .monospaced{font-family:monospace,monospace}indexbyte1 << 8 + indexbyte2) in the constant pool
	Anewarray Opcode = 0xbd

	// objectref → [empty]
	// return a reference from a method
	Areturn Opcode = 0xb0

	// arrayref → length
	// get the length of an array
	Arraylength Opcode = 0xbe

	// 1: index
	// objectref →
	// store a reference into a local variable #index
	Astore Opcode = 0x3a

	// objectref →
	// store a reference into local variable 0
	Astore0 Opcode = 0x4b

	// objectref →
	// store a reference into local variable 1
	Astore1 Opcode = 0x4c

	// objectref →
	// store a reference into local variable 2
	Astore2 Opcode = 0x4d

	// objectref →
	// store a reference into local variable 3
	Astore3 Opcode = 0x4e

	// objectref → [empty], objectref
	// throws an error or exception (notice that the rest of the stack is cleared, leaving only a reference to the Throwable)
	Athrow Opcode = 0xbf

	// arrayref, index → value
	// load a byte or Boolean value from an array
	Baload Opcode = 0x33

	// arrayref, index, value →
	// store a byte or Boolean value into an array
	Bastore Opcode = 0x54

	// 1: byte
	// → value
	// push a byte onto the stack as an integer value
	Bipush Opcode = 0x10

	//
	// reserved for breakpoints in Java debuggers; should not appear in any class file
	Breakpoint Opcode = 0xca

	// arrayref, index → value
	// load a char from an array
	Caload Opcode = 0x34

	// arrayref, index, value →
	// store a char into an array
	Castore Opcode = 0x55

	// 2: indexbyte1, indexbyte2
	// objectref → objectref
	// checks whether an objectref is of a certain type, the class reference of which is in the constant pool at index (indexbyte1 << 8 + indexbyte2)
	Checkcast Opcode = 0xc0

	// value → result
	// convert a double to a float
	D2F Opcode = 0x90

	// value → result
	// convert a double to an int
	D2I Opcode = 0x8e

	// value → result
	// convert a double to a long
	D2L Opcode = 0x8f

	// value1, value2 → result
	// add two doubles
	Dadd Opcode = 0x63

	// arrayref, index → value
	// load a double from an array
	Daload Opcode = 0x31

	// arrayref, index, value →
	// store a double into an array
	Dastore Opcode = 0x52

	// value1, value2 → result
	// compare two doubles
	Dcmpg Opcode = 0x98

	// value1, value2 → result
	// compare two doubles
	Dcmpl Opcode = 0x97

	// → 0.0
	// push the constant 0.0 (a double) onto the stack
	Dconst0 Opcode = 0x0e

	// → 1.0
	// push the constant 1.0 (a double) onto the stack
	Dconst1 Opcode = 0x0f

	// value1, value2 → result
	// divide two doubles
	Ddiv Opcode = 0x6f

	// 1: index
	// → value
	// load a double value from a local variable #index
	Dload Opcode = 0x18

	// → value
	// load a double from local variable 0
	Dload0 Opcode = 0x26

	// → value
	// load a double from local variable 1
	Dload1 Opcode = 0x27

	// → value
	// load a double from local variable 2
	Dload2 Opcode = 0x28

	// → value
	// load a double from local variable 3
	Dload3 Opcode = 0x29

	// value1, value2 → result
	// multiply two doubles
	Dmul Opcode = 0x6b

	// value → result
	// negate a double
	Dneg Opcode = 0x77

	// value1, value2 → result
	// get the remainder from a division between two doubles
	Drem Opcode = 0x73

	// value → [empty]
	// return a double from a method
	Dreturn Opcode = 0xaf

	// 1: index
	// value →
	// store a double value into a local variable #index
	Dstore Opcode = 0x39

	// value →
	// store a double into local variable 0
	Dstore0 Opcode = 0x47

	// value →
	// store a double into local variable 1
	Dstore1 Opcode = 0x48

	// value →
	// store a double into local variable 2
	Dstore2 Opcode = 0x49

	// value →
	// store a double into local variable 3
	Dstore3 Opcode = 0x4a

	// value1, value2 → result
	// subtract a double from another
	Dsub Opcode = 0x67

	// value → value, value
	// duplicate the value on top of the stack
	Dup Opcode = 0x59

	// value2, value1 → value1, value2, value1
	// insert a copy of the top value into the stack two values from the top. value1 and value2 must not be of the type double or long.
	DupX1 Opcode = 0x5a

	// value3, value2, value1 → value1, value3, value2, value1
	// insert a copy of the top value into the stack two (if value2 is double or long it takes up the entry of value3, too) or three values (if value2 is neither double nor long) from the top
	DupX2 Opcode = 0x5b

	// {value2, value1} → {value2, value1}, {value2, value1}
	// duplicate top two stack words (two values, if value1 is not double nor long; a single value, if value1 is double or long)
	Dup2 Opcode = 0x5c

	// value3, {value2, value1} → {value2, value1}, value3, {value2, value1}
	// duplicate two words and insert beneath third word (see explanation above)
	Dup2X1 Opcode = 0x5d

	// {value4, value3}, {value2, value1} → {value2, value1}, {value4, value3}, {value2, value1}
	// duplicate two words and insert beneath fourth word
	Dup2X2 Opcode = 0x5e

	// value → result
	// convert a float to a double
	F2D Opcode = 0x8d

	// value → result
	// convert a float to an int
	F2I Opcode = 0x8b

	// value → result
	// convert a float to a long
	F2L Opcode = 0x8c

	// value1, value2 → result
	// add two floats
	Fadd Opcode = 0x62

	// arrayref, index → value
	// load a float from an array
	Faload Opcode = 0x30

	// arrayref, index, value →
	// store a float in an array
	Fastore Opcode = 0x51

	// value1, value2 → result
	// compare two floats
	Fcmpg Opcode = 0x96

	// value1, value2 → result
	// compare two floats
	Fcmpl Opcode = 0x95

	// → 0.0f
	// push 0.0f on the stack
	Fconst0 Opcode = 0x0b

	// → 1.0f
	// push 1.0f on the stack
	Fconst1 Opcode = 0x0c

	// → 2.0f
	// push 2.0f on the stack
	Fconst2 Opcode = 0x0d

	// value1, value2 → result
	// divide two floats
	Fdiv Opcode = 0x6e

	// 1: index
	// → value
	// load a float value from a local variable #index
	Fload Opcode = 0x17

	// → value
	// load a float value from local variable 0
	Fload0 Opcode = 0x22

	// → value
	// load a float value from local variable 1
	Fload1 Opcode = 0x23

	// → value
	// load a float value from local variable 2
	Fload2 Opcode = 0x24

	// → value
	// load a float value from local variable 3
	Fload3 Opcode = 0x25

	// value1, value2 → result
	// multiply two floats
	Fmul Opcode = 0x6a

	// value → result
	// negate a float
	Fneg Opcode = 0x76

	// value1, value2 → result
	// get the remainder from a division between two floats
	Frem Opcode = 0x72

	// value → [empty]
	// return a float
	Freturn Opcode = 0xae

	// 1: index
	// value →
	// store a float value into a local variable #index
	Fstore Opcode = 0x38

	// value →
	// store a float value into local variable 0
	Fstore0 Opcode = 0x43

	// value →
	// store a float value into local variable 1
	Fstore1 Opcode = 0x44

	// value →
	// store a float value into local variable 2
	Fstore2 Opcode = 0x45

	// value →
	// store a float value into local variable 3
	Fstore3 Opcode = 0x46

	// value1, value2 → result
	// subtract two floats
	Fsub Opcode = 0x66

	// 2: indexbyte1, indexbyte2
	// objectref → value
	// get a field value of an object objectref, where the field is identified by field reference in the constant pool index (indexbyte1 << 8 + indexbyte2)
	Getfield Opcode = 0xb4

	// 2: indexbyte1, indexbyte2
	// → value
	// get a static field value of a class, where the field is identified by field reference in the constant pool index (indexbyte1 << 8 + indexbyte2)
	Getstatic Opcode = 0xb2

	// 2: branchbyte1, branchbyte2
	// [no change]
	// goes to another instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Goto Opcode = 0xa7

	// 4: branchbyte1, branchbyte2, branchbyte3, branchbyte4
	// [no change]
	// goes to another instruction at branchoffset (signed int constructed from unsigned bytes branchbyte1 << 24 + branchbyte2 << 16 + branchbyte3 << 8 + branchbyte4)
	GotoW Opcode = 0xc8

	// value → result
	// convert an int into a byte
	I2B Opcode = 0x91

	// value → result
	// convert an int into a character
	I2C Opcode = 0x92

	// value → result
	// convert an int into a double
	I2D Opcode = 0x87

	// value → result
	// convert an int into a float
	I2F Opcode = 0x86

	// value → result
	// convert an int into a long
	I2L Opcode = 0x85

	// value → result
	// convert an int into a short
	I2S Opcode = 0x93

	// value1, value2 → result
	// add two ints
	Iadd Opcode = 0x60

	// arrayref, index → value
	// load an int from an array
	Iaload Opcode = 0x2e

	// value1, value2 → result
	// perform a bitwise AND on two integers
	Iand Opcode = 0x7e

	// arrayref, index, value →
	// store an int into an array
	Iastore Opcode = 0x4f

	// → -1
	// load the int value −1 onto the stack
	IconstM1 Opcode = 0x02

	// → 0
	// load the int value 0 onto the stack
	Iconst0 Opcode = 0x03

	// → 1
	// load the int value 1 onto the stack
	Iconst1 Opcode = 0x04

	// → 2
	// load the int value 2 onto the stack
	Iconst2 Opcode = 0x05

	// → 3
	// load the int value 3 onto the stack
	Iconst3 Opcode = 0x06

	// → 4
	// load the int value 4 onto the stack
	Iconst4 Opcode = 0x07

	// → 5
	// load the int value 5 onto the stack
	Iconst5 Opcode = 0x08

	// value1, value2 → result
	// divide two integers
	Idiv Opcode = 0x6c

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if references are equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfAcmpeq Opcode = 0xa5

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if references are not equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfAcmpne Opcode = 0xa6

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if ints are equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmpeq Opcode = 0x9f

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is greater than or equal to value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmpge Opcode = 0xa2

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is greater than value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmpgt Opcode = 0xa3

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is less than or equal to value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmple Opcode = 0xa4

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if value1 is less than value2, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmplt Opcode = 0xa1

	// 2: branchbyte1, branchbyte2
	// value1, value2 →
	// if ints are not equal, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	IfIcmpne Opcode = 0xa0

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifeq Opcode = 0x99

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is greater than or equal to 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifge Opcode = 0x9c

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is greater than 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifgt Opcode = 0x9d

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is less than or equal to 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifle Opcode = 0x9e

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is less than 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Iflt Opcode = 0x9b

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is not 0, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifne Opcode = 0x9a

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is not null, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifnonnull Opcode = 0xc7

	// 2: branchbyte1, branchbyte2
	// value →
	// if value is null, branch to instruction at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2)
	Ifnull Opcode = 0xc6

	// 2: index, const
	// [No change]
	// increment local variable #index by signed byte const
	Iinc Opcode = 0x84

	// 1: index
	// → value
	// load an int value from a local variable #index
	Iload Opcode = 0x15

	// → value
	// load an int value from local variable 0
	Iload0 Opcode = 0x1a

	// → value
	// load an int value from local variable 1
	Iload1 Opcode = 0x1b

	// → value
	// load an int value from local variable 2
	Iload2 Opcode = 0x1c

	// → value
	// load an int value from local variable 3
	Iload3 Opcode = 0x1d

	//
	// reserved for implementation-dependent operations within debuggers; should not appear in any class file
	Impdep1 Opcode = 0xfe

	//
	// reserved for implementation-dependent operations within debuggers; should not appear in any class file
	Impdep2 Opcode = 0xff

	// value1, value2 → result
	// multiply two integers
	Imul Opcode = 0x68

	// value → result
	// negate int
	Ineg Opcode = 0x74

	// 2: indexbyte1, indexbyte2
	// objectref → result
	// determines if an object objectref is of a given type, identified by class reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Instanceof Opcode = 0xc1

	// 4: indexbyte1, indexbyte2, 0, 0
	// [arg1, [arg2 ...]] → result
	// invokes a dynamic method and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Invokedynamic Opcode = 0xba

	// 4: indexbyte1, indexbyte2, count, 0
	// objectref, [arg1, arg2, ...] → result
	// invokes an interface method on object objectref and puts the result on the stack (might be void); the interface method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Invokeinterface Opcode = 0xb9

	// 2: indexbyte1, indexbyte2
	// objectref, [arg1, arg2, ...] → result
	// invoke instance method on object objectref and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Invokespecial Opcode = 0xb7

	// 2: indexbyte1, indexbyte2
	// [arg1, arg2, ...] → result
	// invoke a static method and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Invokestatic Opcode = 0xb8

	// 2: indexbyte1, indexbyte2
	// objectref, [arg1, arg2, ...] → result
	// invoke virtual method on object objectref and puts the result on the stack (might be void); the method is identified by method reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Invokevirtual Opcode = 0xb6

	// value1, value2 → result
	// bitwise int OR
	Ior Opcode = 0x80

	// value1, value2 → result
	// logical int remainder
	Irem Opcode = 0x70

	// value → [empty]
	// return an integer from a method
	Ireturn Opcode = 0xac

	// value1, value2 → result
	// int shift left
	Ishl Opcode = 0x78

	// value1, value2 → result
	// int arithmetic shift right
	Ishr Opcode = 0x7a

	// 1: index
	// value →
	// store int value into variable #index
	Istore Opcode = 0x36

	// value →
	// store int value into variable 0
	Istore0 Opcode = 0x3b

	// value →
	// store int value into variable 1
	Istore1 Opcode = 0x3c

	// value →
	// store int value into variable 2
	Istore2 Opcode = 0x3d

	// value →
	// store int value into variable 3
	Istore3 Opcode = 0x3e

	// value1, value2 → result
	// int subtract
	Isub Opcode = 0x64

	// value1, value2 → result
	// int logical shift right
	Iushr Opcode = 0x7c

	// value1, value2 → result
	// int xor
	Ixor Opcode = 0x82

	// 2: branchbyte1, branchbyte2
	// → address
	// jump to subroutine at branchoffset (signed short constructed from unsigned bytes branchbyte1 << 8 + branchbyte2) and place the return address on the stack
	Jsr Opcode = 0xa8

	// 4: branchbyte1, branchbyte2, branchbyte3, branchbyte4
	// → address
	// jump to subroutine at branchoffset (signed int constructed from unsigned bytes branchbyte1 << 24 + branchbyte2 << 16 + branchbyte3 << 8 + branchbyte4) and place the return address on the stack
	JsrW Opcode = 0xc9

	// value → result
	// convert a long to a double
	L2D Opcode = 0x8a

	// value → result
	// convert a long to a float
	L2F Opcode = 0x89

	// value → result
	// convert a long to a int
	L2I Opcode = 0x88

	// value1, value2 → result
	// add two longs
	Ladd Opcode = 0x61

	// arrayref, index → value
	// load a long from an array
	Laload Opcode = 0x2f

	// value1, value2 → result
	// bitwise AND of two longs
	Land Opcode = 0x7f

	// arrayref, index, value →
	// store a long to an array
	Lastore Opcode = 0x50

	// value1, value2 → result
	// push 0 if the two longs are the same, 1 if value1 is greater than value2, -1 otherwise
	Lcmp Opcode = 0x94

	// → 0L
	// push 0L (the number zero with type long) onto the stack
	Lconst0 Opcode = 0x09

	// → 1L
	// push 1L (the number one with type long) onto the stack
	Lconst1 Opcode = 0x0a

	// 1: index
	// → value
	// push a constant #index from a constant pool (String, int, float, Class, java.lang.invoke.MethodType, or java.lang.invoke.MethodHandle) onto the stack
	Ldc Opcode = 0x12

	// 2: indexbyte1, indexbyte2
	// → value
	// push a constant #index from a constant pool (String, int, float, Class, java.lang.invoke.MethodType, or java.lang.invoke.MethodHandle) onto the stack (wide index is constructed as indexbyte1 << 8 + indexbyte2)
	LdcW Opcode = 0x13

	// 2: indexbyte1, indexbyte2
	// → value
	// push a constant #index from a constant pool (double or long) onto the stack (wide index is constructed as indexbyte1 << 8 + indexbyte2)
	Ldc2W Opcode = 0x14

	// value1, value2 → result
	// divide two longs
	Ldiv Opcode = 0x6d

	// 1: index
	// → value
	// load a long value from a local variable #index
	Lload Opcode = 0x16

	// → value
	// load a long value from a local variable 0
	Lload0 Opcode = 0x1e

	// → value
	// load a long value from a local variable 1
	Lload1 Opcode = 0x1f

	// → value
	// load a long value from a local variable 2
	Lload2 Opcode = 0x20

	// → value
	// load a long value from a local variable 3
	Lload3 Opcode = 0x21

	// value1, value2 → result
	// multiply two longs
	Lmul Opcode = 0x69

	// value → result
	// negate a long
	Lneg Opcode = 0x75

	// 8+: <0–3 bytes padding>, defaultbyte1, defaultbyte2, defaultbyte3, defaultbyte4, npairs1, npairs2, npairs3, npairs4, match-offset pairs...
	// key →
	// a target address is looked up from a table using a key and execution continues from the instruction at that address
	Lookupswitch Opcode = 0xab

	// value1, value2 → result
	// bitwise OR of two longs
	Lor Opcode = 0x81

	// value1, value2 → result
	// remainder of division of two longs
	Lrem Opcode = 0x71

	// value → [empty]
	// return a long value
	Lreturn Opcode = 0xad

	// value1, value2 → result
	// bitwise shift left of a long value1 by int value2 positions
	Lshl Opcode = 0x79

	// value1, value2 → result
	// bitwise shift right of a long value1 by int value2 positions
	Lshr Opcode = 0x7b

	// 1: index
	// value →
	// store a long value in a local variable #index
	Lstore Opcode = 0x37

	// value →
	// store a long value in a local variable 0
	Lstore0 Opcode = 0x3f

	// value →
	// store a long value in a local variable 1
	Lstore1 Opcode = 0x40

	// value →
	// store a long value in a local variable 2
	Lstore2 Opcode = 0x41

	// value →
	// store a long value in a local variable 3
	Lstore3 Opcode = 0x42

	// value1, value2 → result
	// subtract two longs
	Lsub Opcode = 0x65

	// value1, value2 → result
	// bitwise shift right of a long value1 by int value2 positions, unsigned
	Lushr Opcode = 0x7d

	// value1, value2 → result
	// bitwise XOR of two longs
	Lxor Opcode = 0x83

	// objectref →
	// enter monitor for object (\"grab the lock\" – start of synchronized() section)
	Monitorenter Opcode = 0xc2

	// objectref →
	// exit monitor for object (\"release the lock\" – end of synchronized() section)
	Monitorexit Opcode = 0xc3

	// 3: indexbyte1, indexbyte2, dimensions
	// count1, [count2,...] → arrayref
	// create a new array of dimensions dimensions of type identified by class reference in constant pool index (indexbyte1 << 8 + indexbyte2); the sizes of each dimension is identified by count1, [count2, etc.]
	Multianewarray Opcode = 0xc5

	// 2: indexbyte1, indexbyte2
	// → objectref
	// create new object of type identified by class reference in constant pool index (indexbyte1 << 8 + indexbyte2)
	New Opcode = 0xbb

	// 1: atype
	// count → arrayref
	// create new array with count elements of primitive type identified by atype
	Newarray Opcode = 0xbc

	// [No change]
	// perform no operation
	Nop Opcode = 0x00

	// value →
	// discard the top value on the stack
	Pop Opcode = 0x57

	// {value2, value1} →
	// discard the top two values on the stack (or one value, if it is a double or long)
	Pop2 Opcode = 0x58

	// 2: indexbyte1, indexbyte2
	// objectref, value →
	// set field to value in an object objectref, where the field is identified by a field reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Putfield Opcode = 0xb5

	// 2: indexbyte1, indexbyte2
	// value →
	// set static field to value in a class, where the field is identified by a field reference index in constant pool (indexbyte1 << 8 + indexbyte2)
	Putstatic Opcode = 0xb3

	// 1: index
	// [No change]
	// continue execution from address taken from a local variable #index (the asymmetry with jsr is intentional)
	Ret Opcode = 0xa9

	// → [empty]
	// return void from method
	Return Opcode = 0xb1

	// arrayref, index → value
	// load short from array
	Saload Opcode = 0x35

	// arrayref, index, value →
	// store short to array
	Sastore Opcode = 0x56

	// 2: byte1, byte2
	// → value
	// push a short onto the stack as an integer value
	Sipush Opcode = 0x11

	// value2, value1 → value1, value2
	// swaps two top words on the stack (note that value1 and value2 must not be double or long)
	Swap Opcode = 0x5f

	// 16+: [0–3 bytes padding], defaultbyte1, defaultbyte2, defaultbyte3, defaultbyte4, lowbyte1, lowbyte2, lowbyte3, lowbyte4, highbyte1, highbyte2, highbyte3, highbyte4, jump offsets...
	// index →
	// continue execution from an address in the table at offset index
	Tableswitch Opcode = 0xaa

	// 3/5: opcode, indexbyte1, indexbyte2or iinc, indexbyte1, indexbyte2, countbyte1, countbyte2
	// [same as for corresponding instructions]
	// execute opcode, where opcode is either iload, fload, aload, lload, dload, istore, fstore, astore, lstore, dstore, or ret, but assume the index is 16 bit; or execute iinc, where the index is 16 bits and the constant to increment by is a signed 16 bit short
	Wide Opcode = 0xc4
)
