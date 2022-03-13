package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTargetCreate,
		ReadContext:   resourceTargetRead,
		UpdateContext: resourceTargetUpdate,
		DeleteContext: resourceTargetDelete,
		Schema: map[string]*schema.Schema{
			"pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"published": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"segment_duration": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"playlist_length": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"container": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scte_35": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"filter_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelTarget(d *schema.ResourceData) (*Target, error) {

	scte35s := d.Get("scte_35").([]interface{})
	ve := Target{
		Name:            d.Get("name").(string),
		Identifier:      d.Get("identifier").(string),
		Published:       d.Get("published").(bool),
		Type:            d.Get("type").(string),
		SegmentDuration: d.Get("segment_duration").(float64),
		PlaylistLength:  d.Get("playlist_length").(int),
		Container:       d.Get("container").(string),
	}

	for _, scte35 := range scte35s {
		scte := scte35.(map[string]interface{})
		sc := Scte35{
			Enabled:    scte["enabled"].(bool),
			FilterType: scte["filter_type"].(string),
		}
		filters := scte["filter_list"].([]interface{})
		for _, filter := range filters {
			sc.FilterList = append(sc.FilterList, filter.(string))
		}
		ve.Scte35 = &sc
	}

	return &ve, nil
}

func resourceTargetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	ve, err1 := marshallModelTarget(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreateTarget(poolId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceTargetRead(ctx, d, m)

	return diags
}

func resourceTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	targetId := d.Id()

	target, err := c.GetTarget(poolId, targetId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", target.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("identifier", target.Identifier); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("published", target.Published); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", target.Type); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("segment_duration", target.SegmentDuration); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("playlist_length", target.PlaylistLength); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("container", target.Container); err != nil {
		return diag.FromErr(err)
	}

	scte35s := flattenTargetScte35(target.Scte35)
	if err := d.Set("scte_35", scte35s); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTargetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)
	targetId := d.Id()

	ve, err1 := marshallModelTarget(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdateTarget(poolId, targetId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceTargetRead(ctx, d, m)

	return diags
}

func resourceTargetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	targetId := d.Id()

	err := c.DeleteTarget(poolId, targetId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenTargetScte35(scte35 *Scte35) []interface{} {

	c := make(map[string]interface{})
	if scte35 != nil {
		c["enabled"] = scte35.Enabled
		c["filter_type"] = scte35.FilterType
		c["filter_list"] = scte35.FilterList
	}
	return []interface{}{c}
}
