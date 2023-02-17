package trafficmanagernestedendpoint


type TrafficManagerNestedEndpointSubnet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint#first TrafficManagerNestedEndpoint#first}.
	First *string `field:"required" json:"first" yaml:"first"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint#last TrafficManagerNestedEndpoint#last}.
	Last *string `field:"optional" json:"last" yaml:"last"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint#scope TrafficManagerNestedEndpoint#scope}.
	Scope *float64 `field:"optional" json:"scope" yaml:"scope"`
}

