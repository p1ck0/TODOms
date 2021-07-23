package models

import "time"

type TODO struct {
	ID    string
	Name  string
	Timer time.Time
}
