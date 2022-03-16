package quortex

// Dataplane -
type Dataplane struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Organization       string `json:"organization,omitempty"`
	Provider           string `json:"provider,omitempty"`
	Region             string `json:"region"`
	Endpoint           string `json:"endpoint"`
	Certificate        string `json:"certificate,omitempty"`
	Token              string `json:"token,omitempty"`
	Livepoint          string `json:"livepoint,omitempty"`
	Rtmpendpoint       string `json:"rtmpendpoint,omitempty"`
	Enable             bool   `json:"enable"`
	ManageDistribution bool   `json:"manage_distribution"`
	IngressClass       string `json:"ingress_class,omitempty"`
}
