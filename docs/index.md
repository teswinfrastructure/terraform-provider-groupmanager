# GroupManager Provider

The GroupManager Provider can be used to create and manage only owned groups in [Azure Active Directory](https://azure.microsoft.com/en-us/services/active-directory/) using the [Microsoft Graph](https://docs.microsoft.com/en-us/graph/overview) API. This provider was created on the basis of the well-known [azuread](https://registry.terraform.io/providers/hashicorp/azuread/latest) provider (version 2.36.0). Documentation regarding the [Data Sources](https://www.terraform.io/docs/language/data-sources/index.html) and [Resources](https://www.terraform.io/docs/language/resources/index.html) supported by the GroupManager Provider can be found in the navigation to the left.

## Example Usage

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
  tenant_id = "00000000-0000-0000-0000-000000000000"
}

# Create a group
data "groupmanager_client_config" "current" {}

resource "groupmanager_group" "example" {
  display_name     = "example"
  owners           = [data.groupmanager_client_config.current.object_id]
  security_enabled = true
}
```

## Authenticating to Azure Active Directory

Terraform supports a number of different methods for authenticating to Azure Active Directory:

* Authenticating to Azure Active Directory using Managed Identity
* Authenticating to Azure Active Directory using a Service Principal and a Client Certificate
* Authenticating to Azure Active Directory using a Service Principal and a Client Secret

## Argument Reference

The following arguments are supported:

* `client_id` - (Optional) The Client ID which should be used when authenticating as a service principal. This can also be sourced from the `ARM_CLIENT_ID` environment variable.
* `tenant_id` - (Optional) The Tenant ID which should be used. This can also be sourced from the `ARM_TENANT_ID` environment variable.

---

When authenticating as a Service Principal using a Client Certificate, the following fields can be set:

* `client_certificate` - (Optional) A base64-encoded PKCS#12 bundle to be used as the client certificate for authentication. This can also be sourced from the `ARM_CLIENT_CERTIFICATE` environment variable.
* `client_certificate_password` - (Optional) The password for decrypting the client certificate bundle. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PASSWORD` environment variable.
* `client_certificate_path` - (Optional) The path to a PKCS#12 bundle (.pfx file) to be used as the client certificate for authentication. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` environment variable.

---

When authenticating as a Service Principal using a Client Secret, the following fields can be set:

* `client_secret` - (Optional) The application password to be used when authenticating using a client secret. This can also be sourced from the `ARM_CLIENT_SECRET` environment variable.

---

When authenticating as a Service Principal using Open ID Connect, the following fields can be set:

* `oidc_request_token` - (Optional) The bearer token for the request to the OIDC provider. This can also be sourced from the `ARM_OIDC_REQUEST_TOKEN` or `ACTIONS_ID_TOKEN_REQUEST_TOKEN` Environment Variables.
* `oidc_request_url` - (Optional) The URL for the OIDC provider from which to request an ID token. This can also be sourced from the `ARM_OIDC_REQUEST_URL` or `ACTIONS_ID_TOKEN_REQUEST_TOKEN` Environment Variables.
* `oidc_token` - (Optional) The ID token when authenticating using OpenID Connect (OIDC). This can also be sourced from the `ARM_OIDC_TOKEN` Environment Variable.
* `oidc_token_file_path` - (Optional) The path to a file containing an ID token when authenticating using OpenID Connect (OIDC). This can also be sourced from the `ARM_OIDC_TOKEN_FILE_PATH` Environment Variable.
* `use_oidc` - (Optional) Should OIDC be used for Authentication? This can also be sourced from the `ARM_USE_OIDC` Environment Variable. Defaults to `false`.

---

When authenticating using Managed Identity, the following fields can be set:

* `msi_endpoint` - (Optional) The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically. This can also be sourced from the `ARM_MSI_ENDPOINT` environment variable.
* `use_msi` - (Optional) Should a Managed Identity be used for authentication? This can also be sourced from the `ARM_USE_MSI` environment variable. Defaults to `false`.

---

## Logging and Tracing

Logging output can be controlled with the `TF_LOG` or `TF_PROVIDER_LOG` environment variables. Exporting `TF_LOG=DEBUG` will increase the log verbosity and emit HTTP request and response traces to stdout when running Terraform. This output is very useful when reporting a bug in the provider.

Note that whilst we make every effort to remove authentication tokens from HTTP traces, they can still contain very identifiable and personal information which you should carefully censor before posting on our issue tracker.
