package quortex

// Webhook -
type Webhook struct {
	Uuid     string `json:"uuid,omitempty"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Category string `json:"category,omitempty"`
	PoolUuid string `json:"pool_uuid,omitempty"`
}
