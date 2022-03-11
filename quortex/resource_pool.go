package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePoolCreate,
		ReadContext:   resourcePoolRead,
		UpdateContext: resourcePoolUpdate,
		DeleteContext: resourcePoolDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"published": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"input_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"streaming_countries": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	ve := Pool{
		Name:        d.Get("name").(string),
		Published:   d.Get("published").(bool),
		InputRegion: d.Get("input_region").(string),
	}
	countries := d.Get("streaming_countries").([]interface{})
	for _, country := range countries {
		ve.StreamingCountries = append(ve.StreamingCountries, country.(string))
	}

	o, err := c.CreatePool(ve)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Uuid)

	return resourcePoolRead(ctx, d, m)
}

func resourcePoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	poolId := d.Id()

	pool, err := c.GetPool(poolId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", pool.Name)
	d.Set("input_region", pool.InputRegion)
	d.Set("published", pool.Published)

	return diags
}

func resourcePoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	poolId := d.Id()

	ve := Pool{
		Name:        d.Get("name").(string),
		Published:   d.Get("published").(bool),
		InputRegion: d.Get("input_region").(string),
	}
	countries := d.Get("streaming_countries").([]interface{})
	for _, country := range countries {
		ve.StreamingCountries = append(ve.StreamingCountries, country.(string))
	}

	_, err := c.UpdatePool(poolId, ve)
	if err != nil {
		return diag.FromErr(err)

	}
	return resourcePoolRead(ctx, d, m)
}

func resourcePoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	poolId := d.Id()

	err := c.DeletePool(poolId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
