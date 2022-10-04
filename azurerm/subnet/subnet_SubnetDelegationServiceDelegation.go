package subnet


type SubnetDelegationServiceDelegation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet#name Subnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet#actions Subnet#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
}

