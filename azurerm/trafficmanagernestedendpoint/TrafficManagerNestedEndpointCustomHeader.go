package trafficmanagernestedendpoint


type TrafficManagerNestedEndpointCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/traffic_manager_nested_endpoint#name TrafficManagerNestedEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/traffic_manager_nested_endpoint#value TrafficManagerNestedEndpoint#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

