package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"sort"
)

type colorData struct {
	R     uint32
	G     uint32
	B     uint32
	count int
}

func main() {
	reader, err := os.Open(`flag_1.png`)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	im, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	bnds := im.Bounds().Size()

	clrMap := make(map[color.Color]colorData)
	for i := 0; i < bnds.X; i++ {
		for j := 0; j < bnds.Y; j++ {
			c := im.At(i, j)
			_, ok := clrMap[c]
			if !ok {
				r, g, b, _ := c.RGBA()
				clrMap[c] = colorData{
					R:     r / 256,
					G:     g / 256,
					B:     b / 256,
					count: 1,
				}
			} else {
				dat := clrMap[c]
				dat.count++
				clrMap[c] = dat
			}
		}
	}
	totalPix := float64(bnds.X * bnds.Y)
	all := make([]colorData, len(clrMap))
	idx := 0
	for _, c := range clrMap {
		all[idx] = c
		idx++
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].count > all[j].count
	})
	for _, c := range all {
		pct := float64(c.count) / totalPix * 100.0
		if pct < 0.1 {
			break
		}
		fmt.Printf("Color [%d, %d, %d] makes up %f percent of the map\n", c.R, c.G, c.B, pct)
	}
}
