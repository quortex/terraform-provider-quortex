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

// Irdeto -
type Irdeto struct {
	MerchantName string `json:"merchant_name,omitempty"`
	DrmServer    string `json:"drm_server,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}

// Ksm -
type Ksm struct {
	DrmServer         string `json:"drm_server,omitempty"`
	ClientCertificate string `json:"client_certificate,omitempty"`
	ClientKey         string `json:"client_key,omitempty"`
}

// Mdrm -
type Mdrm struct {
	AuthServer   string `json:"auth_server,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	DrmServer    string `json:"drm_server,omitempty"`
}

// DrmMerchant -
type DrmMerchant struct {
	Uuid     string    `json:"uuid,omitempty"`
	Name     string    `json:"name,omitempty"`
	Type     string    `json:"type,omitempty"`
	Castlabs *Castlabs `json:"castlabs,omitempty"`
	Irdeto   *Irdeto   `json:"irdeto,omitempty"`
	Ksm      *Ksm      `json:"ksm,omitempty"`
	Mdrm     *Mdrm     `json:"mdrm,omitempty"`
}
