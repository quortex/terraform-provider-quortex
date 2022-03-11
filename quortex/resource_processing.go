package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProcessing() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProcessingCreate,
		ReadContext:   resourceProcessingRead,
		UpdateContext: resourceProcessingUpdate,
		DeleteContext: resourceProcessingDelete,
		Schema: map[string]*schema.Schema{
			"pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"published": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"video": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"codec": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"bitrate": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"framerate": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"resolution": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"width": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"height": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"audio": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"codec": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"channels": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"bitrate": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"samplerate": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"track": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"output": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ad": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"subtitle": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"track": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"hoh": &schema.Schema{
							Type:     schema.TypeBool,
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

	for _, video := range videos {
		vid := video.(map[string]interface{})
		vi := VideoMedia{
			Codec:     vid["codec"].(string),
			Bitrate:   vid["bitrate"].(int),
			Framerate: vid["framerate"].(string),
		}

		if resolution, ok := vid["resolution"]; ok {
			res := resolution.([]interface{})
			re := res[0].(map[string]interface{})

			reso := Resolution{
				Width:  re["width"].(int),
				Height: re["height"].(int),
			}
			vi.Resolution = &reso

		}

		ve.VideoMedias = append(ve.VideoMedias, vi)
	}

	for _, audio := range audios {
		aud := audio.(map[string]interface{})
		au := AudioMedia{
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
			oi["codec"] = video.Codec
			oi["bitrate"] = video.Bitrate
			oi["framerate"] = video.Framerate
			oi["resolution"] = flattenResolution(video.Resolution)
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
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
			oi["hoh"] = subtitle.DeafAndHardOfHearing
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}
