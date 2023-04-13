vnet_address_space = ["10.0.0.0/24"]
subnets = [
  {
    name             = "subnet1"
    address_prefixes = ["10.0.0.0/25"]
  },
  {
    name             = "subnet2"
    address_prefixes = ["10.0.0.128/25"]
  }
]
