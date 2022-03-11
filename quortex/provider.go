package quortex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUORTEX_HOST", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUORTEX_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("QUORTEX_PASSWORD", nil),
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("QUORTEX_APIKEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"quortex_pool":  resourcePool(),
			"quortex_input": resourceInput(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	apikey := d.Get("api_key").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := NewClient(host, &username, &password, nil)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Quortex client",
				Detail:   "Unable to authenticate user for authenticated Quortex client with username password",
			})

			return nil, diags
		}

		return c, diags
	}

	if apikey != "" {
		c, err := NewClient(host, nil, nil, &apikey)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Quortex client",
				Detail:   "Unable to authenticate user for authenticated Quortex client with api key",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := NewClient(host, nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Quortex client",
			Detail:   "Unable to create anonymous Quortex client without authentication",
		})
		return nil, diags
	}

	return c, diags
}
