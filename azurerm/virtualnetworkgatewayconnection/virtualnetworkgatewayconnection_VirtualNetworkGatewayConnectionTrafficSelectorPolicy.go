package virtualnetworkgatewayconnection


type VirtualNetworkGatewayConnectionTrafficSelectorPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#local_address_cidrs VirtualNetworkGatewayConnection#local_address_cidrs}.
	LocalAddressCidrs *[]*string `field:"required" json:"localAddressCidrs" yaml:"localAddressCidrs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_connection#remote_address_cidrs VirtualNetworkGatewayConnection#remote_address_cidrs}.
	RemoteAddressCidrs *[]*string `field:"required" json:"remoteAddressCidrs" yaml:"remoteAddressCidrs"`
}

