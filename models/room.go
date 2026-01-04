package models

type RoomEdit struct {
	Children *[]ResourceIdentifier `json:"children,omitempty"`
	Metadata *MetadataPut          `json:"metadata,omitempty"`
}

type Room struct {
	ID   string `json:"id"`
	IDV1 string `json:"id_v1"`
	// children have the light owner.rid and owner.rtype
	Children []ResourceIdentifier `json:"children"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
	} `json:"metadata"`
	Services []ResourceIdentifier `json:"services"`
	Type     string               `json:"type"`
}
