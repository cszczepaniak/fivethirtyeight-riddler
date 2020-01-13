package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/png"
	"os"
)

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
	q := 10
	return colorData{
		// bin the colors to account for image conversion artifacts
		R: r >> q,
		G: g >> q,
		B: b >> q,
	}
}

type pixelDistribution map[colorData]int

type imageData struct {
	file string
	px   pixelDistribution
}

func decodeImageFromFile(filename string) (image.Image, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	im, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return im, nil
}

func getPixelDistribution(im image.Image) (pixelDistribution, error) {
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
	return clrMap, nil
}

func getFilesInFolder(folder string) ([]os.FileInfo, error) {
	f, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func compareImages(im1, im2 pixelDistribution) int {
	inCommon := 0
	for clr, ct1 := range im1 {
		ct2, ok := im2[clr]
		if ok {
			if ct1 < ct2 {
				inCommon += ct1
			} else {
				inCommon += ct2
			}
		}
	}
	return inCommon
}

func getPixelsForFlags() ([]imageData, error) {
	files, err := getFilesInFolder(`flags/`)
	if err != nil {
		return nil, err
	}
	result := make([]imageData, len(files))
	for i, f := range files {
		im, err := decodeImageFromFile(`flags/` + f.Name())
		if err != nil {
			return nil, err
		}
		px, err := getPixelDistribution(im)
		if err != nil {
			return nil, err
		}
		result[i] = imageData{
			file: f.Name(),
			px:   px,
		}
	}
	return result, nil
}

func findClosestMatch(flags []imageData, mystery imageData) (string, int, error) {
	bestRank, bestImage := 0, ``
	for _, f := range flags {
		rank := compareImages(f.px, mystery.px)
		if rank > bestRank {
			bestRank = rank
			bestImage = f.file
		}
	}
	return bestImage, bestRank, nil
}

func main() {
	inputs, err := getFilesInFolder(`inputData/`)
	if err != nil {
		panic(err)
	}
	flagData, err := getPixelsForFlags()
	if err != nil {
		panic(err)
	}
	for _, in := range inputs {
		im, err := decodeImageFromFile(`inputData/` + in.Name())
		if err != nil {
			panic(err)
		}
		px, err := getPixelDistribution(im)
		if err != nil {
			panic(err)
		}
		inputData := imageData{
			file: in.Name(),
			px:   px,
		}

		m, n, err := findClosestMatch(flagData, inputData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("the closest match to %s is %s with a score of %d\n", in.Name(), m, n)
	}
}
