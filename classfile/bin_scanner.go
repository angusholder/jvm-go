package classfile

import "encoding/binary"

type BinScanner struct {
	slice []uint8
}

type BinScannerOutOfBounds struct{}

func (b *BinScanner) checkBounds(n int) {
	if len(b.slice) < n {
		panic(BinScannerOutOfBounds{})
	}
}

func (b *BinScanner) Uint8() uint8 {
	b.checkBounds(1)
	n := b.slice[0]
	b.slice = b.slice[1:]
	return n
}

func (b *BinScanner) Uint16() uint16 {
	b.checkBounds(2)
	n := binary.BigEndian.Uint16(b.slice)
	b.slice = b.slice[2:]
	return n
}

func (b *BinScanner) Uint32() uint32 {
	b.checkBounds(4)
	n := binary.BigEndian.Uint32(b.slice)
	b.slice = b.slice[4:]
	return n
}

func (b *BinScanner) Uint64() uint64 {
	b.checkBounds(8)
	n := binary.BigEndian.Uint64(b.slice)
	b.slice = b.slice[8:]
	return n
}

func (b *BinScanner) Bytes(count int) []uint8 {
	b.checkBounds(count)
	bytes := b.slice[0:count]
	b.slice = b.slice[count:]
	return bytes
}
