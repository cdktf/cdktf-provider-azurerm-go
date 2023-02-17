package trafficmanagernestedendpoint


type TrafficManagerNestedEndpointCustomHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint#name TrafficManagerNestedEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_nested_endpoint#value TrafficManagerNestedEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

