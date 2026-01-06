package models

import "time"

type Button struct {
	ID           string             `json:"id"`
	IDV1         string             `json:"id_v1,omitempty"`
	Owner        ResourceIdentifier `json:"owner"`
	Type         string             `json:"type"` // "button"
	Metadata     ButtonMetadata     `json:"metadata"`
	Button       ButtonSpecific     `json:"button"`
	ButtonReport *ButtonReport      `json:"button_report,omitempty"`
}

type ButtonMetadata struct {
	ControlID int `json:"control_id"`
}

type ButtonSpecific struct {
	LastEvent string `json:"last_event,omitempty"` // Deprecated
}

type ButtonReport struct {
	Updated        time.Time `json:"updated"`
	Event          string    `json:"event"`
	RepeatInterval *int      `json:"repeat_interval,omitempty"`
	EventValues    []string  `json:"event_values,omitempty"`
}
