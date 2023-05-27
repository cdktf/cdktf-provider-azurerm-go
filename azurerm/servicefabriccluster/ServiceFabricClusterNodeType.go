package servicefabriccluster


type ServiceFabricClusterNodeType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#client_endpoint_port ServiceFabricCluster#client_endpoint_port}.
	ClientEndpointPort *float64 `field:"required" json:"clientEndpointPort" yaml:"clientEndpointPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#http_endpoint_port ServiceFabricCluster#http_endpoint_port}.
	HttpEndpointPort *float64 `field:"required" json:"httpEndpointPort" yaml:"httpEndpointPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#instance_count ServiceFabricCluster#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#is_primary ServiceFabricCluster#is_primary}.
	IsPrimary interface{} `field:"required" json:"isPrimary" yaml:"isPrimary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#name ServiceFabricCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// application_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#application_ports ServiceFabricCluster#application_ports}
	ApplicationPorts *ServiceFabricClusterNodeTypeApplicationPorts `field:"optional" json:"applicationPorts" yaml:"applicationPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#capacities ServiceFabricCluster#capacities}.
	Capacities *map[string]*string `field:"optional" json:"capacities" yaml:"capacities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#durability_level ServiceFabricCluster#durability_level}.
	DurabilityLevel *string `field:"optional" json:"durabilityLevel" yaml:"durabilityLevel"`
	// ephemeral_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#ephemeral_ports ServiceFabricCluster#ephemeral_ports}
	EphemeralPorts *ServiceFabricClusterNodeTypeEphemeralPorts `field:"optional" json:"ephemeralPorts" yaml:"ephemeralPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#is_stateless ServiceFabricCluster#is_stateless}.
	IsStateless interface{} `field:"optional" json:"isStateless" yaml:"isStateless"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#multiple_availability_zones ServiceFabricCluster#multiple_availability_zones}.
	MultipleAvailabilityZones interface{} `field:"optional" json:"multipleAvailabilityZones" yaml:"multipleAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#placement_properties ServiceFabricCluster#placement_properties}.
	PlacementProperties *map[string]*string `field:"optional" json:"placementProperties" yaml:"placementProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/service_fabric_cluster#reverse_proxy_endpoint_port ServiceFabricCluster#reverse_proxy_endpoint_port}.
	ReverseProxyEndpointPort *float64 `field:"optional" json:"reverseProxyEndpointPort" yaml:"reverseProxyEndpointPort"`
}

