package privatednszonevirtualnetworklink


type PrivateDnsZoneVirtualNetworkLinkTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_zone_virtual_network_link#create PrivateDnsZoneVirtualNetworkLink#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_zone_virtual_network_link#delete PrivateDnsZoneVirtualNetworkLink#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_zone_virtual_network_link#read PrivateDnsZoneVirtualNetworkLink#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_zone_virtual_network_link#update PrivateDnsZoneVirtualNetworkLink#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
