package models

import (
	"time"
)

type Light struct {
	ID        string `json:"id"`
	IDV1      string `json:"id_v1,omitempty"`
	Owner     Owner  `json:"owner"`
	Type      string `json:"type"`
	ServiceID int    `json:"service_id"`

	Metadata Metadata `json:"metadata"`

	On       On     `json:"on"`
	Mode     string `json:"mode"`
	Identify *struct {
		Action string `json:"action,omitempty"`
	} `json:"identify,omitempty"`

	Dimming               *Dimming               `json:"dimming,omitempty"`
	DimmingDelta          *DimmingDelta          `json:"dimming_delta,omitempty"`
	ColorTemperature      *ColorTemperature      `json:"color_temperature,omitempty"`
	ColorTemperatureDelta *ColorTemperatureDelta `json:"color_temperature_delta,omitempty"`
	Color                 *Color                 `json:"color,omitempty"`
	Dynamics              *Dynamics              `json:"dynamics,omitempty"`
	Alert                 *Alert                 `json:"alert,omitempty"`
	Signaling             *Signaling             `json:"signaling,omitempty"`
	Gradient              *Gradient              `json:"gradient,omitempty"`
	Effects               *Effects               `json:"effects,omitempty"`
	EffectsV2             *EffectsV2             `json:"effects_v2,omitempty"`
	TimedEffects          *TimedEffects          `json:"timed_effects,omitempty"`
	PowerUp               *PowerUp               `json:"powerup,omitempty"`
}

type LightUpdate struct {
	Type                  *string                `json:"type,omitempty"`
	Metadata              *Metadata              `json:"metadata,omitempty"`
	Identify              *Identify              `json:"identify,omitempty"`
	On                    *On                    `json:"on,omitempty"`
	Dimming               *Dimming               `json:"dimming,omitempty"`
	DimmingDelta          *DimmingDelta          `json:"dimming_delta,omitempty"`
	ColorTemperature      *ColorTemperature      `json:"color_temperature,omitempty"`
	ColorTemperatureDelta *ColorTemperatureDelta `json:"color_temperature_delta,omitempty"`
	Color                 *Color                 `json:"color,omitempty"`
	Dynamics              *Dynamics              `json:"dynamics,omitempty"`
	Alert                 *Alert                 `json:"alert,omitempty"`
	Signaling             *Signaling             `json:"signaling,omitempty"`
	Gradient              *Gradient              `json:"gradient,omitempty"`
	Effects               *Effects               `json:"effects,omitempty"`
	EffectsV2             *EffectsV2             `json:"effects_v2,omitempty"`
	TimedEffects          *TimedEffects          `json:"timed_effects,omitempty"`
	PowerUp               *PowerUp               `json:"powerup,omitempty"`
	ContentConfiguration  *ContentConfiguration  `json:"content_configuration,omitempty"`
}

type Owner struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}

type Metadata struct {
	Name        *string      `json:"name,omitempty"`
	Archetype   *string      `json:"archetype,omitempty"`
	Function    string       `json:"function,omitempty"`
	ProductData *ProductData `json:"product_data,omitempty"`
}

type ProductData struct {
	Name      string `json:"name,omitempty"`
	Archetype string `json:"archetype,omitempty"`
	Function  string `json:"function,omitempty"`
}

type Identify struct {
	Action   *string `json:"action,omitempty"`
	Duration *int64  `json:"duration,omitempty"`
}

type On struct {
	On *bool `json:"on,omitempty"`
}

type Dimming struct {
	Brightness  *float64 `json:"brightness,omitempty"`
	MinDimLevel *float64 `json:"min_dim_level,omitempty"`
}

type DimmingDelta struct {
	Action          *string  `json:"action,omitempty"`
	BrightnessDelta *float64 `json:"brightness_delta,omitempty"`
}

type XY struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ColorTemperature struct {
	Mirek       *int         `json:"mirek,omitempty"`
	MirekValid  *bool        `json:"mirek_valid,omitempty"`
	MirekSchema *MirekSchema `json:"mirek_schema,omitempty"`
}

type MirekSchema struct {
	MirekMinimum int `json:"mirek_minimum"`
	MirekMaximum int `json:"mirek_maximum"`
}

type ColorTemperatureDelta struct {
	Action     *string `json:"action,omitempty"`
	MirekDelta *int    `json:"mirek_delta,omitempty"`
}

