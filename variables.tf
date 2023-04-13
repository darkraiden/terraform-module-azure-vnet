variable "vnet_name" {
  type = string
}

variable "resource_group" {
  // This object will make sense once we have created
  // our root module in the `fixture` folder
  type = object({
    id       = string
    name     = string
    location = string
  })
}

variable "location" {
  type    = string
  default = null
}

variable "address_space" {
  type = list(string)
}

variable "subnets" {
  type = list(object({
    name             = string
    address_prefixes = list(string)
  }))
  default = []
}

variable "enable_ddos_protection_plan" {
  type    = bool
  default = false
}
