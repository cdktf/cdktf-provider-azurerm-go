package vpngateway


type VpnGatewayBgpSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/vpn_gateway#asn VpnGateway#asn}.
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/vpn_gateway#peer_weight VpnGateway#peer_weight}.
	PeerWeight *float64 `field:"required" json:"peerWeight" yaml:"peerWeight"`
	// instance_0_bgp_peering_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/vpn_gateway#instance_0_bgp_peering_address VpnGateway#instance_0_bgp_peering_address}
	Instance0BgpPeeringAddress *VpnGatewayBgpSettingsInstance0BgpPeeringAddress `field:"optional" json:"instance0BgpPeeringAddress" yaml:"instance0BgpPeeringAddress"`
	// instance_1_bgp_peering_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/vpn_gateway#instance_1_bgp_peering_address VpnGateway#instance_1_bgp_peering_address}
	Instance1BgpPeeringAddress *VpnGatewayBgpSettingsInstance1BgpPeeringAddress `field:"optional" json:"instance1BgpPeeringAddress" yaml:"instance1BgpPeeringAddress"`
}

