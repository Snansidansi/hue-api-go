package builders

import "github.com/snansidansi/hueapi/models"

type ZoneBuilder struct {
	edit models.ZoneEdit
}

func NewUpdateZoneBuilder() *ZoneBuilder {
	return &ZoneBuilder{
		edit: models.ZoneEdit{},
	}
}

func NewCreateZoneBuilder(name string, children []models.ResourceIdentifier) *ZoneBuilder {
	metadata := models.MetadataPut{Name: &name}
	return &ZoneBuilder{
		edit: models.ZoneEdit{
			Children: &children,
			Metadata: &metadata,
		},
	}
}

func (b *ZoneBuilder) WithName(name string) *ZoneBuilder {
	if b.edit.Metadata == nil {
		b.edit.Metadata = &models.MetadataPut{}
	}
	b.edit.Metadata.Name = &name
	return b
}

func (b *ZoneBuilder) WithArchetype(archetype string) *ZoneBuilder {
	if b.edit.Metadata == nil {
		b.edit.Metadata = &models.MetadataPut{}
	}
	b.edit.Metadata.Archetype = &archetype
	return b
}

func (b *ZoneBuilder) WithChildren(children []models.ResourceIdentifier) *ZoneBuilder {
	b.edit.Children = &children
	return b
}

func (b *ZoneBuilder) Build() models.ZoneEdit {
	return b.edit
}
