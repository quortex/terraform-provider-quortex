package quortex

// Scte35 -
type Scte35 struct {
	Enabled    bool     `json:"enabled,omitempty"`
	FilterType string   `json:"filter_type,omitempty"`
	FilterList []string `json:"filter_list,omitempty"`
}

// Target -
type Target struct {
	Uuid            string  `json:"uuid,omitempty"`
	Name            string  `json:"name"`
	Identifier      string  `json:"identifier,omitempty"`
	Published       bool    `json:"published"`
	Type            string  `json:"type"`
	SegmentDuration float64 `json:"segment_duration"`
	PlaylistLength  int     `json:"playlist_length"`
	Container       string  `json:"container,omitempty"`
	Scte35          *Scte35 `json:"scte_35,omitempty"`
}
