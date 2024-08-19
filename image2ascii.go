package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <filename>")
	}

	filename := os.Args[1]
	charset := "^_`abcdefghijklmnopqrstuvwxyz~*+-.:<=>{}0123456789?@ABCDEFGHIJKLMNOPQRSTUVWXYZ#$%&"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	fontFace := basicfont.Face7x13
	charW := 7
	charH := 13

	newImg := imageToAscii(img, charset, charW, charH, fontFace)

	ext := strings.ToLower(filepath.Ext(filename))
	outputFilename := fmt.Sprintf("%s_processed%s", strings.TrimSuffix(filename, ext), ext)

	outFile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outFile.Close()

	switch ext {
	case ".png":
		err = png.Encode(outFile, newImg)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outFile, newImg, nil)
	case ".gif":
		err = gif.Encode(outFile, newImg, nil)
	default:
		log.Fatalf("Unsupported output format: %v", ext)
	}

	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	fmt.Println("Completed!")
}

func resizeImage(img image.Image, width, height int) (dst *image.RGBA) {
	dst = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return
}

func imageToAscii(img image.Image, charset string, charW, charH int, fontFace font.Face) *image.RGBA {
	newWidth := img.Bounds().Dx() / charW
	newHeight := img.Bounds().Dy() / charH
	resizedImg := resizeImage(img, newWidth, newHeight)

	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), image.Black, image.Point{}, draw.Src)

	drawer := &font.Drawer{
		Dst:  newImg,
		Src:  image.NewUniform(color.White),
		Face: fontFace,
	}

	for y := 0; y < resizedImg.Bounds().Dy(); y++ {
		for x := 0; x < resizedImg.Bounds().Dx(); x++ {
			r, g, b, _ := resizedImg.At(x, y).RGBA()

			r = r >> 8
			g = g >> 8
			b = b >> 8

			brightness := (float64(r) + float64(g) + float64(b)) / 3.0

			n := brightness / 255.0

			index := int(n * float64(len(charset)-1))

			if index < 0 {
				index = 0
			} else if index >= len(charset) {
				index = len(charset) - 1
			}

			ch := string(charset[index])

			drawer.Src = image.NewUniform(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
			drawer.Dot = fixed.Point26_6{
				X: fixed.I(x * charW),
				Y: fixed.I((y + 1) * charH),
			}
			drawer.DrawString(ch)
		}
	}

	return newImg
}
