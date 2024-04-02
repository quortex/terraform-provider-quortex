package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOttTarget() *schema.Resource {
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
				Default:  "",
			},
			"published": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
				Default:  "",
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
							Default:  false,
						},
						"filter_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "deny",
						},
						"filter_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"encryption_dynamic": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"drm_merchant_uuid": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"encryption": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 5,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"labels": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"stream_type": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"hls_advanced": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"program_datetime": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"version": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"dash_advanced": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"utc_timing": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"utc_timing_server": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"base_url": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"position": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"suggested_presentation_delay": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"start_time_origin_offset": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"input_label_restriction": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"processing_label_restriction": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"latency": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "standard",
				ValidateFunc: validation.StringInSlice([]string{"standard", "low"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelTarget(d *schema.ResourceData) (*Target, error) {

	ve := Target{
		Name:            d.Get("name").(string),
		Identifier:      d.Get("identifier").(string),
		Published:       d.Get("published").(bool),
		Type:            d.Get("type").(string),
		SegmentDuration: d.Get("segment_duration").(float64),
		PlaylistLength:  d.Get("playlist_length").(int),
		Container:       d.Get("container").(string),
		Latency:         d.Get("latency").(string),
	}
	ilrs := d.Get("input_label_restriction").([]interface{})
	for _, ilr := range ilrs {
		ve.InputLabelRestriction = append(ve.InputLabelRestriction, ilr.(string))
	}
	plrs := d.Get("processing_label_restriction").([]interface{})
	for _, plr := range plrs {
		ve.ProcessingLabelRestriction = append(ve.ProcessingLabelRestriction, plr.(string))
	}

	scte35s := d.Get("scte_35").([]interface{})
	for _, scte35 := range scte35s {
		scte := scte35.(map[string]interface{})
		sc := Scte35{
			Enabled:    scte["enabled"].(bool),
			FilterType: scte["filter_type"].(string),
		}
		filters := scte["filter_list"].([]interface{})
		for _, filter := range filters {
			if filter != nil {
				sc.FilterList = append(sc.FilterList, filter.(string))
			}
		}
		ve.Scte35 = &sc
	}

	encdyns := d.Get("encryption_dynamic").([]interface{})
	for _, encdyn := range encdyns {
		if encdyn != nil {
			encdy := encdyn.(map[string]interface{})
			ed := EncryptionDynamic{
				ContentId:       encdy["content_id"].(string),
				DrmMerchantUuid: encdy["drm_merchant_uuid"].(string),
			}

			// Manage encryption
			encrs := encdy["encryption"].([]interface{})
			for _, encr := range encrs {
				enc := encr.(map[string]interface{})
				en := Encryption{
					Uuid:       enc["uuid"].(string),
					StreamType: enc["stream_type"].(string),
				}

				labels := enc["labels"].([]interface{})
				for _, label := range labels {
					en.Labels = append(en.Labels, label.(string))
				}

				ed.Encryption = append(ed.Encryption, en)
			}
			if ed.ContentId != "" {
				ve.EncryptionDynamic = &ed
				ve.EncryptionType = "dynamic"
			}
		}
	}

	hlsadvs := d.Get("hls_advanced").([]interface{})
	for _, hlsadv := range hlsadvs {
		hlsadv := hlsadv.(map[string]interface{})
		hlsa := HlsAdvanced{
			ProgramDatetime: hlsadv["program_datetime"].(string),
			Version:         hlsadv["version"].(int),
		}
		ve.HlsAdvanced = &hlsa
	}

	dashadvs := d.Get("dash_advanced").([]interface{})
	for _, dashadv := range dashadvs {
		dashadv := dashadv.(map[string]interface{})
		dasha := DashAdvanced{
			UtcTiming:                  dashadv["utc_timing"].(string),
			UtcTimingServer:            dashadv["utc_timing_server"].(string),
			BaseUrl:                    dashadv["base_url"].(string),
			Position:                   dashadv["position"].(string),
			SuggestedPresentationDelay: dashadv["suggested_presentation_delay"].(int),
			StartTimeOriginOffset:      dashadv["start_time_origin_offset"].(int),
		}
		ve.DashAdvanced = &dasha
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

	if err := d.Set("input_label_restriction", target.InputLabelRestriction); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("processing_label_restriction", target.ProcessingLabelRestriction); err != nil {
		return diag.FromErr(err)
	}

	encdyns := flattenTargetEncryptionDynamic(target.EncryptionDynamic)
	if err := d.Set("encryption_dynamic", encdyns); err != nil {
		return diag.FromErr(err)
	}

	hlsadv := flattenTargetHlsAdvanced(target.HlsAdvanced)
	if err := d.Set("hls_advanced", hlsadv); err != nil {
		return diag.FromErr(err)
	}

	dashadv := flattenTargetDashAdvanced(target.DashAdvanced)
	if err := d.Set("dash_advanced", dashadv); err != nil {
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

func flattenTargetEncryptionDynamic(enc *EncryptionDynamic) []interface{} {

	c := make(map[string]interface{})
	if enc != nil {
		c["content_id"] = enc.ContentId
		c["drm_merchant_uuid"] = enc.DrmMerchantUuid
		c["encryption"] = flattenTargetEncryption(&enc.Encryption)
	}
	return []interface{}{c}
}

func flattenTargetEncryption(enc *[]Encryption) []interface{} {
	if enc != nil {
		ois := make([]interface{}, len(*enc))

		for i, en := range *enc {
			oi := make(map[string]interface{})
			oi["uuid"] = en.Uuid
			oi["stream_type"] = en.StreamType
			oi["labels"] = en.Labels
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenTargetDashAdvanced(adv *DashAdvanced) []interface{} {

	c := make(map[string]interface{})
	if adv != nil {
		c["utc_timing"] = adv.UtcTiming
		c["utc_timing_server"] = adv.UtcTimingServer
		c["base_url"] = adv.BaseUrl
		c["position"] = adv.Position
		c["suggested_presentation_delay"] = adv.SuggestedPresentationDelay
		c["start_time_origin_offset"] = adv.StartTimeOriginOffset
	}
	return []interface{}{c}
}

func flattenTargetHlsAdvanced(adv *HlsAdvanced) []interface{} {

	c := make(map[string]interface{})
	if adv != nil {
		c["program_datetime"] = adv.ProgramDatetime
		c["version"] = adv.Version
	}
	return []interface{}{c}
}
