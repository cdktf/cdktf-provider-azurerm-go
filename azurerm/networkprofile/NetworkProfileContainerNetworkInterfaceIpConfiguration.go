package networkprofile


type NetworkProfileContainerNetworkInterfaceIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_profile#name NetworkProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_profile#subnet_id NetworkProfile#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

