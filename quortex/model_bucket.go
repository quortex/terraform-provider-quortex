package quortex

// S3 -
type S3 struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// Bucket -
type Bucket struct {
	Uuid   string `json:"uuid,omitempty"`
	Name   string `json:"name"`
	Region string `json:"region"`
	Type   string `json:"type"`
	Label  string `json:"label,omitempty"`
	S3     *S3    `json:"s3,omitempty"`
}
