package builders

import (
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
)

type GroupedLightBuilder struct {
	update models.GroupedLightUpdate
}

func NewGroupedLightBuilder() *GroupedLightBuilder {
	return &GroupedLightBuilder{
		update: models.GroupedLightUpdate{},
	}
}

func (b *GroupedLightBuilder) On(on bool) *GroupedLightBuilder {
	if b.update.On == nil {
		b.update.On = &models.On{}
	}
	b.update.On.On = &on
	return b
}

func (b *GroupedLightBuilder) Brightness(percent float64) *GroupedLightBuilder {
	if b.update.Dimming == nil {
		b.update.Dimming = &models.Dimming{}
	}
	b.update.Dimming.Brightness = &percent
	return b
}

func (b *GroupedLightBuilder) BrightnessDelta(action string, delta float64) *GroupedLightBuilder {
	if b.update.DimmingDelta == nil {
		b.update.DimmingDelta = &models.DimmingDelta{}
	}
	b.update.DimmingDelta.Action = &action
	b.update.DimmingDelta.BrightnessDelta = &delta
	return b
}

func (b *GroupedLightBuilder) Temperature(mirek int) *GroupedLightBuilder {
	if b.update.ColorTemperature == nil {
		b.update.ColorTemperature = &models.ColorTemperature{}
	}
	b.update.ColorTemperature.Mirek = &mirek
	return b
}

func (b *GroupedLightBuilder) TemperatureDelta(action string, delta int) *GroupedLightBuilder {
	if b.update.ColorTemperatureDelta == nil {
		b.update.ColorTemperatureDelta = &models.ColorTemperatureDelta{}
	}
	b.update.ColorTemperatureDelta.Action = &action
	b.update.ColorTemperatureDelta.MirekDelta = &delta
	return b
}

func (b *GroupedLightBuilder) ColorXY(x, y float64) *GroupedLightBuilder {
	if b.update.Color == nil {
		b.update.Color = &models.Color{XY: &models.XY{}}
	}
	b.update.Color.XY.X = x
	b.update.Color.XY.Y = y
	return b
}

func (b *GroupedLightBuilder) ColorRGB(r, g, b_int int) *GroupedLightBuilder {
	x, y := util.RGBToXY(r, g, b_int)
	return b.ColorXY(x, y)
}

func (b *GroupedLightBuilder) Alert() *GroupedLightBuilder {
	if b.update.Alert == nil {
		b.update.Alert = &models.Alert{}
	}

	action := "breathe"
	b.update.Alert.Action = &action
	return b
}

// durationMS must be greater than 999
func (b *GroupedLightBuilder) Signaling(signal string, durationMS int64) *GroupedLightBuilder {
	if b.update.Signaling == nil {
		b.update.Signaling = &models.Signaling{}
	}
	if durationMS < 1000 {
		durationMS = 1000
	}

	b.update.Signaling.Signal = &signal
	b.update.Signaling.Duration = &durationMS
	return b
}

func (b *GroupedLightBuilder) SignalingXY(signal string, durationMS int64, x, y float64) *GroupedLightBuilder {
	b.Signaling(signal, durationMS)

	colors := []models.ColorFeatureBasic{
		{XY: &models.XY{X: x, Y: y}},
	}
	b.update.Signaling.Colors = colors

	return b
}

func (b *GroupedLightBuilder) SignalingRGB(signal string, duration int64, r, g, b_int int) *GroupedLightBuilder {
	x, y := util.RGBToXY(r, g, b_int)
	return b.SignalingXY(signal, duration, x, y)
}

func (b *GroupedLightBuilder) Dynamics(durationMs int64) *GroupedLightBuilder {
	if b.update.Dynamics == nil {
		b.update.Dynamics = &models.Dynamics{}
	}
	b.update.Dynamics.Duration = &durationMs
	return b
}

func (b *GroupedLightBuilder) Build() models.GroupedLightUpdate {
	return b.update
}
