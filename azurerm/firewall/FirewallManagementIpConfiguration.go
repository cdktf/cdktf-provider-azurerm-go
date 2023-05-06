package firewall


type FirewallManagementIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall#name Firewall#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall#public_ip_address_id Firewall#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall#subnet_id Firewall#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

