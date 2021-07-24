package expiration

import "time"

type TimeExp struct {
	ID    string
	Timer *time.Timer
}

func SetTimer(id string, t time.Time) {
	exp := t.Sub(time.Now())
	timeExp := TimeExp{
		ID:    id,
		Timer: time.NewTimer(exp),
	}
	<-timeExp.Timer.C
}
