package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"path_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
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
				Computed: true,
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
			"catchup": {
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
							Default:  false,
						},
						"bucket": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"region": {
										Type:     schema.TypeString,
										Required: true,
									},
									"s3": {
										Type:     schema.TypeList,
										Optional: true,
										MinItems: 0,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"access_key": {
													Type:     schema.TypeString,
													Required: true,
												},
												"secret_key": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"origin": {
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
							Default:  false,
						},
						"whitelist_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"whitelist_cidr": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"authorization_header_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"authorization_header_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"processing_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "standard",
				ValidateFunc: validation.StringInSlice([]string{"standard", "advanced", "warm_disaster_recovery"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelPool(d *schema.ResourceData) (*Pool, error) {

	time_shiftings := d.Get("time_shifting").([]interface{})
	catchups := d.Get("catchup").([]interface{})
	origins := d.Get("origin").([]interface{})
	ve := Pool{
		Name:           d.Get("name").(string),
		Published:      d.Get("published").(bool),
		PathPrefix:     d.Get("path_prefix").(string),
		InputRegion:    d.Get("input_region").(string),
		Label:          d.Get("label").(string),
		ProcessingType: d.Get("processing_type").(string),
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

	for _, catchup := range catchups {
		catch := catchup.(map[string]interface{})
		ca := Catchup{
			Enabled: catch["enabled"].(bool),
		}

		buckets := catch["bucket"].([]interface{})
		for _, bucket := range buckets {
			buck := bucket.(map[string]interface{})
			buc := Bucket2{
				Name:   buck["name"].(string),
				Region: buck["region"].(string),
			}

			if bucs3, ok := buck["s3"]; ok {
				s3s := bucs3.([]interface{})
				if len(s3s) > 0 {
					buc.Type = "s3"
					s3 := s3s[0].(map[string]interface{})
					so := S3{
						AccessKey: s3["access_key"].(string),
						SecretKey: s3["secret_key"].(string),
					}
					buc.S3 = &so
				}
			}
			ca.Bucket2 = &buc
		}
		ve.Catchup = &ca
	}

	for _, origin := range origins {
		origi := origin.(map[string]interface{})
		ori := Origin{
			Enabled:                    origi["enabled"].(bool),
			WhitelistEnabled:           origi["whitelist_enabled"].(bool),
			AuthorizationHeaderEnabled: origi["authorization_header_enabled"].(bool),
			AuthorizationHeaderValue:   origi["authorization_header_value"].(string),
		}
		ori.WhitelistCidr = []string{}
		cidrs := origi["whitelist_cidr"].([]interface{})
		for _, cidr := range cidrs {
			ori.WhitelistCidr = append(ori.WhitelistCidr, cidr.(string))
		}

		ve.Origin = &ori
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

	if err := d.Set("path_prefix", pool.PathPrefix); err != nil {
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

	catchup := flattenPoolCatchup(pool.Catchup)
	if err := d.Set("catchup", catchup); err != nil {
		return diag.FromErr(err)
	}

	origin := flattenPoolOrigin(pool.Origin)
	if err := d.Set("origin", origin); err != nil {
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

func flattenPoolCatchup(catchup *Catchup) []interface{} {

	c := make(map[string]interface{})
	if catchup != nil {
		c["enabled"] = catchup.Enabled
		if catchup.Bucket2 != nil {
			c["bucket"] = flattenPoolCatchupBucket(catchup.Bucket2)
		}

	}
	return []interface{}{c}
}

func flattenPoolCatchupBucket(bucket *Bucket2) []interface{} {

	c := make(map[string]interface{})
	if bucket != nil {
		c["name"] = bucket.Name
		c["region"] = bucket.Region
		if bucket.S3 != nil {
			c["s3"] = flattenPoolCatchupBucketS3(bucket.S3)
		}

	}
	return []interface{}{c}
}

func flattenPoolCatchupBucketS3(s3 *S3) []interface{} {

	c := make(map[string]interface{})
	if s3 != nil {
		c["access_key"] = s3.AccessKey
		c["secret_key"] = s3.SecretKey
	}
	return []interface{}{c}
}

func flattenPoolOrigin(origin *Origin) []interface{} {

	c := make(map[string]interface{})
	if origin != nil {
		c["enabled"] = origin.Enabled
		c["whitelist_enabled"] = origin.WhitelistEnabled
		c["whitelist_cidr"] = origin.WhitelistCidr
		c["authorization_header_enabled"] = origin.AuthorizationHeaderEnabled
		c["authorization_header_value"] = origin.AuthorizationHeaderValue
	}
	return []interface{}{c}
}
