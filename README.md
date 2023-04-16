# terraform-module-azure-vnet

Terraform module to create a new Azure Virtual Network and its respective subnets.

## Module Usage

```hcl
resource "azurerm_resource_group" "this" {
    name     = "my-resource-group"
    location = "northeurope"
}

module "myvnet" {
    source = "github.com/darkraiden/terraform-module-azure-vnet"

    vnet_name      = "my-vnet"
    resource_group = azurerm_resource_group.this
    address_space  = ["10.0.0.0/24"]
    subnets        = [
        {
            name = "subnet1"
            address_prefixes = ["10.0.0.0/25"]
        },
        {
            name = "subnet1"
            address_prefixes = ["10.0.0.128/25"]
        }
    ]
}
```

## Test Module

This module comes with a `terratest` test suite. Before running the tests, ensure you have the following:

- Azure cli installed and configured
- Write access to an Azure Subscription/Tenant
- Go installed

To run the tests, execute the following command:

```bash
export TEST_AZURE_SUBSCRIPTION_ID=<your-subscription-id>
export TEST_AZURE_TENANT_ID=<your-tenant-id>

cd tests/
go test -v -timeout 30m -count 1 ./...
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.4.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 3.40.0, <= 3.50.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | >= 3.40.0, <= 3.50.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_network_ddos_protection_plan.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_ddos_protection_plan) | resource |
| [azurerm_subnet.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet) | resource |
| [azurerm_virtual_network.this](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_address_space"></a> [address\_space](#input\_address\_space) | The list of address space prefixes used by the Virtual Network. | `list(string)` | n/a | yes |
| <a name="input_enable_ddos_protection_plan"></a> [enable\_ddos\_protection\_plan](#input\_enable\_ddos\_protection\_plan) | [Optional] Should a DDoS Protection Plan be associated with the Virtual Network? | `bool` | `false` | no |
| <a name="input_location"></a> [location](#input\_location) | The location/region where the Virtual Network should exist. | `string` | `null` | no |
| <a name="input_resource_group"></a> [resource\_group](#input\_resource\_group) | The Resource Group where the Virtual Network should exist. | <pre>object({<br>    id       = string<br>    name     = string<br>    location = string<br>  })</pre> | n/a | yes |
| <a name="input_subnets"></a> [subnets](#input\_subnets) | [Optional] A list of subnets to be created within the Virtual Network. | <pre>list(object({<br>    name             = string<br>    address_prefixes = list(string)<br>  }))</pre> | `[]` | no |
| <a name="input_vnet_name"></a> [vnet\_name](#input\_vnet\_name) | The name of the Virtual Network. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ddos_protection_plan_id"></a> [ddos\_protection\_plan\_id](#output\_ddos\_protection\_plan\_id) | The ID of the DDoS Protection Plan associated with the Virtual Network. |
| <a name="output_subnet_ids"></a> [subnet\_ids](#output\_subnet\_ids) | The list of subnet IDs associated with the Virtual Network. |
| <a name="output_vnet_address_space"></a> [vnet\_address\_space](#output\_vnet\_address\_space) | The list of address space prefixes used by the Virtual Network. |
| <a name="output_vnet_id"></a> [vnet\_id](#output\_vnet\_id) | The ID of the Virtual Network. |
| <a name="output_vnet_location"></a> [vnet\_location](#output\_vnet\_location) | The location of the Virtual Network. Defaults to the Resource Group's location if not set. |
| <a name="output_vnet_name"></a> [vnet\_name](#output\_vnet\_name) | The name of the Virtual Network. |
<!-- END_TF_DOCS -->
