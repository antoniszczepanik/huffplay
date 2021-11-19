package bits

import (
	"fmt"
	"io"
	"errors"
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

func (bs *BitSet) ReadAll() ([]byte, error) {
	byteBuffer := make([]byte, getByteCount(len(bs.bits)))
	n, err := bs.Read(byteBuffer)
	if n != len(byteBuffer) {
		return []byte{}, fmt.Errorf("could not read all bytes: %d out of %d", n, len(byteBuffer))
	}
	if !errors.Is(err, io.EOF) {
		return []byte{}, fmt.Errorf("read BitSet: %w", err)
	}
	return byteBuffer, nil
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

func (bs *BitSet) ReadBits() []bool {
	return bs.bits
}

func (bs *BitSet) AppendBits(bits []bool) error {
	bs.bits = append(bs.bits, bits...)
	return nil
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

func getByteCount(bitCount int) int {
	if bitCount == 0 {
		return 0
	}
	return (bitCount - 1) / 8 + 1
}

func b2i(b bool) byte {
	if b {
		return 1
	}
	return 0
}
