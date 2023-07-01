package virtualnetworkgateway


type VirtualNetworkGatewayBgpSettingsPeeringAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/virtual_network_gateway#apipa_addresses VirtualNetworkGateway#apipa_addresses}.
	ApipaAddresses *[]*string `field:"optional" json:"apipaAddresses" yaml:"apipaAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/virtual_network_gateway#ip_configuration_name VirtualNetworkGateway#ip_configuration_name}.
	IpConfigurationName *string `field:"optional" json:"ipConfigurationName" yaml:"ipConfigurationName"`
}

