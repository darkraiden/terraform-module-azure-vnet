variable "vnet_name" {
  description = "The name of the Virtual Network."
  type        = string
}

variable "resource_group" {
  description = "The Resource Group where the Virtual Network should exist."
  type = object({
    id       = string
    name     = string
    location = string
  })
}

variable "location" {
  description = "The location/region where the Virtual Network should exist."
  type        = string
  default     = null
}

variable "address_space" {
  description = "The list of address space prefixes used by the Virtual Network."
  type        = list(string)
}

variable "subnets" {
  description = "[Optional] A list of subnets to be created within the Virtual Network."
  type = list(object({
    name             = string
    address_prefixes = list(string)
  }))
  default = []
}

variable "enable_ddos_protection_plan" {
  description = "[Optional] Should a DDoS Protection Plan be associated with the Virtual Network?"
  type        = bool
  default     = false
}
