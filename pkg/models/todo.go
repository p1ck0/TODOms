package models

import "time"

type TODO struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Timer time.Time `json:"time"`
}
