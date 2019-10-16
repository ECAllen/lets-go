package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Memory struct {
	ID int
	Title string
	Content string
	Created string
}
