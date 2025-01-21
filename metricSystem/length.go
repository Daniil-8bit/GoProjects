// Length types with methods
package metricsystem

// Kilometr
type Kilometr float64

func (k Kilometr) ToMetr() Metr {
	return Metr(k) * 1000
}

func (k Kilometr) ToСentimeter() Сentimeter {
	return Сentimeter(k) * 100000
}

func (k Kilometr) ToMillimetr() Millimetr {
	return Millimetr(k) * 10000000
}

//Metr
type Metr float64

func (m Metr) ToKilometr() Kilometr {
	return Kilometr(m) / 1000
}

func (m Metr) ToСentimeter() Сentimeter {
	return Сentimeter(m) * 100
}

func (m Metr) ToMillimetr() Millimetr {
	return Millimetr(m) * 10000
}

// Сentimeter
type Сentimeter float64

func (с Сentimeter) ToKilometr() Kilometr {
	return Kilometr(с) / 100000
}

func (с Сentimeter) ToMetr() Metr {
	return Metr(с) / 100
}

func (с Сentimeter) ToMillimetr() Millimetr {
	return Millimetr(с) * 100
}

// Millimetr
type Millimetr float64

func (m Millimetr) ToKilometr() Kilometr {
	return Kilometr(m) / 1000000
}

func (m Millimetr) ToMetr() Metr {
	return Metr(m) / 1000
}

func (m Millimetr) ToСentimeter() Сentimeter {
	return Сentimeter(m) / 10
}
