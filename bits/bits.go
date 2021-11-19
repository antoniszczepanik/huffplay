package main

import (
	"fmt"
	"io"
)

type BitSet struct {
	bits []bool
}

type byteArray [8]bool

func NewBitSet(bits []bool) *BitSet {
	return &BitSet{
		bits: bits,
	}
}

func (bs *BitSet) Read(p []byte) (int, error) {
	chunks := getChunks(bs.bits)
	if len(p) < len(chunks) {
		return 0, fmt.Errorf("TODO: buffer to small (len=%d) to read BitSet (len=%d)", len(p), len(chunks))
	}
	var i int
	for i = 0; i < len(chunks); i += 1 {
		p[i] = bitsToByte(chunks[i])
	}
	return i, io.EOF
}

// Split BitSet into byte length arrays.
func getChunks(bits []bool) []byteArray {
	const chunkSize = 8
	var result []byteArray
	var byteBits byteArray
	for start := 0; start < len(bits); start += chunkSize {
		end := start + chunkSize
		if end > len(bits) {
			// Less than 8 bits.
			end = len(bits)
			// Make sure it's zeroed.
			byteBits = byteArray{false}
			for i := 7; i > 7-(end-start); i -= 1 {
				byteBits[i] = bits[7-i]
			}
			result = append(result, byteBits)
		} else {
			// Full 8 bits.
			copied := copy(byteBits[:], bits[start:end])
			if copied != 8 {
				panic("ERROR: chunk should copy exactly 8 bits")
			}
			result = append(result, byteBits)
		}
	}
	return result
}

func (bs *BitSet) ReadBits() []bool {
	return bs.bits
}

func (bs *BitSet) AppendBits(bits []bool) error {
	bs.bits = append(bs.bits, bits...)
	return nil
}

// Convert byte lenght array of bits into byte.
func bitsToByte(bits byteArray) byte {
	var result byte
	for i, bit := range bits {
		result |= b2i(bit)
		// Do not shift the last bit.
		if i != 7 {
			result = result << 1
		}
	}
	return result
}

func b2i(b bool) byte {
	if b {
		return 1
	}
	return 0
}
