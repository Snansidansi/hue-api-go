package models

type Entertainment struct {
	ID                string                 `json:"id"`
	IDV1              string                 `json:"id_v1,omitempty"`
	Owner             ResourceIdentifier     `json:"owner"`
	Type              string                 `json:"type"` // "entertainment"
	Renderer          bool                   `json:"renderer"`
	RendererReference *ResourceIdentifier    `json:"renderer_reference,omitempty"`
	Proxy             bool                   `json:"proxy"`
	Equalizer         bool                   `json:"equalizer"`
	MaxStreams        *int                   `json:"max_streams,omitempty"`
	Segments          *EntertainmentSegments `json:"segments,omitempty"`
}

type EntertainmentSegments struct {
	Configurable bool                   `json:"configurable"`
	MaxSegments  int                    `json:"max_segments"`
	Segments     []EntertainmentSegment `json:"segments"`
}

type EntertainmentSegment struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}
