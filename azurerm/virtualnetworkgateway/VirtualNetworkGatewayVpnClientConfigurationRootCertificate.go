package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationRootCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway#public_cert_data VirtualNetworkGateway#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

