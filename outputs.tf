output "vnet_id" {
  value = azurerm_virtual_network.this.id
}

output "vnet_name" {
  value = azurerm_virtual_network.this.name
}

output "vnet_location" {
  value = azurerm_virtual_network.this.location
}

output "vnet_address_space" {
  value = azurerm_virtual_network.this.address_space
}

output "subnet_ids" {
  value = [for subnet in azurerm_subnet.this : subnet.id]
}

output "ddos_protection_plan_id" {
  value = var.enable_ddos_protection_plan ? azurerm_network_ddos_protection_plan.this[0].id : null
}
