package vpnserverconfiguration


type VpnServerConfigurationRadiusServer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#address VpnServerConfiguration#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#score VpnServerConfiguration#score}.
	Score *float64 `field:"required" json:"score" yaml:"score"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#secret VpnServerConfiguration#secret}.
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

