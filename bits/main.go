package main

import (
	"fmt"
	"github.com/antoniszczepanik/lzhcomp/bits"
)

func main() {
	bs := bits.NewBitSet([]bool{false, true, false})
	fmt.Println(bs.ReadBits())
}
