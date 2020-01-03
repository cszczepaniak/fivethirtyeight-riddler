package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/png"
	"os"
	"strconv"
)

const binSize = 8

type bin int

func (b bin) String() string {
	return strconv.Itoa(int(b))
}

func newBin(n uint32) bin {
	val := n / binSize * binSize
	return bin(val)
}

type colorData struct {
	R uint32
	G uint32
	B uint32
}

func (cd colorData) String() string {
	return fmt.Sprintf(`{R:%d G:%d B:%d}`, cd.R, cd.G, cd.B)
}

func newColorData(c color.Color) colorData {
	r, g, b, _ := c.RGBA()
	return colorData{
		R: r,
		G: g,
		B: b,
	}
}

func main() {
	reader, err := os.Open(`inputData/flag_1.png`)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	im, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	bnds := im.Bounds().Size()

	clrMap := make(map[colorData]int)
	for i := 0; i < bnds.X; i++ {
		for j := 0; j < bnds.Y; j++ {
			c := im.At(i, j)
			cd := newColorData(c)
			_, ok := clrMap[cd]
			if !ok {
				clrMap[cd] = 1
			} else {
				clrMap[cd]++
			}
		}
	}
}

// func binEdges() []int {
// 	e := make([]int, 16)
// 	for i := range e {
// 		e[i] = i*16
// 	}
// 	return e
// }

// func makeBin(n uint32, bins []int) (bin, error) {
// 	if n < 0 || n > 255 {
// 		return [2]int{0, 0}, fmt.Errorf(`value of n [%d] is out of range`, n)
// 	}
// 	signed := int(n)
// 	low, hi := signed-10, signed+10
// 	if low < 0 {
// 		low = 0
// 	}
// 	if high
// }
