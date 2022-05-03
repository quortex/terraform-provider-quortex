package quortex

// DashAdvanced -
type DashAdvanced struct {
	UtcTiming                  string `json:"utc_timing,omitempty"`
	UtcTimingServer            string `json:"utc_timing_server,omitempty"`
	BaseUrl                    string `json:"base_url,omitempty"`
	Position                   string `json:"position,omitempty"`
	SuggestedPresentationDelay int    `json:"suggested_presentation_delay,omitempty"`
	StartTimeOriginOffset      int    `json:"start_time_origin_offset,omitempty"`
}

// HlsAdvanced -
type HlsAdvanced struct {
	ProgramDatetime string `json:"program_datetime,omitempty"`
	Version         int    `json:"version,omitempty"`
}

// Encryption -
type Encryption struct {
	Uuid       string   `json:"uuid,omitempty"`
	Iv         string   `json:"iv"`
	IvMode     string   `json:"iv_mode"`
	StreamType string   `json:"stream_type"`
	Labels     []string `json:"labels"`
}

// EncryptionDynamic -
type EncryptionDynamic struct {
	ContentId       string       `json:"content_id"`
	DrmMerchantUuid string       `json:"drm_merchant_uuid"`
	Encryption      []Encryption `json:"encryptions"`
}

// Scte35 -
type Scte35 struct {
	Enabled    bool     `json:"enabled,omitempty"`
	FilterType string   `json:"filter_type,omitempty"`
	FilterList []string `json:"filter_list,omitempty"`
}

// Target -
type Target struct {
	Uuid                       string             `json:"uuid,omitempty"`
	Name                       string             `json:"name"`
	Identifier                 string             `json:"identifier,omitempty"`
	Published                  bool               `json:"published"`
	Type                       string             `json:"type"`
	SegmentDuration            float64            `json:"segment_duration"`
	PlaylistLength             int                `json:"playlist_length"`
	Container                  string             `json:"container,omitempty"`
	Scte35                     *Scte35            `json:"scte_35,omitempty"`
	InputLabelRestriction      []string           `json:"input_label_restriction,omitempty"`
	ProcessingLabelRestriction []string           `json:"processing_label_restriction,omitempty"`
	EncryptionType             string             `json:"encryption_type,omitempty"`
	EncryptionDynamic          *EncryptionDynamic `json:"encryption_dynamic,omitempty"`
	DashAdvanced               *DashAdvanced      `json:"dash_advanced,omitempty"`
	HlsAdvanced                *HlsAdvanced       `json:"hls_advanced,omitempty"`
}
