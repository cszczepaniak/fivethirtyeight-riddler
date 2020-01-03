package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	"image/png"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL   = `https://www.cia.gov/library/publications/the-world-factbook/`
	homePage  = baseURL + `docs/flagsoftheworld.html`
	flagsBase = `attachments/flags/`

	flagsDir = `../../flags/`

	flagFileRegex = `[A-Z]{2}-flag\.gif$`
)

func main() {
	err := downloadImages()
	if err != nil {
		log.Fatal(err)
	}
}

func downloadImages() error {
	err := downloadMysteryFlags()
	if err != nil {
		return err
	}
	err = downloadWorldFlagImages()
	if err != nil {
		return err
	}
	return nil
}

func downloadMysteryFlags() error {
	fmt.Println(`downloading mystery flag images...`)
	const urlBase = `https://fivethirtyeight.com/wp-content/uploads/2020/01/`
	for i := 0; i < 3; i++ {
		filename := fmt.Sprintf(`flag_%d.png`, i+1)
		im, err := readImg(urlBase, filename)
		if err != nil {
			return err
		}
		err = saveImg(im, `../../inputData/`+filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadWorldFlagImages() error {
	fmt.Println(`downloading world flag images...`)
	resp, err := http.Get(homePage)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	imgs := doc.Find(`div.modalFlagBox`).Children()
	regex := regexp.MustCompile(flagFileRegex)
	n := imgs.Length()
	for i := 0; i < n; i++ {
		src, ok := imgs.Eq(i).Attr(`src`)
		if !ok {
			return errors.New(`src attribute not found`)
		}
		filename := regex.FindString(src)
		if filename == `` {
			continue
		}
		im, err := readImg(baseURL+flagsBase, filename)
		if err != nil {
			return err
		}
		trimmed, err := trimBorder(im, 1)
		if err != nil {
			return err
		}
		b := trimmed.Bounds()
		fmt.Printf("image has size %dx%d\n", b.Dx(), b.Dy())
		err = saveImg(trimmed, flagsDir+filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func trimBorder(im *image.RGBA, px int) (*image.RGBA, error) {
	b := im.Bounds()
	r := image.Rect(0, 0, b.Max.X, b.Max.Y).Inset(px)
	sub := im.SubImage(r)
	res, ok := sub.(*image.RGBA)
	if !ok {
		return nil, errors.New(`type assertion failed`)
	}
	return res, nil
}

func readImg(baseURL, filename string) (*image.RGBA, error) {
	time.Sleep(time.Second)
	resp, err := http.Get(baseURL + filename)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	im, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	rgb, ok := im.(*image.RGBA)
	if ok {
		return rgb, nil
	}
	// assertion failed. we need to create a new image and draw into it
	b := im.Bounds()
	rgb = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgb, b, im, b.Min, draw.Src)
	return rgb, nil
}

func saveImg(im *image.RGBA, target string) error {
	fmt.Printf("saving image to %s...\n", target)
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}
