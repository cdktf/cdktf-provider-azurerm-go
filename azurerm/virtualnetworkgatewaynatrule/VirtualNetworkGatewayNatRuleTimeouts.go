package virtualnetworkgatewaynatrule


type VirtualNetworkGatewayNatRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_nat_rule#create VirtualNetworkGatewayNatRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_nat_rule#delete VirtualNetworkGatewayNatRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_nat_rule#read VirtualNetworkGatewayNatRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_network_gateway_nat_rule#update VirtualNetworkGatewayNatRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
