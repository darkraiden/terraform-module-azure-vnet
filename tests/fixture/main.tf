resource "random_id" "this" {
  byte_length = 8
}

locals {
  // this will allow us to have a unique name to use when creating resources
  // it is very useful when testing infrastructure that relies on unique names
  // especially when running parallel tests
  resources_name = format("%s-%s", var.base_name, random_id.this.hex)
}

resource "azurerm_resource_group" "this" {
  name     = local.resources_name
  location = var.location
}

module "vnet" {
  source    = "../../"
  vnet_name = local.resources_name // using same name for both vnet and rg is totally fine
  // by passing the entire rg object,
  // we have access to the attributes we defined previously in the child module
  // (id, name, and location)
  resource_group = azurerm_resource_group.this
  location       = var.vnet_location
  address_space  = var.vnet_address_space
  subnets        = var.subnets
}
