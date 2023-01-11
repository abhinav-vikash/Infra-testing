resource "azurerm_resource_group" "rg" {
  name     = var.dns_resource_group_name
  location = var.dns_location
}

resource "azurerm_private_dns_zone" "zone" {
  name                = var.domain_name
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_private_dns_a_record" "record" {
  name                = var.record_name
  zone_name           = azurerm_private_dns_zone.zone.name
  resource_group_name = azurerm_resource_group.rg.name
  ttl                 = 300
  records             = ["191.10.0.110","191.10.0.111"]
}