package quortex

// PublishingPoint -
type PublishingPoint struct {
	Uuid           string `json:"uuid,omitempty"`
	TargetType     string `json:"target_type,omitempty"`
	InputUuid      string `json:"input_uuid,omitempty"`
	ProcessingUuid string `json:"processing_uuid,omitempty"`
	TargetUuid     string `json:"target_uuid,omitempty"`
	CustomPath     string `json:"custom_path"`
	Published      bool   `json:"published"`
}
