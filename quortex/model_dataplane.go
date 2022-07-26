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
	LiveEndpoint       string `json:"livepoint,omitempty"`
	RtmpEndpoint       string `json:"rtmpendpoint,omitempty"`
	MeshEndpoint       string `json:"meshendpoint,omitempty"`
	GrafanaEndpoint    string `json:"grafanaendpoint,omitempty"`
	Enable             bool   `json:"enable"`
	ManageDistribution bool   `json:"manage_distribution"`
	IngressClass       string `json:"ingress_class,omitempty"`
	SmartTrafficQuery  string `json:"smart_traffic_query,omitempty"`
	CreateHpas         bool   `json:"create_hpas"`
	CdnReconciliation  bool   `json:"cdn_reconciliation"`
}
