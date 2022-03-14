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

	for _, stream := range streams {
		st := stream.(map[string]interface{})

		str := Stream{
			Uuid:    st["uuid"].(string),
			Name:    st["name"].(string),
			Enabled: st["enabled"].(bool),
		}

		// Manage srt
		if srt, ok := st["srt"]; ok {
			srts := srt.([]interface{})
			sr := srts[0].(map[string]interface{})
			str.Type = "srt"

			srtt := Srt{
				Latency: sr["latency"].(int),
			}

			// Manage srt listener/caller
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
					}
					srtt.Listener = &listener
				}

			}
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

			// Manage rtmp
		} else if rtmp, ok := st["rtmp"]; ok {
			rtmps := rtmp.([]interface{})
			rt := rtmps[0].(map[string]interface{})
			str.Type = "rtmp"

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
			oi["srt"] = flattenSrt(stream.Srt)
			//oi["rtmp"] = flattenRtmp(stream.Rtmp)
			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

func flattenSrt(srt *Srt) []interface{} {
	c := make(map[string]interface{})
	//c["connection_type"] = (*srt).ConnectionType
	c["latency"] = (*srt).Latency

	return []interface{}{c}
}

func flattenRtmp(rtmp *Rtmp) []interface{} {
	c := make(map[string]interface{})

	return []interface{}{c}
}
