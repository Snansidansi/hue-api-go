package models

type ZoneEdit struct {
	Children *[]ResourceIdentifier `json:"children,omitempty"`
	Metadata *MetadataPut          `json:"metadata,omitempty"`
}

type Zone struct {
	ID   string `json:"id"`
	IDV1 string `json:"id_v1"`
	// children have the direct light id and type not the owner.rid and owner.rtype
	Children []ResourceIdentifier `json:"children"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
	} `json:"metadata"`
	Services []ResourceIdentifier `json:"services"`
	Type     string               `json:"type"`
}
