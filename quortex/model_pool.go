package quortex

// Bucket -
type Bucket2 struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Type   string `json:"type"`
	S3     *S3    `json:"s3,omitempty"`
}

// Catchup -
type Catchup struct {
	Enabled bool     `json:"enabled,omitempty"`
	Bucket2 *Bucket2 `json:"bucket,omitempty"`
}

// TimeShifting -
type TimeShifting struct {
	Enabled           bool `json:"enabled,omitempty"`
	StartoverDuration int  `json:"startover_duration,omitempty"`
}

// Pool -
type Pool struct {
	Uuid               string        `json:"uuid,omitempty"`
	Name               string        `json:"name"`
	Published          bool          `json:"published"`
	InputRegion        string        `json:"input_region"`
	StreamingCountries []string      `json:"streaming_countries"`
	Label              string        `json:"label,omitempty"`
	TimeShifting       *TimeShifting `json:"time_shifting,omitempty"`
	Catchup            *Catchup      `json:"catchup,omitempty"`
	Origin             *Origin       `json:"origin,omitempty"`
}

// Origin -
type Origin struct {
	Enabled                    bool     `json:"enabled"`
	WhitelistEnabled           bool     `json:"whitelist_enabled"`
	WhitelistCidr              []string `json:"whitelist_cidr"`
	AuthorizationHeaderEnabled bool     `json:"authorization_header_enabled"`
	AuthorizationHeaderValue   string   `json:"authorization_header_value"`
}
