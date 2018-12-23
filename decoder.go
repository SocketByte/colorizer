package main

import (
	"image"
	"io"
	"log"
)

func GetPixels(file io.Reader) ([][]Pixel, int, int) {
	decoded, _, err := image.Decode(file)

	if err != nil {
		log.Println(err)
		return nil, 0, 0
	}

	bounds := decoded.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, convert(decoded.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, width, height
}

func convert(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
