output "resource_group_name" {
  value = azurerm_resource_group.this.name
}

output "resource_group_location" {
  value = azurerm_resource_group.this.location
}

output "vnet_name" {
  value = module.vnet.vnet_name
}

output "vnet_location" {
  value = module.vnet.vnet_location
}

output "subnet_ids" {
  value = module.vnet.subnet_ids
}
