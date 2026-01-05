package models

import "time"

type SceneMetadata struct {
	Name    string              `json:"name,omitempty"`
	Image   *ResourceIdentifier `json:"image,omitempty"`
	Appdata string              `json:"appdata,omitempty"`
}

type ActionPayload struct {
	On               *On               `json:"on,omitempty"`
	Dimming          *Dimming          `json:"dimming,omitempty"`
	Color            *Color            `json:"color,omitempty"`
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Gradient         *Gradient         `json:"gradient,omitempty"`
	Effects          *Effects          `json:"effects,omitempty"` // Deprecated
	EffectsV2        *EffectsV2        `json:"effects_v2,omitempty"`
	Dynamics         *Dynamics         `json:"dynamics,omitempty"`
}

type Palette struct {
	Color            []ColorPalette            `json:"color,omitempty"`
	Dimming          []Dimming                 `json:"dimming,omitempty"` // Nutzt jetzt direkt das Shared Struct
	ColorTemperature []ColorTemperaturePalette `json:"color_temperature,omitempty"`
	Effects          []Effects                 `json:"effects,omitempty"` // Deprecated
	EffectsV2        []EffectsV2               `json:"effects_v2,omitempty"`
}

type ColorPalette struct {
	Color   Color   `json:"color"`
	Dimming Dimming `json:"dimming"`
}

// DimmingPalette wurde entfernt, da 'Dimming' (Struct) identisch ist.

type ColorTemperaturePalette struct {
	ColorTemperature ColorTemperature `json:"color_temperature"`
	Dimming          Dimming          `json:"dimming"`
}

type SceneStatus struct {
	Active     string    `json:"active"` // inactive, static, dynamic_palette
	LastRecall time.Time `json:"last_recall,omitempty"`
}

type Scene struct {
	ID          string             `json:"id"`
	IDV1        string             `json:"id_v1,omitempty"`
	Type        string             `json:"type"`
	Metadata    SceneMetadata      `json:"metadata"`
	Group       ResourceIdentifier `json:"group"`
	Actions     []SceneAction      `json:"actions"`
	Palette     Palette            `json:"palette"`
	Speed       float64            `json:"speed"`
	AutoDynamic bool               `json:"auto_dynamic"`
	Status      SceneStatus        `json:"status"`
}

type SceneAction struct {
	Target ResourceIdentifier `json:"target"`
	Action ActionPayload      `json:"action"`
}

type SceneNew struct {
	Metadata    SceneMetadata      `json:"metadata"`
	Group       ResourceIdentifier `json:"group"`
	Actions     []SceneAction      `json:"actions"`
	Palette     *Palette           `json:"palette,omitempty"`
	Speed       *float64           `json:"speed,omitempty"`
	AutoDynamic *bool              `json:"auto_dynamic,omitempty"`
}

type SceneUpdate struct {
	Metadata    *SceneMetadata     `json:"metadata,omitempty"`
	Actions     *[]SceneAction     `json:"actions,omitempty"`
	Palette     *Palette           `json:"palette,omitempty"`
	Speed       *float64           `json:"speed,omitempty"`
	AutoDynamic *bool              `json:"auto_dynamic,omitempty"`
	Recall      *SceneRecallAction `json:"recall,omitempty"`
}

type SceneRecallAction struct {
	Action   string   `json:"action,omitempty"`   // active, dynamic_palette, static
	Duration *int     `json:"duration,omitempty"` // transition time in ms
	Dimming  *Dimming `json:"dimming,omitempty"`  // override brightness
}
