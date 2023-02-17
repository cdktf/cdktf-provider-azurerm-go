package localnetworkgateway


type LocalNetworkGatewayBgpSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/local_network_gateway#asn LocalNetworkGateway#asn}.
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/local_network_gateway#bgp_peering_address LocalNetworkGateway#bgp_peering_address}.
	BgpPeeringAddress *string `field:"required" json:"bgpPeeringAddress" yaml:"bgpPeeringAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/local_network_gateway#peer_weight LocalNetworkGateway#peer_weight}.
	PeerWeight *float64 `field:"optional" json:"peerWeight" yaml:"peerWeight"`
}

