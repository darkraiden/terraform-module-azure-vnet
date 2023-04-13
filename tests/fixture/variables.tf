variable "subscription_id" {}

variable "tenant_id" {}

variable "base_name" {
  type = string
  // I personally like to prefix all my resources with `terratest`
  // this helps me target them with clean up scripts that I run daily
  // (yes, it does happen that terratest fails and orphan resources are left behind)
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
  type = string
  // defaulting once again to null to allow us to use the location of the resource group
  default = null
}
