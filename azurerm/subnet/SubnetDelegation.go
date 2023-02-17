package subnet


type SubnetDelegation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet#name Subnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// service_delegation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subnet#service_delegation Subnet#service_delegation}
	ServiceDelegation *SubnetDelegationServiceDelegation `field:"required" json:"serviceDelegation" yaml:"serviceDelegation"`
}

