package models

import "time"

const (
	EventTypeAdd    = "add"
	EventTypeUpdate = "update"
	EventTypeDelete = "delete"
)

const (
	GroupTypeRoom       = "room"
	GroupTypeZone       = "zone"
	GroupTypeLightGroup = "grouped_light"
)

type BaseEventFields struct {
	EventType    string // add, update, delete
	ID           string
	Timestamp    time.Time // from creationtime
	StateChanges bool      // true ONLY if it is an 'update' AND state data is present AND NO config data is present
}

type LightChangeEvent struct {
	BaseEventFields
	On               *On
	Dimming          *Dimming
	Color            *Color
	ColorTemperature *ColorTemperature
	Dynamics         *Dynamics
}

type GroupChangeEvent struct {
	BaseEventFields
	Type    string   // room, zone, grouped_light
	On      *On      // State
	Dimming *Dimming // State
}

type ButtonEvent struct {
	BaseEventFields
	Button string // initial_press, repeat, ...
}

type SceneEvent struct {
	BaseEventFields
	Status *SceneStatusEvent // State
}

type SceneStatusEvent struct {
	Active     string     `json:"active,omitempty"`
	LastRecall *time.Time `json:"last_recall,omitempty"`
}

type EntertainmentConfigurationEvent struct {
	BaseEventFields
	Status         string              // State (active/inactive)
	ActiveStreamer *ResourceIdentifier // State
}
