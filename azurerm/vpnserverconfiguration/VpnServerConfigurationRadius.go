package vpnserverconfiguration


type VpnServerConfigurationRadius struct {
	// client_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#client_root_certificate VpnServerConfiguration#client_root_certificate}
	ClientRootCertificate interface{} `field:"optional" json:"clientRootCertificate" yaml:"clientRootCertificate"`
	// server block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#server VpnServerConfiguration#server}
	Server interface{} `field:"optional" json:"server" yaml:"server"`
	// server_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#server_root_certificate VpnServerConfiguration#server_root_certificate}
	ServerRootCertificate interface{} `field:"optional" json:"serverRootCertificate" yaml:"serverRootCertificate"`
}
