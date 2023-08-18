package vpngatewaynatrule


type VpnGatewayNatRuleInternalMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/vpn_gateway_nat_rule#address_space VpnGatewayNatRule#address_space}.
	AddressSpace *string `field:"required" json:"addressSpace" yaml:"addressSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/vpn_gateway_nat_rule#port_range VpnGatewayNatRule#port_range}.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
}

