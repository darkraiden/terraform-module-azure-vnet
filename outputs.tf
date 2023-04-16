output "vnet_id" {
  description = "The ID of the Virtual Network."
  value       = azurerm_virtual_network.this.id
}

output "vnet_name" {
  description = "The name of the Virtual Network."
  value       = azurerm_virtual_network.this.name
}

output "vnet_location" {
  description = "The location of the Virtual Network. Defaults to the Resource Group's location if not set."
  value       = azurerm_virtual_network.this.location
}

output "vnet_address_space" {
  description = "The list of address space prefixes used by the Virtual Network."
  value       = azurerm_virtual_network.this.address_space
}

output "subnet_ids" {
  description = "The list of subnet IDs associated with the Virtual Network."
  value       = [for subnet in azurerm_subnet.this : subnet.id]
}

output "ddos_protection_plan_id" {
  description = "The ID of the DDoS Protection Plan associated with the Virtual Network."
  value       = var.enable_ddos_protection_plan ? azurerm_network_ddos_protection_plan.this[0].id : ""
}
