package virtualnetwork


type VirtualNetworkSubnet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/virtual_network#address_prefix VirtualNetwork#address_prefix}.
	AddressPrefix *string `field:"optional" json:"addressPrefix" yaml:"addressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/virtual_network#id VirtualNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/virtual_network#name VirtualNetwork#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/virtual_network#security_group VirtualNetwork#security_group}.
	SecurityGroup *string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

