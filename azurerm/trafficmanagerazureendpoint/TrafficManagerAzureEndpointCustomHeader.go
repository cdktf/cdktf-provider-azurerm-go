package trafficmanagerazureendpoint


type TrafficManagerAzureEndpointCustomHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#name TrafficManagerAzureEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#value TrafficManagerAzureEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

