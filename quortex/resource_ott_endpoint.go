package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttEndpoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOTTEndpointCreate,
		ReadContext:   resourceOTTEndpointRead,
		UpdateContext: resourceOTTEndpointUpdate,
		DeleteContext: resourceOTTEndpointDelete,
		Schema: map[string]*schema.Schema{
			"pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"custom_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"input_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"processing_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelOTTEndpoint(d *schema.ResourceData) (*OTTEndpoint, error) {

	ve := OTTEndpoint{
		Enabled:        d.Get("enabled").(bool),
		CustomPath:     d.Get("custom_path").(string),
		InputUuid:      d.Get("input_uuid").(string),
		ProcessingUuid: d.Get("processing_uuid").(string),
		TargetUuid:     d.Get("target_uuid").(string),
	}

	return &ve, nil
}

func resourceOTTEndpointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	ve, err1 := marshallModelOTTEndpoint(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreateOTTEndpoint(poolId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceOTTEndpointRead(ctx, d, m)

	return diags
}

func resourceOTTEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ottEndpointId := d.Id()

	ottEndpoint, err := c.GetOTTEndpoint(poolId, ottEndpointId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("input_uuid", ottEndpoint.InputUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("processing_uuid", ottEndpoint.ProcessingUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("target_uuid", ottEndpoint.TargetUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("custom_path", ottEndpoint.CustomPath); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enabled", ottEndpoint.Enabled); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceOTTEndpointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)
	ottEndpointId := d.Id()

	ve, err1 := marshallModelOTTEndpoint(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdateOTTEndpoint(poolId, ottEndpointId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceOTTEndpointRead(ctx, d, m)

	return diags
}

func resourceOTTEndpointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ottEndpointId := d.Id()

	err := c.DeleteOTTEndpoint(poolId, ottEndpointId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
