package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBucketCreate,
		ReadContext:   resourceBucketRead,
		UpdateContext: resourceBucketUpdate,
		DeleteContext: resourceBucketDelete,
		Schema: map[string]*schema.Schema{
			"dataplane_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"s3": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func marshallModelBucket(d *schema.ResourceData) (*Bucket, error) {

	ve := Bucket{
		Name:   d.Get("name").(string),
		Region: d.Get("region").(string),
		Type:   d.Get("type").(string),
		Label:  d.Get("label").(string),
	}

	s3s := d.Get("s3").([]interface{})
	for _, s3 := range s3s {
		s3e := s3.(map[string]interface{})
		sc := S3{
			AccessKey: s3e["access_key"].(string),
			SecretKey: s3e["secret_key"].(string),
		}
		ve.S3 = &sc
	}

	return &ve, nil
}

func resourceBucketCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	dataplaneId := d.Get("dataplane_id").(string)

	ve, err1 := marshallModelBucket(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.CreateBucket(dataplaneId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceBucketRead(ctx, d, m)

	return diags
}

func resourceBucketRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	dataplaneId := d.Get("dataplane_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucketId := d.Id()

	bucket, err := c.GetBucket(dataplaneId, bucketId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", bucket.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("region", bucket.Region); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("label", bucket.Label); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", bucket.Type); err != nil {
		return diag.FromErr(err)
	}

	s3s := flattenBucketS3(bucket.S3)
	if err := d.Set("s3", s3s); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceBucketUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*Client)
	dataplaneId := d.Get("dataplane_id").(string)
	bucketId := d.Id()

	ve, err1 := marshallModelBucket(d)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	o, err2 := c.UpdateBucket(dataplaneId, bucketId, *ve)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	d.SetId(o.Uuid)

	resourceBucketRead(ctx, d, m)

	return diags
}

func resourceBucketDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	dataplaneId := d.Get("dataplane_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bucketId := d.Id()

	err := c.DeleteBucket(dataplaneId, bucketId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenBucketS3(s3 *S3) []interface{} {

	c := make(map[string]interface{})
	if s3 != nil {
		c["access_key"] = s3.AccessKey
		c["secret_key"] = s3.SecretKey
	}
	return []interface{}{c}
}
