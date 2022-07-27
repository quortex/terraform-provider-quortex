package quortex

// Dataplane -
type Dataplane struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Organization       string `json:"organization,omitempty"`
	Provider           string `json:"provider,omitempty"`
	Region             string `json:"region"`
	KubeEndpoint       string `json:"kube_endpoint"`
	KubeCertificate    string `json:"kube_certificate,omitempty"`
	KubeToken          string `json:"kube_token,omitempty"`
	LiveEndpoint       string `json:"live_endpoint,omitempty"`
	RtmpEndpoint       string `json:"rtmp_endpoint,omitempty"`
	MeshEndpoint       string `json:"mesh_endpoint,omitempty"`
	GrafanaEndpoint    string `json:"grafana_endpoint,omitempty"`
	Enable             bool   `json:"enable"`
	ManageDistribution bool   `json:"manage_distribution"`
	IngressClass       string `json:"ingress_class,omitempty"`
	SmartTrafficQuery  string `json:"smart_traffic_query,omitempty"`
	CreateHpas         bool   `json:"create_hpas"`
	CdnReconciliation  bool   `json:"cdn_reconciliation"`
}
