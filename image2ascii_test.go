package main

import (
	"image"
	"image/color"
	"os"
	"testing"

	"golang.org/x/image/font/basicfont"
)

var charset = "^_`abcdefghijklmnopqrstuvwxyz~*+-.:<=>{}0123456789?@ABCDEFGHIJKLMNOPQRSTUVWXYZ#$%&"

func TestCharsetMapping(t *testing.T) {
	tests := []struct {
		brightness float64
		expected   string
	}{
		{0.0, "^"},
		{1.0, "&"},
		{0.5, "0"},
	}

	for _, tt := range tests {
		index := int(tt.brightness * float64(len(charset)-1))
		if index < 0 || index >= len(charset) {
			t.Fatalf("Index out of bounds: %d for brightness %f", index, tt.brightness)
		}
		ch := string(charset[index])
		if ch != tt.expected {
			t.Errorf("For brightness %f, expected %s, got %s", tt.brightness, tt.expected, ch)
		}
	}
}

func TestResizeImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	char_w := 7
	char_h := 14

	expectedWidth := img.Bounds().Dx() / char_w
	expectedHeight := img.Bounds().Dy() / char_h

	resizedImg := resizeImage(img, expectedWidth, expectedHeight)

	if resizedImg.Bounds().Dx() != expectedWidth || resizedImg.Bounds().Dy() != expectedHeight {
		t.Fatalf("Expected resized dimensions %dx%d, got %dx%d",
			expectedWidth, expectedHeight,
			resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy())
	}
}

func TestFileHandling(t *testing.T) {
	testFiles := []string{
		"testdata/test_image.png",
		"testdata/test_image.jpg",
	}

	for _, file := range testFiles {
		_, err := os.Open(file)
		if err != nil {
			t.Errorf("Failed to open file: %s", file)
		}
	}
}

func TestImageConversion(t *testing.T) {
	original := image.NewRGBA(image.Rect(0, 0, 100, 100))

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			original.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}

	outputImage := imageToAscii(original, charset, 7, 13, basicfont.Face7x13)

	if outputImage == nil {
		t.Fatal("Expected output image, got nil")
	}
}
