package virtualnetworkgateway


type VirtualNetworkGatewayIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_network_gateway#public_ip_address_id VirtualNetworkGateway#public_ip_address_id}.
	PublicIpAddressId *string `field:"required" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_network_gateway#subnet_id VirtualNetworkGateway#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_network_gateway#private_ip_address_allocation VirtualNetworkGateway#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"optional" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
}

