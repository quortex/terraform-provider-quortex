package quortex

// OTTEndpoint -
type OTTEndpoint struct {
	Uuid           string `json:"uuid,omitempty"`
	Enabled        bool   `json:"enabled"`
	CustomPath     string `json:"custom_path"`
	InputUuid      string `json:"input_uuid,omitempty"`
	ProcessingUuid string `json:"processing_uuid,omitempty"`
	TargetUuid     string `json:"target_uuid,omitempty"`
}
