package calendar

import (
	"errors"
	"fmt"
)

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(day int) error {
	if day > 0 && day < 32 {
		d.day = day
		return nil
	} else {
		return errors.New("Value of day invalid! It must be 0 < day < 32!")
	}
}

func (d *Date) SetMonth(month int) error {
	if month > 0 && month < 13 {
		d.month = month
		return nil
	} else {
		return errors.New("Value of month invalid! It must be 0 < month < 13!")
	}
}

func (d *Date) SetYear(year int) error {
	if year > 0 && year < 10000 {
		d.year = year
		return nil
	} else {
		return errors.New("Value of year invalid! It must be 0 < year < 10000!")
	}
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) GetInfo() string {
	return fmt.Sprintf("Date of event: %d.%d.%d\n", d.day, d.month, d.year)
}
