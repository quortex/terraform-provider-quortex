package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePoolCreate,
		ReadContext:   resourcePoolRead,
		UpdateContext: resourcePoolUpdate,
		DeleteContext: resourcePoolDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"published": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"input_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"streaming_countries": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_shifting": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"startover_duration": {
							Type:     schema.TypeInt,
							Optional: true,
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

func marshallModelPool(d *schema.ResourceData) (*Pool, error) {

	time_shiftings := d.Get("time_shifting").([]interface{})
	ve := Pool{
		Name:        d.Get("name").(string),
		Published:   d.Get("published").(bool),
		InputRegion: d.Get("input_region").(string),
		Label:       d.Get("label").(string),
	}
	countries := d.Get("streaming_countries").([]interface{})
	for _, country := range countries {
		ve.StreamingCountries = append(ve.StreamingCountries, country.(string))
	}

	for _, time_shifting := range time_shiftings {
		time_shift := time_shifting.(map[string]interface{})
		ts := TimeShifting{
			Enabled:           time_shift["enabled"].(bool),
			StartoverDuration: time_shift["startover_duration"].(int),
		}
		ve.TimeShifting = &ts
	}

	return &ve, nil
}

func resourcePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	ve, err1 := marshallModelPool(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.CreatePool(*ve)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Uuid)

	resourcePoolRead(ctx, d, m)

	return diags
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

	if err := d.Set("name", pool.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("input_region", pool.InputRegion); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("published", pool.Published); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("streaming_countries", pool.StreamingCountries); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("label", pool.Label); err != nil {
		return diag.FromErr(err)
	}

	timeshift := flattenPoolTimeShifting(pool.TimeShifting)
	if err := d.Set("time_shifting", timeshift); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Id()

	ve, err1 := marshallModelPool(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.UpdatePool(poolId, *ve)
	if err != nil {
		return diag.FromErr(err)

	}

	d.SetId(o.Uuid)

	resourcePoolRead(ctx, d, m)

	return diags
}

func resourcePoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeletePool(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenPoolTimeShifting(timeshifting *TimeShifting) []interface{} {

	c := make(map[string]interface{})
	if timeshifting != nil {
		c["enabled"] = timeshifting.Enabled
		c["startover_duration"] = timeshifting.StartoverDuration

	}
	return []interface{}{c}
}
