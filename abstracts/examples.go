package abstracts

import "fmt"

type Example string

func (e Example) Error() string {
	return string(e)
}

type HighTemp float64

func (h HighTemp) Error() string {
	return fmt.Sprintf("Temperature is %.2f degrees\n", h)
}

func checkAvailableTemp(currentValue float64, maxValue float64) error {
	result := currentValue - maxValue
	if result > 0 {
		return HighTemp(result)
	}
	return nil
}

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %s %s, year: %d\n", c.Brand, c.Model, c.Year)
}
