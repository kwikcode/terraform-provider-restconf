package restconf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"restconf_config_block": resourceConfigBlock(),
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RESTCONF_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RESTCONF_PASSWORD", ""),
			},
			"device_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RESTCONF_DEVICE_HOST", ""),
			},
			"device_port": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("RESTCONF_DEVICE_PORT", "443"),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	deviceHost := d.Get("device_host").(string)
	devicePort := d.Get("device_port").(string)

	client, err := NewClient(username, password, deviceHost, devicePort)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}
