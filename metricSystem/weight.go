package metricsystem

type Kilogram float64

func (k Kilogram) ToGram() Gram {
	return Gram(k * 1000)
}

func (k Kilogram) ToMilligram() Milligram {
	return Milligram(k * 1000000)
}

type Gram float64

func (g Gram) ToKilogram() Kilogram {
	return Kilogram(g / 1000)
}

func (g Gram) ToMilligram() Milligram {
	return Milligram(g * 1000)
}

type Milligram float64

func (m Milligram) ToGram() Gram {
	return Gram(m / 1000)
}

func (m Milligram) ToKilogram() Kilogram {
	return Kilogram(m / 1000000)
}
