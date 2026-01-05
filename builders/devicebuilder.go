package builders

import "github.com/snansidansi/hueapi/models"

type DeviceBuilder struct {
	put models.DevicePut
}

func NewDeviceBuilder() *DeviceBuilder {
	return &DeviceBuilder{
		put: models.DevicePut{},
	}
}

func (b *DeviceBuilder) SetName(name string) *DeviceBuilder {
	if b.put.Metadata == nil {
		b.put.Metadata = &models.DeviceMetadata{}
	}
	b.put.Metadata.Name = &name
	return b
}

func (b *DeviceBuilder) SetArchetype(archetype string) *DeviceBuilder {
	if b.put.Metadata == nil {
		b.put.Metadata = &models.DeviceMetadata{}
	}
	b.put.Metadata.Archetype = &archetype
	return b
}

func (b *DeviceBuilder) Identify(action string) *DeviceBuilder {
	if b.put.Identify == nil {
		b.put.Identify = &models.Identify{}
	}
	b.put.Identify.Action = &action
	return b
}

func (b *DeviceBuilder) UserTest(enabled bool) *DeviceBuilder {
	if b.put.UserTest == nil {
		b.put.UserTest = &models.UserTest{}
	}
	b.put.UserTest.UserTest = enabled
	return b
}

func (b *DeviceBuilder) DeviceMode(mode string) *DeviceBuilder {
	if b.put.DeviceMode == nil {
		b.put.DeviceMode = &models.DeviceMode{}
	}
	b.put.DeviceMode.Mode = mode
	return b
}

func (b *DeviceBuilder) Build() models.DevicePut {
	return b.put
}
