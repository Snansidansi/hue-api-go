package builders

import "github.com/Snansidansi/hue-api-go/models"

type RoomBuilder struct {
	edit models.RoomEdit
}

func NewUpdateRoomBuilder() *RoomBuilder {
	return &RoomBuilder{
		edit: models.RoomEdit{},
	}
}

// children have the light owner.rid and owner.rtype
func NewRoom(name, archetype string, children []models.ResourceIdentifier) *models.RoomEdit {
	metadata := models.MetadataPut{Name: &name, Archetype: &archetype}
	return &models.RoomEdit{
		Children: &children,
		Metadata: &metadata,
	}
}

func (b *RoomBuilder) WithName(name string) *RoomBuilder {
	if b.edit.Metadata == nil {
		b.edit.Metadata = &models.MetadataPut{}
	}
	b.edit.Metadata.Name = &name
	return b
}

func (b *RoomBuilder) WithArchetype(archetype string) *RoomBuilder {
	if b.edit.Metadata == nil {
		b.edit.Metadata = &models.MetadataPut{}
	}
	b.edit.Metadata.Archetype = &archetype
	return b
}

// children have the light owner.rid and owner.rtype
func (b *RoomBuilder) WithChildren(children []models.ResourceIdentifier) *RoomBuilder {
	b.edit.Children = &children
	return b
}

func (b *RoomBuilder) Build() models.RoomEdit {
	return b.edit
}
