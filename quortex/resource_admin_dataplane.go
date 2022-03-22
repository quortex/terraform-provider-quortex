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
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"organization": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"cloud_vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "aws",
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},

			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"livepoint": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"rtmpendpoint": {
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
		Provider:           d.Get("cloud_vendor").(string),
		Region:             d.Get("region").(string),
		Endpoint:           d.Get("endpoint").(string),
		Certificate:        d.Get("certificate").(string),
		Token:              d.Get("token").(string),
		Livepoint:          d.Get("livepoint").(string),
		Rtmpendpoint:       d.Get("rtmpendpoint").(string),
		Enable:             d.Get("enable").(bool),
		ManageDistribution: d.Get("manage_distribution").(bool),
		IngressClass:       d.Get("ingress_class").(string),
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

	if err := d.Set("cloud_vendor", dataplane.Provider); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("region", dataplane.Region); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("endpoint", dataplane.Endpoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("certificate", dataplane.Certificate); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("token", dataplane.Token); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("livepoint", dataplane.Livepoint); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("rtmpendpoint", dataplane.Rtmpendpoint); err != nil {
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

	return diags
}

func resourceDataplaneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	dataplaneId := d.Id()

	ve := Dataplane{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		Organization:       d.Get("organization").(string),
		Provider:           d.Get("cloud_vendor").(string),
		Region:             d.Get("region").(string),
		Endpoint:           d.Get("endpoint").(string),
		Certificate:        d.Get("certificate").(string),
		Token:              d.Get("token").(string),
		Livepoint:          d.Get("livepoint").(string),
		Rtmpendpoint:       d.Get("rtmpendpoint").(string),
		Enable:             d.Get("enable").(bool),
		ManageDistribution: d.Get("manage_distribution").(bool),
		IngressClass:       d.Get("ingress_class").(string),
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
