---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "restconf_config_block Resource - terraform-provider-restconf"
subcategory: ""
description: |-
  
---

# restconf_config_block (Resource)
This resource provides an idempotent way to manage configuration blocks on network devices using the RESTCONF API.
It will handle the creation, update and deletion of configuration blocks on network devices.
You must import the resource before you can use it if it already exists on the device.

# Examples

Here are a few examples of how to use the RESTCONF provider in your Terraform configuration:

## Example 1: Update NTP settings on a Cisco IOS XE device

```hcl
terraform {
  required_providers {
    restconf = {
      source  = "kwikcode/restconf"
    }
  }
}

provider "restconf" {
  username = "admin"
  password = "admin"
}

resource "restconf_config" "example" {
  path  = "https://192.0.2.1/restconf/data/Cisco-IOS-XE-native:native/ntp"
  value = jsonencode({
                "Cisco-IOS-XE-native:ntp": {
                    "Cisco-IOS-XE-ntp:server": {
                        "server-list": [
                            {
                                "ip-address": "ntp1.example.com"
                            },
                            {
                                "ip-address": "ntp2.example.com"
                            }
                        ]
                    }
                }
           })
}
```

## Example 2 - Update Banner settings on a Cisco IOS XE device

```hcl
terraform {
  required_providers {
    restconf = {
      source  = "kwikcode/restconf"
    }
  }
}

provider "restconf" {
  username = "admin"
  password = "admin"
}

resource "restconf_config" "example" {
    path  = "https://192.0.2.1/restconf/data/Cisco-IOS-XE-native:native/banner"
    value = jsonencode({
                "Cisco-IOS-XE-native:banner": {
                    "motd": "Welcome to the network!"
                }
            })
    }
```

## Example 3 - Import NTP settings from a Cisco IOS XE device
```
terraform import restconf_config_block.example "https://192.0.2.1/restconf/data/Cisco-IOS-XE-native:native/ntp"
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String)
- `path` (String)

### Read-Only

- `id` (String) The ID of this resource.


