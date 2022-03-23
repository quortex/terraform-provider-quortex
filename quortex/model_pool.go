package quortex

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
}
