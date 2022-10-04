package vpngatewayconnection


type VpnGatewayConnectionVpnLinkCustomBgpAddress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ip_address VpnGatewayConnection#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ip_configuration_id VpnGatewayConnection#ip_configuration_id}.
	IpConfigurationId *string `field:"required" json:"ipConfigurationId" yaml:"ipConfigurationId"`
}

