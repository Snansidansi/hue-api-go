package models

type LightChangeEvent struct {
	ID      string
	On      *On
	Dimming *Dimming
	Color   *Color
}

type GroupChangeEvent struct {
	ID      string
	On      *On
	Dimming *Dimming
	Type    string // "grouped_light", "zone", "room"
}

type ButtonEvent struct {
	ID        string
	EventType string
}
