package quortex

// Pool -
type Pool struct {
	Uuid               string   `json:"uuid,omitempty"`
	Name               string   `json:"name"`
	Published          bool     `json:"published"`
	InputRegion        string   `json:"input_region"`
	StreamingCountries []string `json:"streaming_countries"`
}

// Caller -
type Caller struct {
}

// Listener -
type Listener struct {
}

// Srt -
type Srt struct {
	ConnectionType string    `json:"connection_type"`
	Latency        int       `json:"latency,omitempty"`
	Caller         *Caller   `json:"caller,omitempty"`
	Listener       *Listener `json:"listener,omitempty"`
}

// Rtmp -
type Rtmp struct{}

// Stream -
type Stream struct {
	Uuid    string `json:"uuid,omitempty"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled,omitempty"`
	Srt     *Srt   `json:"srt,omitempty"`
	Rtmp    *Rtmp  `json:"rtmp,omitempty"`
}

// Input -
type Input struct {
	Uuid       string   `json:"uuid,omitempty"`
	Name       string   `json:"name"`
	Identifier string   `json:"identifier,omitempty"`
	Published  bool     `json:"published"`
	Streams    []Stream `json:"streams"`
}
