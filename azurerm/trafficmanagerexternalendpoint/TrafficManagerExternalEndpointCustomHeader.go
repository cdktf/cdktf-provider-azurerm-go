package trafficmanagerexternalendpoint


type TrafficManagerExternalEndpointCustomHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#name TrafficManagerExternalEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#value TrafficManagerExternalEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

