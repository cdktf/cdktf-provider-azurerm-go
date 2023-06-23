package subnet


type SubnetDelegationServiceDelegation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/subnet#name Subnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.0/docs/resources/subnet#actions Subnet#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
}

