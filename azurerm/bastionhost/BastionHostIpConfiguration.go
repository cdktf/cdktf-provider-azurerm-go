package bastionhost


type BastionHostIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bastion_host#name BastionHost#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bastion_host#public_ip_address_id BastionHost#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/bastion_host#subnet_id BastionHost#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}
