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
	im1Ct, im2Ct := 0, 0
	for _, ct := range im1 {
		im1Ct += ct
	}
	for _, ct := range im2 {
		im2Ct += ct
	}
	if im1Ct != im2Ct {
		return 0, errors.New(`pixel count between the images are not the same`)
	}
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

	targetW, targetH := getImageDimensions(im)
	fmt.Println(targetW)
	fmt.Println(targetH)
	for _, f := range files {
		thisFlag, err := decodeImage(`flags/` + f.Name())
		if err != nil {
			panic(err)
		}
		w, h := getImageDimensions(thisFlag)
		if w != targetW || h != targetH {
			continue
		}
		fmt.Println(`found image with proper dimensions: ` + f.Name())
	}
}
