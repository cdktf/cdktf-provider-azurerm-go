package trafficmanagerexternalendpoint


type TrafficManagerExternalEndpointSubnet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#first TrafficManagerExternalEndpoint#first}.
	First *string `field:"required" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#last TrafficManagerExternalEndpoint#last}.
	Last *string `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_external_endpoint#scope TrafficManagerExternalEndpoint#scope}.
	Scope *float64 `field:"optional" json:"scope" yaml:"scope"`
}

