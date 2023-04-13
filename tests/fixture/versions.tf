terraform {
  required_version = "1.4.0"
}

provider "azurerm" {
  // the following parameters will be passed by the user in testing phase
  subscription_id = var.subscription_id
  tenant_id       = var.tenant_id
  features {}
}
