package models

type BridgeHome struct {
	ID       string               `json:"id"`
	IDV1     string               `json:"id_v1,omitempty"`
	Type     string               `json:"type"` // "bridge_home"
	Children []ResourceIdentifier `json:"children"`
	Services []ResourceIdentifier `json:"services"`
}
