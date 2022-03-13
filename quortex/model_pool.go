package quortex

// Pool -
type Pool struct {
	Uuid               string   `json:"uuid,omitempty"`
	Name               string   `json:"name"`
	Published          bool     `json:"published"`
	InputRegion        string   `json:"input_region"`
	StreamingCountries []string `json:"streaming_countries"`
}
