variable "resource_group_location" {
  default     = "eastus"
  description = "Location of the resource group."
}

variable "resource_group_name" {
  default     = "rg_ada_adls"
  description = "Prefix of the resource group name that's combined with a random ID so name is unique in your Azure subscription."
}

variable "storage_account_name" {
  default     = "adaadls"
  description = "storage account name"
}