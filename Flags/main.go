package main

import (
	"errors"
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
	return colorData{
		R: r,
		G: g,
		B: b,
	}
}

func decodeImage(filename string) (*image.RGBA, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	im, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	rgba, ok := im.(*image.RGBA)
	if !ok {
		return nil, errors.New(`conversion to RGBA failed on ` + filename)
	}
	return rgba, nil
}

func getImageDimensions(im *image.RGBA) (int, int) {
	b := im.Bounds()
	return b.Dx(), b.Dy()
}

func getPixels(im *image.RGBA) (map[colorData]int, error) {
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

func compareImages(im1, im2 map[colorData]int) (int, error) {
	im1Ct, im2Ct, inCommon := 0, 0, 0
	for clr, ct1 := range im1 {
		ct2, ok := im2[clr]
		if ok {
			if ct1 < ct2 {
				inCommon += ct1
			} else {
				inCommon += ct2
			}
		}

		im1Ct += ct1
	}
	for _, ct := range im2 {
		im2Ct += ct
	}
	if im1Ct != im2Ct {
		return 0, errors.New(`pixel count between the images are not the same`)
	}
	return inCommon, nil
}

func main() {
	files, err := getFilesInFolder(`flags/`)
	if err != nil {
		panic(err)
	}

	im, err := decodeImage(`inputData/flag_1.png`)
	if err != nil {
		panic(err)
	}
	mysteryPx, err := getPixels(im)
	if err != nil {
		panic(err)
	}

	targetW, targetH := getImageDimensions(im)
	bestRank, bestImage := 0, ``
	for _, f := range files {
		thisFlag, err := decodeImage(`flags/` + f.Name())
		if err != nil {
			fmt.Printf("an error occurred decoding %s. skipping...\n", f.Name())
			continue
		}
		w, h := getImageDimensions(thisFlag)
		if w != targetW || h != targetH {
			continue
		}
		thesePx, err := getPixels(thisFlag)
		if err != nil {
			panic(err)
		}
		rank, err := compareImages(mysteryPx, thesePx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("pixels in common with %s: %d\n", f.Name(), rank)
		if rank > bestRank {
			fmt.Printf("this rank of %d beats the best rank of %d\n", rank, bestRank)
			bestRank = rank
			bestImage = f.Name()
		}
	}
	fmt.Printf("the flag with the most pixels in common is %s\n", bestImage)
}
