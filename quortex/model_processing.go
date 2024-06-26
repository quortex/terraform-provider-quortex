package quortex

// Resolution -
type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Advanced -
type Advanced struct {
	Profile             string `json:"profile,omitempty"`
	Level               string `json:"level,omitempty"`
	Quality             string `json:"quality,omitempty"`
	EncodingMode        string `json:"encoding_mode,omitempty"`
	EncodingQuality     int    `json:"encoding_quality,omitempty"`
	QualityOptimization string `json:"quality_optimization,omitempty"`
	ClosedGop           *bool  `json:"closed_gop,omitempty"`
	GopSize             int    `json:"gop_size,omitempty"`
	GopMaxSize          int    `json:"gop_max_size,omitempty"`
	Bframe              *bool  `json:"bframe,omitempty"`
	BframeNumber        int    `json:"bframe_number,omitempty"`
	Maxrate             int    `json:"maxrate,omitempty"`
	KeyFrameIntervalMs  int    `json:"key_frame_interval_ms,omitempty"`
	HorizontalSharpness int    `json:"horizontal_sharpness,omitempty"`
	VerticalSharpness   int    `json:"vertical_sharpness,omitempty"`
	LogoEnabled         bool   `json:"logo_enabled,omitempty"`
}

// VideoMedia -
type VideoMedia struct {
	Label      string      `json:"label,omitempty"`
	Codec      string      `json:"codec"`
	Bitrate    int         `json:"bitrate"`
	Framerate  string      `json:"framerate"`
	Resolution *Resolution `json:"resolution"`
	Advanced   *Advanced   `json:"advanced,omitempty"`
}

// AudioMedia -
type AudioMedia struct {
	Label            string `json:"label,omitempty"`
	Codec            string `json:"codec"`
	Channels         string `json:"channels"`
	Bitrate          int    `json:"bitrate"`
	Samplerate       string `json:"samplerate"`
	Track            string `json:"track"`
	Output           string `json:"output"`
	AudioDescription bool   `json:"audio_description"`
}

// SubtitleMedia -
type SubtitleMedia struct {
	Track                string `json:"track"`
	Output               string `json:"output"`
	DeafAndHardOfHearing bool   `json:"deaf_and_hard_of_hearing"`
}

// Processing -
type Processing struct {
	Uuid           string          `json:"uuid,omitempty"`
	Name           string          `json:"name"`
	Identifier     string          `json:"identifier,omitempty"`
	Published      bool            `json:"published"`
	VideoMedias    []VideoMedia    `json:"video_medias,omitempty"`
	AudioMedias    []AudioMedia    `json:"audio_medias,omitempty"`
	SubtitleMedias []SubtitleMedia `json:"subtitle_medias,omitempty"`
	Labels         []string        `json:"labels,omitempty"`
}
