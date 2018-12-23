package main

func CalculatePercentage(pairs []*PixelPair) []*PixelPair {
	amount := 0
	for _, pair := range pairs {
		amount += pair.Amount
	}

	for _, pair := range pairs {
		pair.Percentage = int((float64(pair.Amount) / float64(amount)) * 100.0)
	}
	return pairs
}
