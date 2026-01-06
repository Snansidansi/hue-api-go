package builders

import (
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
)

type LightBuilder struct {
	put models.LightUpdate
}

func NewLightBuilder() *LightBuilder {
	return &LightBuilder{
		put: models.LightUpdate{},
	}
}

func (b *LightBuilder) On() *LightBuilder {
	state := true
	if b.put.On == nil {
		b.put.On = &models.On{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) Off() *LightBuilder {
	state := false
	if b.put.On == nil {
		b.put.On = &models.On{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) SetOnOff(state bool) *LightBuilder {
	if b.put.On == nil {
		b.put.On = &models.On{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) Brightness(percent float64) *LightBuilder {
	if b.put.Dimming == nil {
		b.put.Dimming = &models.Dimming{}
	}
	b.put.Dimming.Brightness = &percent
	return b
}

func (b *LightBuilder) ColorXY(x, y float64) *LightBuilder {
	if b.put.Color == nil {
		b.put.Color = &models.Color{XY: &models.XY{}}
	}
	b.put.Color.XY.X = x
	b.put.Color.XY.Y = y
	return b
}

func (b *LightBuilder) ColorRGB(r, g, b_int int) *LightBuilder {
	x, y := util.RGBToXY(r, g, b_int)
	return b.ColorXY(x, y)
}

func (b *LightBuilder) Temperature(mirek int) *LightBuilder {
	if b.put.ColorTemperature == nil {
		b.put.ColorTemperature = &models.ColorTemperature{}
	}
	b.put.ColorTemperature.Mirek = &mirek
	return b
}

func (b *LightBuilder) Duration(ms int64) *LightBuilder {
	if b.put.Dynamics == nil {
		b.put.Dynamics = &models.Dynamics{}
	}
	b.put.Dynamics.Duration = &ms
	return b
}

func (b *LightBuilder) Build() models.LightUpdate {
	return b.put
}
