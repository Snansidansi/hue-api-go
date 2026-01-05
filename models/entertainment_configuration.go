package models

type EntertainmentConfiguration struct {
	ID                string                 `json:"id"`
	IDV1              string                 `json:"id_v1,omitempty"`
	Type              string                 `json:"type"`
	Metadata          EntertainmentMetadata  `json:"metadata"`
	ConfigurationType string                 `json:"configuration_type"` // screen, monitor, music, 3dspace, other
	Status            string                 `json:"status"`             // active, inactive
	ActiveStreamer    *ResourceIdentifier    `json:"active_streamer,omitempty"`
	StreamProxy       StreamProxy            `json:"stream_proxy"`
	Channels          []EntertainmentChannel `json:"channels"`
	Locations         EntertainmentLocations `json:"locations"`
}

type EntertainmentConfigurationNew struct {
	Type              *string                `json:"type,omitempty"` // "entertainment_configuration"
	Metadata          EntertainmentMetadata  `json:"metadata"`
	ConfigurationType string                 `json:"configuration_type"`
	StreamProxy       *StreamProxy           `json:"stream_proxy,omitempty"`
	Locations         EntertainmentLocations `json:"locations"`
}

type EntertainmentConfigurationUpdate struct {
	Type              *string                 `json:"type,omitempty"` // "entertainment_configuration"
	Metadata          *EntertainmentMetadata  `json:"metadata,omitempty"`
	ConfigurationType *string                 `json:"configuration_type,omitempty"`
	Action            *string                 `json:"action,omitempty"` // start, stop
	StreamProxy       *StreamProxy            `json:"stream_proxy,omitempty"`
	Locations         *EntertainmentLocations `json:"locations,omitempty"`
}

type EntertainmentMetadata struct {
	Name string `json:"name"`
}

type StreamProxy struct {
	Mode string              `json:"mode"` // auto, manual
	Node *ResourceIdentifier `json:"node,omitempty"`
}

type EntertainmentLocations struct {
	ServiceLocations []EntertainmentServiceLocation `json:"service_locations"`
}

type EntertainmentServiceLocation struct {
	Service            ResourceIdentifier `json:"service"`
	Positions          []Position         `json:"positions"`
	EqualizationFactor *float64           `json:"equalization_factor,omitempty"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type EntertainmentChannel struct {
	ChannelID int                `json:"channel_id"`
	Position  Position           `json:"position"`
	Members   []SegmentReference `json:"members"`
}

type SegmentReference struct {
	Service ResourceIdentifier `json:"service"`
	Index   int                `json:"index"`
}
