package bits

import (
	"testing"
)

func TestBitSlice_ReadBits(t *testing.T) {
	expectedBits := []bool{false, true}
	bs := NewBitSet(expectedBits)
	actualBits := bs.ReadBits()
	compareBoolSlices(t, expectedBits, actualBits)
}

func TestBitSlice_Append(t *testing.T) {
	someBits := []bool{false, true}
	bs := NewBitSet(someBits)
	bs.AppendBits(someBits)
	actualBits := bs.ReadBits()
	compareBoolSlices(t, append(someBits, someBits...), actualBits)
}

func TestBitSlice_ReadAll(t *testing.T) {
	inputBits := []bool{true, true, true, true, true, true, true, true, true}
	bs := NewBitSet(inputBits)
	actualBytes, err := bs.ReadAll()
	if err != nil {
		t.Errorf("ReadAll: %w", err)
	}
	compareByteSlices(t, []byte{255 , 1}, actualBytes)
}

func TestBitSlice_Read(t *testing.T) {
	tests := []struct {
		name                string
		inputBits           []bool
		expectedOutputBytes []byte
		expectedErr         error
	}{
		{
			name:                "Single bit on",
			inputBits:           []bool{true},
			expectedOutputBytes: []byte{1},
			expectedErr:         nil,
		},
		{
			name:                "Exactly 8 bits on",
			inputBits:           []bool{true, true, true, true, true, true, true, true},
			expectedOutputBytes: []byte{255},
			expectedErr:         nil,
		},
		{
			name:                "Exactly 9 bits on",
			inputBits:           []bool{true, true, true, true, true, true, true, true, true},
			expectedOutputBytes: []byte{255, 1},
			expectedErr:         nil,
		},
		{
			name:                "Exactly 9 bits off",
			inputBits:           []bool{false, false, false, false, false, false, false, false, false},
			expectedOutputBytes: []byte{0, 0},
			expectedErr:         nil,
		},
		{
			name:                "Empty input",
			inputBits:           []bool{},
			expectedOutputBytes: []byte{},
			expectedErr:         nil,
		},
	}
	for _, test := range tests {
		bs := NewBitSet(test.inputBits)
		expectedOutputLenght := getByteCount(len(test.inputBits))
		outputBuffer := make([]byte, expectedOutputLenght)
		p, err := bs.Read(outputBuffer)
		if test.expectedErr != nil {
			if test.expectedErr != err {
				t.Fatalf("%s: unexpected error", test.name)
			}
		}
		if p != expectedOutputLenght {
			t.Fatalf("%s: unexpected output lenght %d != %d", test.name, p, expectedOutputLenght)
		}

		compareByteSlices(t, outputBuffer, test.expectedOutputBytes)
	}
}

func compareBoolSlices(t *testing.T, s1 []bool, s2 []bool) {
	if len(s1) != len(s2) {
		t.Errorf("slice lenghts not equal (%d != %d)", len(s1), len(s2))
	}
	for i := 0; i < len(s1); i += 1 {
		if s1[i] != s2[i] {
			t.Errorf("slice elements at %d differ: %v != %v", i, s1[i], s2[i])
		}
	}
}

// This is ridicullus.
func compareByteSlices(t *testing.T, s1 []byte, s2 []byte) {
	if len(s1) != len(s2) {
		t.Errorf("slice lenghts not equal (%d != %d)", len(s1), len(s2))
	}
	for i := 0; i < len(s1); i += 1 {
		if s1[i] != s2[i] {
			t.Errorf("slice elements at %d differ: %v != %v", i, s1[i], s2[i])
		}
	}
}

func getExpectedOuputLength(inputBits []bool) int {
	if len(inputBits) == 0 {
		return 0
	}
	return (len(inputBits) - 1) / 8 + 1
}
