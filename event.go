package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() {
	//add em um db
	events = append(events, e)
}
