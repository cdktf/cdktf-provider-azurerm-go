package sqlvirtualnetworkrule


type SqlVirtualNetworkRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/sql_virtual_network_rule#create SqlVirtualNetworkRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/sql_virtual_network_rule#delete SqlVirtualNetworkRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/sql_virtual_network_rule#read SqlVirtualNetworkRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/sql_virtual_network_rule#update SqlVirtualNetworkRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

