package builders

import (
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
)

type SceneActionBuilder struct {
	payload models.ActionPayload
	target  models.ResourceIdentifier
}

func (b *SceneActionBuilder) SetOnOff(state bool) *SceneActionBuilder {
	if b.payload.On == nil {
		b.payload.On = &models.On{}
	}
	b.payload.On.On = &state
	return b
}

func (b *SceneActionBuilder) Brightness(brightness float64) *SceneActionBuilder {
	if b.payload.Dimming == nil {
		b.payload.Dimming = &models.Dimming{}
	}
	b.payload.Dimming.Brightness = &brightness
	return b
}

func (b *SceneActionBuilder) ColorXY(x, y float64) *SceneActionBuilder {
	if b.payload.Color == nil {
		b.payload.Color = &models.Color{}
	}
	if b.payload.Color.XY == nil {
		b.payload.Color.XY = &models.XY{}
	}
	b.payload.Color.XY.X = x
	b.payload.Color.XY.Y = y
	return b
}

func (b *SceneActionBuilder) ColorRGB(r, g, b_int int) *SceneActionBuilder {
	x, y := util.RGBToXY(r, g, b_int)
	return b.ColorXY(x, y)
}

func (b *SceneActionBuilder) ColorTemp(mirek int) *SceneActionBuilder {
	if b.payload.ColorTemperature == nil {
		b.payload.ColorTemperature = &models.ColorTemperature{}
	}
	b.payload.ColorTemperature.Mirek = &mirek
	return b
}

func (b *SceneActionBuilder) Build() models.SceneAction {
	return models.SceneAction{
		Target: b.target,
		Action: b.payload,
	}
}

type RecallBuilder struct {
	recall models.SceneRecallAction
}

func NewRecallBuilder() *RecallBuilder {
	return &RecallBuilder{
		recall: models.SceneRecallAction{},
	}
}

// Action accepts: "active", "dynamic_palette", "static"
func (b *RecallBuilder) WithAction(actionType string) *RecallBuilder {
	b.recall.Action = actionType
	return b
}

func (b *RecallBuilder) WithDuration(durationMs int) *RecallBuilder {
	b.recall.Duration = &durationMs
	return b
}

func (b *RecallBuilder) WithBrightnessOverride(brightness float64) *RecallBuilder {
	if b.recall.Dimming == nil {
		b.recall.Dimming = &models.Dimming{}
	}
	b.recall.Dimming.Brightness = &brightness
	return b
}

func (b *RecallBuilder) Build() models.SceneRecallAction {
	return b.recall
}

type SceneCreateBuilder struct {
	scene models.SceneNew
}

func NewCreateSceneBuilder(name string, groupResource models.ResourceIdentifier) *SceneCreateBuilder {
	return &SceneCreateBuilder{
		scene: models.SceneNew{
			Metadata: models.SceneMetadata{Name: name},
			Group:    groupResource,
			Actions:  []models.SceneAction{},
		},
	}
}

// light ID and Type is used insted of Owner.RID and Owner.RType
// Example: .AddAction(target, func(ab *SceneActionBuilder) { ab.SetOnOff(true) })
func (b *SceneCreateBuilder) AddAction(targetResource models.ResourceIdentifier, setupFunc func(*SceneActionBuilder)) *SceneCreateBuilder {
	ab := &SceneActionBuilder{
		target:  targetResource,
		payload: models.ActionPayload{},
	}

	setupFunc(ab)
	b.scene.Actions = append(b.scene.Actions, ab.Build())
	return b
}

func (b *SceneCreateBuilder) WithSpeed(speed float64) *SceneCreateBuilder {
	b.scene.Speed = &speed
	return b
}

func (b *SceneCreateBuilder) WithAutoDynamic(auto bool) *SceneCreateBuilder {
	b.scene.AutoDynamic = &auto
	return b
}

func (b *SceneCreateBuilder) Build() *models.SceneNew {
	return &b.scene
}

type SceneUpdateBuilder struct {
	scene models.SceneUpdate
}

func NewUpdateSceneBuilder() *SceneUpdateBuilder {
	return &SceneUpdateBuilder{
		scene: models.SceneUpdate{},
	}
}

func (b *SceneUpdateBuilder) SetName(name string) *SceneUpdateBuilder {
	if b.scene.Metadata == nil {
		b.scene.Metadata = &models.SceneMetadata{}
	}
	b.scene.Metadata.Name = name
	return b
}

func (b *SceneUpdateBuilder) AddAction(targetResource models.ResourceIdentifier, setupFunc func(*SceneActionBuilder)) *SceneUpdateBuilder {
	ab := &SceneActionBuilder{
		target:  targetResource,
		payload: models.ActionPayload{},
	}

	setupFunc(ab)

	if b.scene.Actions == nil {
		b.scene.Actions = &[]models.SceneAction{}
	}
	*b.scene.Actions = append(*b.scene.Actions, ab.Build())

	return b
}

func (b *SceneUpdateBuilder) SetRecall(rb *RecallBuilder) *SceneUpdateBuilder {
	recall := rb.Build()
	b.scene.Recall = &recall
	return b
}

func (b *SceneUpdateBuilder) SetRecallDirect(action string) *SceneUpdateBuilder {
	b.scene.Recall = &models.SceneRecallAction{
		Action: action,
	}
	return b
}

func (b *SceneUpdateBuilder) Build() *models.SceneUpdate {
	return &b.scene
}
