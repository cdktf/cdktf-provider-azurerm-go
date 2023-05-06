package apimanagementgatewayhostnameconfiguration


type ApiManagementGatewayHostNameConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/api_management_gateway_host_name_configuration#create ApiManagementGatewayHostNameConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/api_management_gateway_host_name_configuration#delete ApiManagementGatewayHostNameConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/api_management_gateway_host_name_configuration#read ApiManagementGatewayHostNameConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/api_management_gateway_host_name_configuration#update ApiManagementGatewayHostNameConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

