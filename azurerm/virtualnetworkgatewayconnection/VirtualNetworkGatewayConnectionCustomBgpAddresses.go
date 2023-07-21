package virtualnetworkgatewayconnection


type VirtualNetworkGatewayConnectionCustomBgpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/virtual_network_gateway_connection#primary VirtualNetworkGatewayConnection#primary}.
	Primary *string `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/virtual_network_gateway_connection#secondary VirtualNetworkGatewayConnection#secondary}.
	Secondary *string `field:"required" json:"secondary" yaml:"secondary"`
}

