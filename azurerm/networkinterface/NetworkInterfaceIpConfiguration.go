package networkinterface


type NetworkInterfaceIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#name NetworkInterface#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#private_ip_address_allocation NetworkInterface#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"required" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#gateway_load_balancer_frontend_ip_configuration_id NetworkInterface#gateway_load_balancer_frontend_ip_configuration_id}.
	GatewayLoadBalancerFrontendIpConfigurationId *string `field:"optional" json:"gatewayLoadBalancerFrontendIpConfigurationId" yaml:"gatewayLoadBalancerFrontendIpConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#primary NetworkInterface#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#private_ip_address NetworkInterface#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#private_ip_address_version NetworkInterface#private_ip_address_version}.
	PrivateIpAddressVersion *string `field:"optional" json:"privateIpAddressVersion" yaml:"privateIpAddressVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#public_ip_address_id NetworkInterface#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_interface#subnet_id NetworkInterface#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

