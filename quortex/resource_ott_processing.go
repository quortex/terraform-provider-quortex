package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttProcessing() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProcessingCreate,
		ReadContext:   resourceProcessingRead,
		UpdateContext: resourceProcessingUpdate,
		DeleteContext: resourceProcessingDelete,
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
			"video": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"codec": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bitrate": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"framerate": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resolution": {
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"width": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"height": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
						"advanced": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"profile": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  nil,
									},
									"level": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  nil,
									},
									"quality": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  nil,
									},
									"encoding_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  nil,
									},
									"encoding_quality": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"quality_optimization": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  nil,
									},
									"closed_gop": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  nil,
									},
									"gop_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"gop_max_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"bframe": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  nil,
									},
									"bframe_number": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"maxrate": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"key_frame_interval": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"horizontal_sharpness": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"vertical_sharpness": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  nil,
									},
									"logo_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  nil,
									},
								},
							},
						},
					},
				},
			},
			"audio": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"codec": {
							Type:     schema.TypeString,
							Required: true,
						},
						"channels": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bitrate": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"samplerate": {
							Type:     schema.TypeString,
							Required: true,
						},
						"track": {
							Type:     schema.TypeString,
							Required: true,
						},
						"output": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ad": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"subtitle": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"track": {
							Type:     schema.TypeString,
							Required: true,
						},
						"output": {
							Type:     schema.TypeString,
							Required: true,
						},
						"hoh": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
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

func marshallModelProcessing(d *schema.ResourceData) (*Processing, error) {

	videos := d.Get("video").([]interface{})
	audios := d.Get("audio").([]interface{})
	subtitles := d.Get("subtitle").([]interface{})

	ve := Processing{
		Name:           d.Get("name").(string),
		Identifier:     d.Get("identifier").(string),
		Published:      d.Get("published").(bool),
		VideoMedias:    []VideoMedia{},
		AudioMedias:    []AudioMedia{},
		SubtitleMedias: []SubtitleMedia{},
	}

	labels := d.Get("labels").([]interface{})
	for _, label := range labels {
		ve.Labels = append(ve.Labels, label.(string))
	}

	for _, video := range videos {
		vid := video.(map[string]interface{})
		vi := VideoMedia{
			Codec:     vid["codec"].(string),
			Label:     vid["label"].(string),
			Bitrate:   vid["bitrate"].(int),
			Framerate: vid["framerate"].(string),
		}

		if resolution, ok := vid["resolution"]; ok {
			reso := resolution.([]interface{})
			res := reso[0].(map[string]interface{})
			re := Resolution{
				Width:  res["width"].(int),
				Height: res["height"].(int),
			}
			vi.Resolution = &re
		}

		if advanced, ok := vid["advanced"]; ok {
			ad := Advanced{}

			if advanced != nil {
				adva := advanced.([]interface{})
				if len(adva) > 0 {
					first := adva[0]
					if first != nil {
						adv := first.(map[string]interface{})
						if val, ok := adv["bframe"]; ok {
							bframe := new(bool)
							*bframe = val.(bool)
							ad.Bframe = bframe
						}
						if val, ok := adv["closed_gop"]; ok {
							closedgop := new(bool)
							*closedgop = val.(bool)
							ad.ClosedGop = closedgop
						}
						ad.Profile = adv["profile"].(string)
						ad.Level = adv["level"].(string)
						ad.Quality = adv["quality"].(string)
						ad.EncodingMode = adv["encoding_mode"].(string)
						ad.EncodingQuality = adv["encoding_quality"].(int)
						ad.QualityOptimization = adv["quality_optimization"].(string)
						ad.GopSize = adv["gop_size"].(int)
						ad.GopMaxSize = adv["gop_max_size"].(int)
						ad.BframeNumber = adv["bframe_number"].(int)
						ad.Maxrate = adv["maxrate"].(int)
						ad.KeyFrameIntervalMs = adv["key_frame_interval"].(int)
						ad.HorizontalSharpness = adv["horizontal_sharpness"].(int)
						ad.VerticalSharpness = adv["vertical_sharpness"].(int)
						ad.LogoEnabled = adv["logo_enabled"].(bool)
					}
				}
			}
			vi.Advanced = &ad
		}
		ve.VideoMedias = append(ve.VideoMedias, vi)
	}

	for _, audio := range audios {
		aud := audio.(map[string]interface{})
		au := AudioMedia{
			Label:            aud["label"].(string),
			Codec:            aud["codec"].(string),
			Channels:         aud["channels"].(string),
			Bitrate:          aud["bitrate"].(int),
			Samplerate:       aud["samplerate"].(string),
			Track:            aud["track"].(string),
			Output:           aud["output"].(string),
			AudioDescription: aud["ad"].(bool),
		}
		ve.AudioMedias = append(ve.AudioMedias, au)
	}

	for _, subtitle := range subtitles {
		sub := subtitle.(map[string]interface{})
		su := SubtitleMedia{
			Track:                sub["track"].(string),
			Output:               sub["output"].(string),
			DeafAndHardOfHearing: sub["hoh"].(bool),
		}
		ve.SubtitleMedias = append(ve.SubtitleMedias, su)
	}

	return &ve, nil
}

func resourceProcessingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	ve, err1 := marshallModelProcessing(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreateProcessing(poolId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceProcessingRead(ctx, d, m)

	return diags
}

func resourceProcessingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	processingId := d.Id()

	processing, err := c.GetProcessing(poolId, processingId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", processing.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("identifier", processing.Identifier); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("published", processing.Published); err != nil {
		return diag.FromErr(err)
	}

	videos := flattenProcessingVideos(&processing.VideoMedias)
	if err := d.Set("video", videos); err != nil {
		return diag.FromErr(err)
	}

	audios := flattenProcessingAudios(&processing.AudioMedias)
	if err := d.Set("audio", audios); err != nil {
		return diag.FromErr(err)
	}

	subtitles := flattenProcessingSubtitles(&processing.SubtitleMedias)
	if err := d.Set("subtitle", subtitles); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("labels", processing.Labels); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProcessingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)
	processingId := d.Id()

	ve, err1 := marshallModelProcessing(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdateProcessing(poolId, processingId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceProcessingRead(ctx, d, m)

	return diags
}

func resourceProcessingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	processingId := d.Id()

	err := c.DeleteProcessing(poolId, processingId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenProcessingVideos(videos *[]VideoMedia) []interface{} {
	if videos != nil {
		ois := make([]interface{}, len(*videos))

		for i, video := range *videos {
			oi := make(map[string]interface{})
			oi["label"] = video.Label
			oi["codec"] = video.Codec
			oi["bitrate"] = video.Bitrate
			oi["framerate"] = video.Framerate
			oi["resolution"] = flattenResolution(video.Resolution)
			oi["advanced"] = flattenAdvanced(video.Advanced)
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenAdvanced(advanced *Advanced) []interface{} {
	c := make(map[string]interface{})
	c["profile"] = (*advanced).Profile
	c["level"] = (*advanced).Level
	c["quality"] = (*advanced).Quality
	c["encoding_mode"] = (*advanced).EncodingMode
	c["encoding_quality"] = (*advanced).EncodingQuality
	c["quality_optimization"] = (*advanced).QualityOptimization
	c["closed_gop"] = (*advanced).ClosedGop
	c["gop_size"] = (*advanced).GopSize
	c["gop_max_size"] = (*advanced).GopMaxSize
	c["bframe"] = (*advanced).Bframe
	c["bframe_number"] = (*advanced).BframeNumber
	c["maxrate"] = (*advanced).Maxrate
	c["key_frame_interval"] = (*advanced).KeyFrameIntervalMs
	c["horizontal_sharpness"] = (*advanced).HorizontalSharpness
	c["vertical_sharpness"] = (*advanced).VerticalSharpness
	c["logo_enabled"] = (*advanced).LogoEnabled

	return []interface{}{c}
}

func flattenResolution(resolution *Resolution) []interface{} {
	c := make(map[string]interface{})
	c["width"] = (*resolution).Width
	c["height"] = (*resolution).Height

	return []interface{}{c}
}

func flattenProcessingAudios(audios *[]AudioMedia) []interface{} {
	if audios != nil {
		ois := make([]interface{}, len(*audios))

		for i, audio := range *audios {
			oi := make(map[string]interface{})
			oi["label"] = audio.Label
			oi["codec"] = audio.Codec
			oi["channels"] = audio.Channels
			oi["bitrate"] = audio.Bitrate
			oi["samplerate"] = audio.Samplerate
			oi["track"] = audio.Track
			oi["output"] = audio.Output
			oi["ad"] = audio.AudioDescription
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenProcessingSubtitles(subtitles *[]SubtitleMedia) []interface{} {
	if subtitles != nil {
		ois := make([]interface{}, len(*subtitles))

		for i, subtitle := range *subtitles {
			oi := make(map[string]interface{})
			oi["track"] = subtitle.Track
			oi["output"] = subtitle.Output
			oi["hoh"] = subtitle.DeafAndHardOfHearing
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}
