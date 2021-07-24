package models

import "time"

type TODO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Timer Timer  `json:"timer"`
}

type Timer struct {
	IsSet     bool      `json:"is_set"`
	IsTimeOut bool      `json:"is_timeout"`
	Time      time.Time `json:"time"`
}
