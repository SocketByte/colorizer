package main

import "sort"

type PixelPair struct {
	Pixel      Pixel
	Amount     int
	Percentage int
}

func SortPixels(pixels map[Pixel]int, include int) *[]*PixelPair {
	var pairs []*PixelPair
	for k, v := range pixels {
		pairs = append(pairs, &PixelPair{k, v, 0})
	}

	sort.Slice(pairs, func(x, y int) bool {
		return pairs[x].Amount > pairs[y].Amount
	})

	includePairs := make([]*PixelPair, include)
	index := 0
	for x := range pairs {
		includePairs[x] = pairs[x]

		index++
		if index >= include {
			break
		}
	}

	includePairs = CalculatePercentage(includePairs)

	return &includePairs
}
