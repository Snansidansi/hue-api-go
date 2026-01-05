package models

type DevicePower struct {
	ID         string             `json:"id"`
	IDV1       string             `json:"id_v1,omitempty"`
	Owner      ResourceIdentifier `json:"owner"`
	Type       string             `json:"type"` // "device_power"
	PowerState PowerState         `json:"power_state"`
}

type PowerState struct {
	BatteryState *string `json:"battery_state,omitempty"`
	BatteryLevel *int    `json:"battery_level,omitempty"`
}
