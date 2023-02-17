package vpngateway


type VpnGatewayBgpSettingsInstance0BgpPeeringAddress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway#custom_ips VpnGateway#custom_ips}.
	CustomIps *[]*string `field:"required" json:"customIps" yaml:"customIps"`
}

