package restconf

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigBlock() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConfigBlockCreate,
		ReadContext:   resourceConfigBlockRead,
		UpdateContext: resourceConfigBlockUpdate,
		DeleteContext: resourceConfigBlockDelete,

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				id := d.Id()
				d.Set("path", id)

				diags := resourceConfigBlockRead(ctx, d, m)
				if diags.HasError() {
					return nil, fmt.Errorf("failed to import resource: %s", diags[0].Summary)
				}

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConfigBlockCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)

	path := d.Get("path").(string)
	content := d.Get("content").(string)

	// Check if the configuration already exists
	existingConfig, err := client.ReadConfigBlock(ctx, path)
	if err != nil && !strings.Contains(err.Error(), "404 Not Found") {
		return diag.FromErr(err)
	}

	if existingConfig != "" {
		return diag.Errorf("Configuration already exists. Please import the existing resource: terraform import restconf_config_block.example %s", path)
	}

	err = client.CreateConfigBlock(ctx, path, content)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(path)

	return resourceConfigBlockRead(ctx, d, meta)
}

func resourceConfigBlockRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	path := d.Get("path").(string)

	// Read the configuration from the device
	apiResponse, err := client.ReadConfigBlock(ctx, path)
	if err != nil {
		return diag.FromErr(err)
	}

	// Decode and re-encode the API response to ensure consistent formatting
	var apiResponseFormatted map[string]interface{}
	if err := json.Unmarshal([]byte(apiResponse), &apiResponseFormatted); err != nil {
		return diag.FromErr(err)
	}
	apiResponseJson, err := json.Marshal(apiResponseFormatted)
	if err != nil {
		return diag.FromErr(err)
	}

	// Read the stored configuration from the state
	content := d.Get("content").(string)

	// Decode and re-encode the stored configuration to ensure consistent formatting
	var contentFormatted map[string]interface{}
	if content == "" {
		contentFormatted = make(map[string]interface{})
	} else {
		if err := json.Unmarshal([]byte(content), &contentFormatted); err != nil {
			return diag.FromErr(err)
		}
	}

	contentJson, err := json.Marshal(contentFormatted)
	if err != nil {
		return diag.FromErr(err)
	}

	// Compare the formatted API response and the formatted stored configuration
	if string(apiResponseJson) != string(contentJson) {
		d.Set("content", string(apiResponseJson))
	}

	return nil
}

func resourceConfigBlockUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	path := d.Id()
	content := d.Get("content").(string)

	err := client.UpdateConfigBlock(ctx, path, content)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceConfigBlockRead(ctx, d, m)
}

func resourceConfigBlockDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	path := d.Id()

	err := client.DeleteConfigBlock(ctx, path)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
