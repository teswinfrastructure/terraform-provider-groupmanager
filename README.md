# GroupManager Provider for Owned Azure Active Directory Group

The GroupManager Provider can be used to create and manage only owned groups in [Azure Active Directory](https://azure.microsoft.com/en-us/services/active-directory/) using the [Microsoft Graph](https://docs.microsoft.com/en-us/graph/overview) API. This provider was created on the basis of the well-known [azuread](https://registry.terraform.io/providers/hashicorp/azuread/latest) provider (version 2.36.0).

**NOTE:** Version 1.3.0 and above of this provider requires Terraform 0.12 or later.

## Usage Example

```hcl
# Configure Terraform
terraform {
  required_providers {
    groupmanager = {
      source = "teswinfrastructure/groupmanager"
      version = "1.1.0"
    }
  }
}

# Configure the GroupManager Provider
provider "groupmanager" {
  # NOTE: Environment Variables can also be used for Service Principal authentication

  # client_id     = "..."
  # client_secret = "..."
  # tenant_id     = "..."
}

# Create a group
data "groupmanager_client_config" "current" {}

resource "groupmanager_group" "example" {
  display_name     = "example"
  owners           = [data.groupmanager_client_config.current.object_id]
  security_enabled = true
}
```
