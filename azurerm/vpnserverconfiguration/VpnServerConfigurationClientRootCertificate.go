package vpnserverconfiguration


type VpnServerConfigurationClientRootCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#name VpnServerConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration#public_cert_data VpnServerConfiguration#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