type Color struct {
	XY        *XY    `json:"xy,omitempty"`
	Gamut     *Gamut `json:"gamut,omitempty"`
	GamutType string `json:"gamut_type,omitempty"`
}

type Gamut struct {
	Red   XY `json:"red"`
	Green XY `json:"green"`
	Blue  XY `json:"blue"`
}

type Dynamics struct {
	Status       *string  `json:"status,omitempty"`
	StatusValues []string `json:"status_values,omitempty"`
	Speed        *float64 `json:"speed,omitempty"`
	SpeedValid   *bool    `json:"speed_valid,omitempty"`
	Duration     *int64   `json:"duration,omitempty"`
}

type Alert struct {
	Action       *string  `json:"action,omitempty"`
	ActionValues []string `json:"action_values,omitempty"`
}

type Signaling struct {
	Signal       *string             `json:"signal,omitempty"`
	Duration     *int64              `json:"duration,omitempty"`
	SignalValues []string            `json:"signal_values,omitempty"`
	Status       *SignalingStatus    `json:"status,omitempty"`
	Colors       []ColorFeatureBasic `json:"colors,omitempty"`
}

type SignalingStatus struct {
	Signal       string     `json:"signal"`
	EstimatedEnd *time.Time `json:"estimated_end,omitempty"`
}

type ColorFeatureBasic struct {
	XY *XY `json:"xy,omitempty"`
}

type Gradient struct {
	Points        []GradientPoint `json:"points,omitempty"`
	Mode          *string         `json:"mode,omitempty"`
	PointsCapable *int            `json:"points_capable,omitempty"`
	ModeValues    []string        `json:"mode_values,omitempty"`
	PixelCount    *int            `json:"pixel_count,omitempty"`
}

type GradientPoint struct {
	Color *ColorFeatureBasic `json:"color,omitempty"`
}

type Effects struct {
	Effect       *string  `json:"effect,omitempty"`
	Status       *string  `json:"status,omitempty"`
	StatusValues []string `json:"status_values,omitempty"`
	EffectValues []string `json:"effect_values,omitempty"`
}

type EffectsV2 struct {
	Action       *EffectActionV2 `json:"action,omitempty"`
	Status       *EffectStatus   `json:"status,omitempty"`
	EffectValues []string        `json:"effect_values,omitempty"`
}

type EffectStatus struct {
	Effect string `json:"effect"`
}

type EffectActionV2 struct {
	Effect     *string           `json:"effect,omitempty"`
	Parameters *EffectParameters `json:"parameters,omitempty"`
}

type EffectParameters struct {
	Color            *ColorFeatureBasic `json:"color,omitempty"`
	ColorTemperature *ColorTemperature  `json:"color_temperature,omitempty"`
	Speed            *float64           `json:"speed,omitempty"`
}

type TimedEffects struct {
	Effect       *string  `json:"effect,omitempty"`
	Duration     *int64   `json:"duration,omitempty"`
	Status       *string  `json:"status,omitempty"`
	StatusValues []string `json:"status_values,omitempty"`
	EffectValues []string `json:"effect_values,omitempty"`
}

type PowerUp struct {
	Preset               *string               `json:"preset,omitempty"`
	Configured           *bool                 `json:"configured,omitempty"`
	On                   *PowerUpOn            `json:"on,omitempty"`
	Dimming              *PowerUpDimming       `json:"dimming,omitempty"`
	Color                *PowerUpColor         `json:"color,omitempty"`
	ContentConfiguration *ContentConfiguration `json:"content_configuration,omitempty"`
}

type PowerUpOn struct {
	Mode *string `json:"mode,omitempty"`
	On   *On     `json:"on,omitempty"`
}

type PowerUpDimming struct {
	Mode    *string  `json:"mode,omitempty"`
	Dimming *Dimming `json:"dimming,omitempty"`
}

type PowerUpColor struct {
	Mode             *string           `json:"mode,omitempty"`
	ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
	Color            *Color            `json:"color,omitempty"`
}

type ContentConfiguration struct {
	Orientation *Orientation `json:"orientation,omitempty"`
	Order       *Order       `json:"order,omitempty"`
}

type Orientation struct {
	Status       *string `json:"status,omitempty"`
	Configurable *bool   `json:"configurable,omitempty"`
	Orientation  *string `json:"orientation,omitempty"`
}

type Order struct {
	Status       *string `json:"status,omitempty"`
	Configurable *bool   `json:"configurable,omitempty"`
	Order        *string `json:"order,omitempty"`
}
