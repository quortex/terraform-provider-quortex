package quortex

// Audio -
type OverAudio struct {
	Language string `json:"language,omitempty"`
	Ad       bool   `json:"ad,omitempty"`
}

type OverTeletext struct {
	Page     string `json:"page"`
	Language string `json:"language,omitempty"`
	Sdh      bool   `json:"sdh,omitempty"`
}

// Override -
type Override struct {
	Pid          int           `json:"pid,omitempty"`
	Type         string        `json:"type"`
	Enabled      bool          `json:"enabled,omitempty"`
	OverAudio    *OverAudio    `json:"audio,omitempty"`
	OverTeletext *OverTeletext `json:"teletext,omitempty"`
}

// Caller -
type Caller struct {
	Address    string `json:"address"`
	Passphrase string `json:"passphrase,omitempty"`
}

// Listener -
type Listener struct {
	Cidr []string `json:"cidr,omitempty"`
}

// Srt -
type Srt struct {
	ConnectionType string     `json:"connection_type"`
	Listener       *Listener  `json:"listener,omitempty"`
	Caller         *Caller    `json:"caller,omitempty"`
	Latency        int        `json:"latency,omitempty"`
	Overrides      []Override `json:"overrides,omitempty"`
}

// Rtmp -
type Rtmp struct {
	Overrides []Override `json:"overrides,omitempty"`
}

// Stream -
type Stream struct {
	Uuid        string `json:"uuid,omitempty"`
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled,omitempty"`
	LogoUrl     string `json:"logo_url,omitempty"`
	FallbackUrl string `json:"fallback_url,omitempty"`
	Type        string `json:"type"`
	Srt         *Srt   `json:"srt,omitempty"`
	Rtmp        *Rtmp  `json:"rtmp,omitempty"`
}

// Input -
type Input struct {
	Uuid       string   `json:"uuid,omitempty"`
	Name       string   `json:"name"`
	Identifier string   `json:"identifier,omitempty"`
	Published  bool     `json:"published"`
	Streams    []Stream `json:"streams,omitempty"`
	Labels     []string `json:"labels,omitempty"`
}
