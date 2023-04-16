# Terraform RESTCONF Provider

The Terraform RESTCONF provider allows you to manage configuration blocks on network devices using the RESTCONF API.

## Examples

Here are a few examples of how to use the RESTCONF provider in your Terraform configuration:

### Example 1: Update NTP settings on a Cisco IOS XE device

```hcl
provider "restconf" {
  host     = "https://192.0.2.1"
  username = "admin"
  password = "admin"
}

resource "restconf_config" "example" {
  path  = "/restconf/data/Cisco-IOS-XE-native:native/ntp"
  value = jsonencodejsonencode({
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
provider "restconf" {
  host     = "https://192.0.2.1"
  username = "admin"
  password = "admin"
}

resource "restconf_config" "example" {
    path  = "/restconf/data/Cisco-IOS-XE-native:native/banner"
    value = jsonencodejsonencode({
                "Cisco-IOS-XE-native:banner": {
                    "motd": "Welcome to the network!"
                }
            })
    }
```

## Example 3 - Import NTP settings from a Cisco IOS XE device
```
terraform import restconf_config_block.example /restconf/data/Cisco-IOS-XE-native:native/ntp
```

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or later
- [Go](https://golang.org/doc/install) 1.16 or later (to build the provider plugin)
- [GoReleaser](https://goreleaser.com/install/) (optional, for building releases)
- [Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation) (optional, for debugging)

## Building the Provider

1. Clone the provider repository:

```bash
git clone https://github.com/kwikcode/terraform-provider-restconf.git
```

2. Change into the repository directory:

```bash
cd terraform-provider-restconf
```

3. Build the provider binary:

```bash
go build -o terraform-provider-restconf
```

4. Move the provider binary to the Terraform plugins directory:

For Terraform 0.12.x and earlier, you can put the provider binary in the same directory as your Terraform configuration files or in the user plugins directory, which is usually `~/.terraform.d/plugins` on UNIX-like systems and `%APPDATA%\terraform.d\plugins` on Windows.

## Building Releases with GoReleaser

To build releases using GoReleaser, follow these steps:

1. Install GoReleaser:

```bash
curl -sL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
```

2. Configure your `.goreleaser.yml` file to define the build settings and artifacts.

3. Run GoReleaser to build the release artifacts:

```bash
goreleaser --rm-dist
```

## Debugging with Delve and Visual Studio Code

1. Install Delve:

Follow the [Delve installation guide](https://github.com/go-delve/delve/tree/master/Documentation/installation) for your platform.

2. Configure Visual Studio Code:

- Install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) for Visual Studio Code.

3. Start debugging:
- Set a breakpoint in the code where you want to debug.
- Open the Visual Studio Code command palette (Cmd+Shift+P or Ctrl+Shift+P) and run the "Debug: Start Debugging" command.
- DLV will output the debug command such as:
```
TF_REATTACH_PROVIDERS='{"github.com/kwikcode/restconf":{"Protocol":"grpc","ProtocolVersion":5,"Pid":86798,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/9c/cw8hxtyj4fz2kfmglh82g0s80000gq/T/plugin1217677869"}}}'
```
- Copy the command and paste it into your terminal
- Run terraform plan or apply
- The debugger will stop at the breakpoint you set
