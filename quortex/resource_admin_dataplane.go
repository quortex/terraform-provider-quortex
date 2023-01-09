package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminDataplane() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDataplaneCreate,
		ReadContext:   resourceDataplaneRead,
		UpdateContext: resourceDataplaneUpdate,
		DeleteContext: resourceDataplaneDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"organization": {
				ForceNew: true,
				Optional: true,
				Type:     schema.TypeString,
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
			},

			"kube_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},

			"kube_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"kube_token": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"live_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"rtmp_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"grafana_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mesh_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"enable": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"manage_distribution": {
				Type:     schema.TypeBool,
				Required: true,
			},

			"ingress_class": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "traefik",
			},

			"smart_traffic_query": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "smart_traffic",
			},

			"create_hpas": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"cdn_reconciliation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceDataplaneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	ve := Dataplane{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		Organization:       d.Get("organization").(string),
		Region:             d.Get("region").(string),
		KubeEndpoint:       d.Get("kube_endpoint").(string),
		KubeCertificate:    d.Get("kube_certificate").(string),
		KubeToken:          d.Get("kube_token").(string),
		LiveEndpoint:       d.Get("live_endpoint").(string),
		RtmpEndpoint:       d.Get("rtmp_endpoint").(string),
		MeshEndpoint:       d.Get("mesh_endpoint").(string),
		GrafanaEndpoint:    d.Get("grafana_endpoint").(string),
		Enable:             d.Get("enable").(bool),
		ManageDistribution: d.Get("manage_distribution").(bool),
		IngressClass:       d.Get("ingress_class").(string),
		SmartTrafficQuery:  d.Get("smart_traffic_query").(string),
		CreateHpas:         d.Get("create_hpas").(bool),
		CdnReconciliation:  d.Get("cdn_reconciliation").(bool),
	}

	o, err := c.CreateDataplane(ve)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Name)

	return resourceDataplaneRead(ctx, d, m)
}

func resourceDataplaneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	dataplaneId := d.Id()

	dataplane, err := c.GetDataplane(dataplaneId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", dataplane.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", dataplane.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("organization", dataplane.Organization); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("region", dataplane.Region); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("kube_endpoint", dataplane.KubeEndpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("kube_certificate", dataplane.KubeCertificate); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("kube_token", dataplane.KubeToken); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("live_endpoint", dataplane.LiveEndpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("rtmp_endpoint", dataplane.RtmpEndpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("mesh_endpoint", dataplane.MeshEndpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("grafana_endpoint", dataplane.GrafanaEndpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enable", dataplane.Enable); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("manage_distribution", dataplane.ManageDistribution); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("ingress_class", dataplane.IngressClass); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("smart_traffic_query", dataplane.SmartTrafficQuery); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("create_hpas", dataplane.CreateHpas); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cdn_reconciliation", dataplane.CdnReconciliation); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDataplaneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	dataplaneId := d.Id()

	ve := Dataplane{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		Organization:       d.Get("organization").(string),
		Region:             d.Get("region").(string),
		KubeEndpoint:       d.Get("kube_endpoint").(string),
		KubeCertificate:    d.Get("kube_certificate").(string),
		KubeToken:          d.Get("kube_token").(string),
		LiveEndpoint:       d.Get("live_endpoint").(string),
		RtmpEndpoint:       d.Get("rtmp_endpoint").(string),
		MeshEndpoint:       d.Get("mesh_endpoint").(string),
		GrafanaEndpoint:    d.Get("grafana_endpoint").(string),
		Enable:             d.Get("enable").(bool),
		ManageDistribution: d.Get("manage_distribution").(bool),
		IngressClass:       d.Get("ingress_class").(string),
		SmartTrafficQuery:  d.Get("smart_traffic_query").(string),
		CreateHpas:         d.Get("create_hpas").(bool),
		CdnReconciliation:  d.Get("cdn_reconciliation").(bool),
	}

	_, err := c.UpdateDataplane(dataplaneId, ve)
	if err != nil {
		return diag.FromErr(err)

	}
	return resourceDataplaneRead(ctx, d, m)
}

func resourceDataplaneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteDataplane(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
