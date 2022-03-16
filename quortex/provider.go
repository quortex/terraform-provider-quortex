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
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUORTEX_HOST", nil),
			},
			"oauth": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_server": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_AUTH_SERVER", nil),
						},
						"client_id": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_CLIENT_ID", nil),
						},
						"client_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_CLIENT_SECRET", nil),
						},
					},
				},
			},
			"api_key": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_server": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_AUTH_SERVER", nil),
						},
						"api_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_APIKEY", nil),
						},
					},
				},
			},
			"basic_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_USERNAME", nil),
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							DefaultFunc: schema.EnvDefaultFunc("QUORTEX_PASSWORD", nil),
						},
					},
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"quortex_admin_dataplane": resourceAdminDataplane(),
			"quortex_ott_pool":        resourceOttPool(),
			"quortex_ott_input":       resourceOttInput(),
			"quortex_ott_processing":  resourceOttProcessing(),
			"quortex_ott_target":      resourceOttTarget(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	var host *string
	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	oauths := d.Get("oauth").([]interface{})
	apikeys := d.Get("api_key").([]interface{})
	basicauths := d.Get("basic_auth").([]interface{})

	for _, oauth := range oauths {
		oaut := oauth.(map[string]interface{})
		authserver := oaut["auth_server"].(string)
		clientid := oaut["client_id"].(string)
		clientsecret := oaut["client_secret"].(string)
		if (authserver != "") && (clientid != "") && (clientsecret != "") {
			c, err := NewClientOauth(host, &authserver, &clientid, &clientsecret)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create Quortex client",
					Detail:   "Unable to authenticate user for authenticated Quortex client with oauth auth",
				})

				return nil, diags
			}
			return c, diags
		}
	}

	for _, apikey := range apikeys {
		apike := apikey.(map[string]interface{})
		key := apike["api_key"].(string)
		authserver := apike["auth_server"].(string)
		if key != "" {
			c, err := NewClientApiKey(host, &key, &authserver)
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
	}

	for _, basicauth := range basicauths {
		basicaut := basicauth.(map[string]interface{})
		username := basicaut["username"].(string)
		password := basicaut["password"].(string)
		if (username != "") && (password != "") {
			c, err := NewClientBasicAuth(host, &username, &password)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create Quortex client",
					Detail:   "Unable to authenticate user for authenticated Quortex client with basic auth",
				})

				return nil, diags
			}
			return c, diags
		}
	}

	c, err := NewClientUnprotected(host)
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
