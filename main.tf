locals {
  location = var.location == null ? var.resource_group.location : var.location
}

resource "azurerm_virtual_network" "this" {
  name                = var.vnet_name
  resource_group_name = var.resource_group.name
  location            = local.location
  address_space       = var.address_space
}

resource "azurerm_subnet" "this" {
  for_each = { for subnet in var.subnets : subnet.name => subnet }

  name                 = each.key
  resource_group_name  = var.resource_group.name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = each.value.address_prefixes
}

resource "azurerm_network_ddos_protection_plan" "this" {
  count = var.enable_ddos_protection_plan ? 1 : 0

  name                = var.vnet_name
  location            = local.location
  resource_group_name = var.resource_group.name
}
