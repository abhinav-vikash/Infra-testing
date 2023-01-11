output "domain_name" {
  value = azurerm_private_dns_zone.zone.name
}

output "record_name" {
  value = azurerm_private_dns_a_record.record.name
}

output "a_record_ip" {
  value = azurerm_private_dns_a_record.record.records
}

output "record_id" {
  value = azurerm_private_dns_a_record.record.id
}

output "fully_qualified_domain_name" {
  value = azurerm_private_dns_a_record.record.fqdn
}