package main

import (
	"gopkg.in/go-playground/colors.v1"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func (pixel Pixel) ToHex() string {
	rgba, _ := colors.RGBA(uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), float64(pixel.A/255.0))
	return rgba.ToHEX().String()
}

func (pixel Pixel) ToRGBString() string {
	rgba, _ := colors.RGBA(uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), float64(pixel.A/255.0))
	return rgba.ToRGB().String()
}
