package vpngateway


type VpnGatewayBgpSettingsInstance1BgpPeeringAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/vpn_gateway#custom_ips VpnGateway#custom_ips}.
	CustomIps *[]*string `field:"required" json:"customIps" yaml:"customIps"`
}

