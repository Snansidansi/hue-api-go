package builders

import "github.com/snansidansi/hueapi/models"

type RoomBuilder struct {
	edit models.RoomEdit
}

func NewUpdateRoomBuilder() *RoomBuilder {
	return &RoomBuilder{
		edit: models.RoomEdit{},
	}
}

func NewCreateRoomBuilder(name string, children []models.ResourceIdentifier) *RoomBuilder {
	metadata := models.MetadataPut{Name: &name}
	return &RoomBuilder{
		edit: models.RoomEdit{
			Children: &children,
			Metadata: &metadata,
		},
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

func (b *RoomBuilder) WithChildren(children []models.ResourceIdentifier) *RoomBuilder {
	b.edit.Children = &children
	return b
}

func (b *RoomBuilder) Build() models.RoomEdit {
	return b.edit
}
