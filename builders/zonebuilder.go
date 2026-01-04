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

// children have the direct light id and type not the owner.rid and owner.rtype
func NewZone(name, archetype string, children []models.ResourceIdentifier) *models.ZoneEdit {
	metadata := models.MetadataPut{Name: &name, Archetype: &archetype}
	return &models.ZoneEdit{
		Children: &children,
		Metadata: &metadata,
	}
}

func (b *ZoneBuilder) WithName(name string) *ZoneBuilder {
	if b.edit.Metadata == nil {
		b.edit.Metadata = &models.MetadataPut{}
	}
	b.edit.Metadata.Name = &name
	return b
}

// children have the direct light id and type not the owner.rid and owner.rtype
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
