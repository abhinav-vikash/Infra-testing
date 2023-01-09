output "resource_group_name" {
  value = azurerm_resource_group.rg.name
}

output "storage_account_name" {
  value = azurerm_storage_account.sgaccount.name
}

output "storage_account_account_tier" {
  value = azurerm_storage_account.sgaccount.account_tier
}

output "storage_account_account_kind" {
  value = azurerm_storage_account.sgaccount.account_kind
}

output "storageV2_container_name" {
  value = azurerm_storage_data_lake_gen2_filesystem.container.name
}