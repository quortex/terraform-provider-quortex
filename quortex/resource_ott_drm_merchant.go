package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOttDrmMerchant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDrmMerchantCreate,
		ReadContext:   resourceDrmMerchantRead,
		UpdateContext: resourceDrmMerchantUpdate,
		DeleteContext: resourceDrmMerchantDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"castlabs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merchant_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"aes_iv": {
							Type:     schema.TypeString,
							Required: true,
						},
						"aes_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"drm_server": {
							Type:     schema.TypeString,
							Required: true,
						},
						"key_seed_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"auth_creds_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"irdeto": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merchant_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"drm_server": {
							Type:     schema.TypeString,
							Required: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"ksm": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drm_server": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"mdrm": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_server": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_secret": {
							Type:     schema.TypeString,
							Required: true,
						},
						"drm_server": {
							Type:     schema.TypeString,
							Required: true,
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

func marshallModelDrmMerchant(d *schema.ResourceData) (*DrmMerchant, error) {

	castlabs := d.Get("castlabs").([]interface{})
	irdetos := d.Get("irdeto").([]interface{})
	ksms := d.Get("ksm").([]interface{})
	mdrms := d.Get("mdrm").([]interface{})
	ve := DrmMerchant{
		Name: d.Get("name").(string),
	}

	for _, castlab := range castlabs {
		if castlab != nil {
			cast := castlab.(map[string]interface{})
			cas := Castlabs{
				MerchantName: cast["merchant_name"].(string),
				AesIv:        cast["aes_iv"].(string),
				AesKey:       cast["aes_key"].(string),
				DrmServer:    cast["drm_server"].(string),
				KeySeedId:    cast["key_seed_id"].(string),
				AuthCredsId:  cast["auth_creds_id"].(string),
			}
			ve.Castlabs = &cas
			ve.Type = "castlabs"
		}
	}

	for _, irdeto := range irdetos {
		if irdeto != nil {
			irdet := irdeto.(map[string]interface{})
			irde := Irdeto{
				MerchantName: irdet["merchant_name"].(string),
				DrmServer:    irdet["drm_server"].(string),
				Username:     irdet["username"].(string),
				Password:     irdet["password"].(string),
			}
			ve.Irdeto = &irde
			ve.Type = "irdeto"
		}
	}

	for _, ksm := range ksms {
		if ksm != nil {
			ks := ksm.(map[string]interface{})
			k := Ksm{
				DrmServer: ks["drm_server"].(string),
			}
			ve.Ksm = &k
			ve.Type = "ksm"
		}
	}

	for _, mdrm := range mdrms {
		if mdrm != nil {
			mdr := mdrm.(map[string]interface{})
			md := Mdrm{
				AuthServer:   mdr["auth_server"].(string),
				ClientId:     mdr["client_id"].(string),
				ClientSecret: mdr["client_secret"].(string),
				DrmServer:    mdr["drm_server"].(string),
			}
			ve.Mdrm = &md
			ve.Type = "mdrm"
		}
	}

	return &ve, nil
}

func resourceDrmMerchantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*Client)

	ve, err1 := marshallModelDrmMerchant(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.CreateDrmMerchant(*ve)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(o.Uuid)

	resourceDrmMerchantRead(ctx, d, m)

	return diags
}

func resourceDrmMerchantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	drmmerchantId := d.Id()

	drmmerchant, err := c.GetDrmMerchant(drmmerchantId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", drmmerchant.Name); err != nil {
		return diag.FromErr(err)
	}

	castlabs := flattenDrmMerchantCastlabs(drmmerchant.Castlabs)
	if err := d.Set("castlabs", castlabs); err != nil {
		return diag.FromErr(err)
	}

	irdeto := flattenDrmMerchantIrdeto(drmmerchant.Irdeto)
	if err := d.Set("irdeto", irdeto); err != nil {
		return diag.FromErr(err)
	}

	ksm := flattenDrmMerchantKsm(drmmerchant.Ksm)
	if err := d.Set("ksm", ksm); err != nil {
		return diag.FromErr(err)
	}

	mdrm := flattenDrmMerchantMdrm(drmmerchant.Mdrm)
	if err := d.Set("mdrm", mdrm); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDrmMerchantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	drmmerchantId := d.Id()

	ve, err1 := marshallModelDrmMerchant(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err := c.UpdateDrmMerchant(drmmerchantId, *ve)
	if err != nil {
		return diag.FromErr(err)

	}

	d.SetId(o.Uuid)

	resourceDrmMerchantRead(ctx, d, m)

	return diags
}

func resourceDrmMerchantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteDrmMerchant(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenDrmMerchantCastlabs(castlabs *Castlabs) []interface{} {

	c := make(map[string]interface{})
	if castlabs != nil {
		c["merchant_name"] = castlabs.MerchantName
		c["aes_iv"] = castlabs.AesIv
		c["aes_key"] = castlabs.AesKey
		c["drm_server"] = castlabs.DrmServer
		c["key_seed_id"] = castlabs.KeySeedId
		c["auth_creds_id"] = castlabs.AuthCredsId
	}
	return []interface{}{c}
}

func flattenDrmMerchantIrdeto(irdeto *Irdeto) []interface{} {

	c := make(map[string]interface{})
	if irdeto != nil {
		c["merchant_name"] = irdeto.MerchantName
		c["drm_server"] = irdeto.DrmServer
		c["username"] = irdeto.Username
		c["password"] = irdeto.Password
	}
	return []interface{}{c}
}

func flattenDrmMerchantKsm(ksm *Ksm) []interface{} {

	c := make(map[string]interface{})
	if ksm != nil {
		c["drm_server"] = ksm.DrmServer
	}
	return []interface{}{c}
}

func flattenDrmMerchantMdrm(mdrm *Mdrm) []interface{} {

	c := make(map[string]interface{})
	if mdrm != nil {
		c["auth_server"] = mdrm.AuthServer
		c["client_id"] = mdrm.ClientId
		c["client_secret"] = mdrm.ClientSecret
		c["drm_server"] = mdrm.DrmServer
	}
	return []interface{}{c}
}
