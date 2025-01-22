package calendar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(s string) error {
	if utf8.RuneCountInString(s) > 50 {
		return errors.New("You title is invalid! Too big value!")
	}

	e.title = s
	return nil
}

func (e *Event) GetEventInfo() string {
	return fmt.Sprintf("Event: %s, date: %d.%d.%d\n", e.Title(), e.Day(), e.Month(), e.Year())
}
