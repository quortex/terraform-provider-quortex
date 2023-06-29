package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrganizationWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelWebhook(d *schema.ResourceData) (*Webhook, error) {

	ve := Webhook{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		Url:      d.Get("url").(string),
		Category: d.Get("category").(string),
		PoolUuid: d.Get("pool_uuid").(string),
	}

	return &ve, nil
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	ve, err1 := marshallModelWebhook(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.CreateWebhook(*ve)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Uuid)

	resourceWebhookRead(ctx, d, m)

	return diags
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	webhookId := d.Id()

	webhook, err := c.GetWebhook(webhookId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", webhook.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", webhook.Type); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("url", webhook.Url); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("category", webhook.Category); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("pool_uuid", webhook.PoolUuid); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	webhookId := d.Id()

	ve, err1 := marshallModelWebhook(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.UpdateWebhook(webhookId, *ve)
	if err != nil {
		return diag.FromErr(err)

	}

	d.SetId(o.Uuid)

	resourceWebhookRead(ctx, d, m)

	return diags
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteWebhook(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
