package virtualnetworkgatewayconnection


type VirtualNetworkGatewayConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/virtual_network_gateway_connection#create VirtualNetworkGatewayConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/virtual_network_gateway_connection#delete VirtualNetworkGatewayConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/virtual_network_gateway_connection#read VirtualNetworkGatewayConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/virtual_network_gateway_connection#update VirtualNetworkGatewayConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

