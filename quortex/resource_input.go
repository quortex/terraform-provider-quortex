package quortex

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInput() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInputCreate,
		ReadContext:   resourceInputRead,
		UpdateContext: resourceInputUpdate,
		DeleteContext: resourceInputDelete,
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
			"stream": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"srt": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"latency": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"connection_type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"rtmp": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem:     &schema.Resource{},
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

		if srt, ok := st["srt"]; ok {
			srts := srt.([]interface{})
			sr := srts[0].(map[string]interface{})
			log.Println(sr)

			str.Type = "srt"

			srtt := Srt{
				ConnectionType: sr["connection_type"].(string),
				Latency:        sr["latency"].(int),
			}

			if sr["connection_type"] == "listener" {
				listener := Listener{}
				srtt.Listener = &listener
			}
			str.Srt = &srtt

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

	d.Set("name", input.Name)
	d.Set("identifier", input.Identifier)
	d.Set("published", input.Published)
	streams := flattenInputStreams(&input.Streams)
	d.Set("stream", streams)

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
		ois := make([]interface{}, len(*streams), len(*streams))

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
	c["connection_type"] = (*srt).ConnectionType
	c["latency"] = (*srt).Latency

	return []interface{}{c}
}

func flattenRtmp(rtmp *Rtmp) []interface{} {
	c := make(map[string]interface{})

	return []interface{}{c}
}
