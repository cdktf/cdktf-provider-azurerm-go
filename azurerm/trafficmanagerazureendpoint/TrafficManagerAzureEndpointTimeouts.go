package trafficmanagerazureendpoint


type TrafficManagerAzureEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#create TrafficManagerAzureEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#delete TrafficManagerAzureEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#read TrafficManagerAzureEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_azure_endpoint#update TrafficManagerAzureEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
