package trafficmanagerexternalendpoint


type TrafficManagerExternalEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#create TrafficManagerExternalEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#delete TrafficManagerExternalEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#read TrafficManagerExternalEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#update TrafficManagerExternalEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
