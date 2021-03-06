package quortex

// Castlabs -
type Castlabs struct {
	MerchantName string `json:"merchant_name,omitempty"`
	AesIv        string `json:"aes_iv,omitempty"`
	AesKey       string `json:"aes_key,omitempty"`
	DrmServer    string `json:"drm_server,omitempty"`
	KeySeedId    string `json:"key_seed_id,omitempty"`
	AuthCredsId  string `json:"auth_creds_id,omitempty"`
}

// DrmMerchant -
type DrmMerchant struct {
	Uuid     string    `json:"uuid,omitempty"`
	Name     string    `json:"name,omitempty"`
	Type     string    `json:"type,omitempty"`
	Castlabs *Castlabs `json:"castlabs,omitempty"`
}
