package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttPublishingPoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePublishingPointCreate,
		ReadContext:   resourcePublishingPointRead,
		UpdateContext: resourcePublishingPointUpdate,
		DeleteContext: resourcePublishingPointDelete,
		Schema: map[string]*schema.Schema{
			"pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"published": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"input_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"processing_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"target_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"custom_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelPublishingPoint(d *schema.ResourceData) (*PublishingPoint, error) {

	ve := PublishingPoint{
		InputUuid:      d.Get("input_uuid").(string),
		ProcessingUuid: d.Get("processing_uuid").(string),
		TargetUuid:     d.Get("target_uuid").(string),
		CustomPath:     d.Get("custom_path").(string),
		TargetType:     d.Get("target_type").(string),
		Published:      d.Get("published").(bool),
	}

	return &ve, nil
}

func resourcePublishingPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	ve, err1 := marshallModelPublishingPoint(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreatePublishingPoint(poolId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourcePublishingPointRead(ctx, d, m)

	return diags
}

func resourcePublishingPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	publishingPointId := d.Id()

	publishingPoint, err := c.GetPublishingPoint(poolId, publishingPointId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("input_uuid", publishingPoint.InputUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("processing_uuid", publishingPoint.ProcessingUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("target_uuid", publishingPoint.TargetUuid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("custom_path", publishingPoint.CustomPath); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("target_type", publishingPoint.TargetType); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("published", publishingPoint.Published); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePublishingPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)
	publishingPointId := d.Id()

	ve, err1 := marshallModelPublishingPoint(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdatePublishingPoint(poolId, publishingPointId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourcePublishingPointRead(ctx, d, m)

	return diags
}

func resourcePublishingPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	publishingPointId := d.Id()

	err := c.DeletePublishingPoint(poolId, publishingPointId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
