package quortex

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
}
