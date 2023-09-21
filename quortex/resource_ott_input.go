package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttInput() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInputCreate,
		ReadContext:   resourceInputRead,
		UpdateContext: resourceInputUpdate,
		DeleteContext: resourceInputDelete,
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
			"stream": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"fallback_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"srt": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"latency": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1000,
									},
									"listener": {
										Type:     schema.TypeList,
										Optional: true,
										MinItems: 0,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cidr": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"passphrase": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "",
												},
											},
										},
									},
									"caller": {
										Type:     schema.TypeList,
										Optional: true,
										MinItems: 0,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:     schema.TypeString,
													Required: true,
												},
												"passphrase": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "",
												},
											},
										},
									},
									"overrides": {
										Type:     schema.TypeList,
										Optional: true,
										MinItems: 0,
										MaxItems: 10,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pid": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"audio": {
													Type:     schema.TypeList,
													Optional: true,
													MinItems: 0,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"language": {
																Type:     schema.TypeString,
																Optional: true,
																Default:  "",
															},
															"ad": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
															},
														},
													},
												},
												"teletext": {
													Type:     schema.TypeList,
													Optional: true,
													MinItems: 0,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"page": {
																Type:     schema.TypeString,
																Required: true,
															},
															"language": {
																Type:     schema.TypeString,
																Optional: true,
																Default:  "",
															},
															"sdh": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
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
						"rtmp": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"overrides": {
										Type:     schema.TypeList,
										Optional: true,
										MinItems: 0,
										MaxItems: 10,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"audio": {
													Type:     schema.TypeList,
													Optional: true,
													MinItems: 0,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"language": {
																Type:     schema.TypeString,
																Optional: true,
																Default:  "",
															},
															"ad": {
																Type:     schema.TypeString,
																Optional: true,
																Default:  false,
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

func marshallModelInput(d *schema.ResourceData) (*Input, error) {

	streams := d.Get("stream").([]interface{})

	ve := Input{
		Name:       d.Get("name").(string),
		Identifier: d.Get("identifier").(string),
		Published:  d.Get("published").(bool),
		Streams:    []Stream{},
	}

	labels := d.Get("labels").([]interface{})
	for _, label := range labels {
		ve.Labels = append(ve.Labels, label.(string))
	}

	for _, stream := range streams {
		st := stream.(map[string]interface{})

		str := Stream{
			Uuid:        st["uuid"].(string),
			Name:        st["name"].(string),
			Enabled:     st["enabled"].(bool),
			FallbackUrl: st["fallback_url"].(string),
		}

		// Manage srt
		if srt, ok := st["srt"]; ok {
			srts := srt.([]interface{})
			if len(srts) > 0 {
				str.Type = "srt"
				sr := srts[0].(map[string]interface{})

				srtt := Srt{
					Latency: sr["latency"].(int),
				}

				// Manage srt listener
				if srta, ok := sr["listener"]; ok {
					srtas := srta.([]interface{})
					if len(srtas) > 0 {
						srtt.ConnectionType = "listener"
						listener := Listener{}
						first := srtas[0]
						if first != nil {
							srt := first.(map[string]interface{})
							cidrs := srt["cidr"].([]interface{})
							for _, cidr := range cidrs {
								listener.Cidr = append(listener.Cidr, cidr.(string))
							}
							listener.Passphrase = srt["passphrase"].(string)
						}
						srtt.Listener = &listener
					}

				}

				// Manage srt caller
				if srta, ok := sr["caller"]; ok {
					srtas := srta.([]interface{})
					if len(srtas) > 0 {
						srtt.ConnectionType = "caller"
						caller := Caller{}
						first := srtas[0]
						if first != nil {
							srt := first.(map[string]interface{})
							caller.Address = srt["address"].(string)
							caller.Passphrase = srt["passphrase"].(string)
						}
						srtt.Caller = &caller
					}
				}

				// Manage overrides
				overrides := sr["overrides"].([]interface{})
				for _, override := range overrides {
					over := override.(map[string]interface{})
					ov := Override{
						Pid:     over["pid"].(int),
						Type:    over["type"].(string),
						Enabled: over["enabled"].(bool),
					}
					srtt.Overrides = append(srtt.Overrides, ov)
				}
				str.Srt = &srtt
			}
		}

		// Manage rtmp
		if rtmp, ok := st["rtmp"]; ok {
			rtmps := rtmp.([]interface{})
			if len(rtmps) > 0 {
				str.Type = "rtmp"
				rt := rtmps[0].(map[string]interface{})

				rtt := Rtmp{}

				// Manage overrides
				overrides := rt["overrides"].([]interface{})
				for _, override := range overrides {
					over := override.(map[string]interface{})
					ov := Override{
						Pid:     0,
						Type:    over["type"].(string),
						Enabled: over["enabled"].(bool),
					}
					rtt.Overrides = append(rtt.Overrides, ov)
				}
				str.Rtmp = &rtt
			}
		}
		ve.Streams = append(ve.Streams, str)
	}

	return &ve, nil
}

func resourceInputCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	ve, err1 := marshallModelInput(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreateInput(poolId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceInputRead(ctx, d, m)

	return diags
}

func resourceInputRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	inputId := d.Id()

	input, err := c.GetInput(poolId, inputId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", input.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("identifier", input.Identifier); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("published", input.Published); err != nil {
		return diag.FromErr(err)
	}

	streams := flattenInputStreams(&input.Streams)
	if err := d.Set("stream", streams); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("labels", input.Labels); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceInputUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	poolId := d.Get("pool_id").(string)
	inputId := d.Id()

	ve, err1 := marshallModelInput(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdateInput(poolId, inputId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceInputRead(ctx, d, m)

	return diags
}

func resourceInputDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	poolId := d.Get("pool_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	inputId := d.Id()

	err := c.DeleteInput(poolId, inputId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenInputStreams(streams *[]Stream) []interface{} {
	if streams != nil {
		ois := make([]interface{}, len(*streams))
		for i, stream := range *streams {
			oi := make(map[string]interface{})
			oi["uuid"] = stream.Uuid
			oi["name"] = stream.Name
			oi["enabled"] = stream.Enabled
			oi["fallback_url"] = stream.FallbackUrl
			if stream.Type == "srt" {
				oi["srt"] = flattenSrt(stream.Srt)
			} else if stream.Type == "rtmp" {
				oi["rtmp"] = flattenRtmp(stream.Rtmp)
			}
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenSrt(srt *Srt) []interface{} {
	c := make(map[string]interface{})
	c["latency"] = srt.Latency
	if srt.ConnectionType == "listener" {
		c["listener"] = flattenListener(srt.Listener)
	} else if srt.ConnectionType == "caller" {
		c["caller"] = flattenCaller(srt.Caller)
	}
	c["overrides"] = flattenOverrides(&srt.Overrides)
	return []interface{}{c}
}

func flattenListener(listener *Listener) []interface{} {
	c := make(map[string]interface{})
	c["cidr"] = listener.Cidr
	c["passphrase"] = listener.Passphrase
	return []interface{}{c}
}

func flattenCaller(caller *Caller) []interface{} {
	c := make(map[string]interface{})
	c["address"] = caller.Address
	c["passphrase"] = caller.Passphrase
	return []interface{}{c}
}

func flattenRtmp(rtmp *Rtmp) []interface{} {
	c := make(map[string]interface{})
	c["overrides"] = flattenOverrides(&rtmp.Overrides)
	return []interface{}{c}
}

func flattenOverrides(overrides *[]Override) []interface{} {
	if overrides != nil {
		ois := make([]interface{}, len(*overrides))
		for i, override := range *overrides {
			oi := make(map[string]interface{})
			oi["type"] = override.Type
			oi["enabled"] = override.Enabled
			if override.Pid != 0 {
				oi["pid"] = override.Pid
			}

			if override.OverAudio != nil {
				oi["audio"] = flattenOverAudio(override.OverAudio)
			}

			if override.OverTeletext != nil {
				oi["teletext"] = flattenOverTeletext(override.OverTeletext)
			}

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenOverAudio(overaudio *OverAudio) []interface{} {
	c := make(map[string]interface{})
	c["language"] = overaudio.Language
	c["ad"] = overaudio.Ad
	return []interface{}{c}
}

func flattenOverTeletext(overteletext *OverTeletext) []interface{} {
	c := make(map[string]interface{})
	c["page"] = overteletext.Page
	c["language"] = overteletext.Language
	c["sdh"] = overteletext.Sdh
	return []interface{}{c}
}
