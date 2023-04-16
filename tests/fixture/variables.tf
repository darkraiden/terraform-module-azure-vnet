variable "subscription_id" {}

variable "tenant_id" {}

variable "base_name" {
  type    = string
  default = "terratest-vnet"
}

variable "vnet_address_space" {}

variable "location" {
  default = "northeurope"
}

variable "subnets" {
  default = []
}

variable "vnet_location" {
  type    = string
  default = null
}

variable "enable_ddos_protection_plan" {
  default = false
}
