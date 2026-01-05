package models

type GroupedLight struct {
	ID               string             `json:"id"`
	IDV1             string             `json:"id_v1,omitempty"`
	Owner            ResourceIdentifier `json:"owner"`
	Type             string             `json:"type"` // "grouped_light"
	On               *On                `json:"on,omitempty"`
	Dimming          *Dimming           `json:"dimming,omitempty"`
	ColorTemperature *ColorTemperature  `json:"color_temperature,omitempty"`
	Color            *Color             `json:"color,omitempty"`
	Alert            *Alert             `json:"alert,omitempty"`
	Signaling        *Signaling         `json:"signaling,omitempty"`
	Dynamics         *Dynamics          `json:"dynamics,omitempty"`
}

type GroupedLightUpdate struct {
	Type                  string                 `json:"type,omitempty"` // "grouped_light"
	On                    *On                    `json:"on,omitempty"`
	Dimming               *Dimming               `json:"dimming,omitempty"`
	DimmingDelta          *DimmingDelta          `json:"dimming_delta,omitempty"`
	ColorTemperature      *ColorTemperature      `json:"color_temperature,omitempty"`
	ColorTemperatureDelta *ColorTemperatureDelta `json:"color_temperature_delta,omitempty"`
	Color                 *Color                 `json:"color,omitempty"`
	Alert                 *Alert                 `json:"alert,omitempty"`     // "breathe"
	Signaling             *Signaling             `json:"signaling,omitempty"` //duration must be greater than 999
	Dynamics              *Dynamics              `json:"dynamics,omitempty"`
}
