package pointtositevpngateway


type PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#address_prefixes PointToSiteVpnGateway#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
}

