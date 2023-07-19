package servicefabricmanagedcluster


type ServiceFabricManagedClusterLbRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/service_fabric_managed_cluster#backend_port ServiceFabricManagedCluster#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/service_fabric_managed_cluster#frontend_port ServiceFabricManagedCluster#frontend_port}.
	FrontendPort *float64 `field:"required" json:"frontendPort" yaml:"frontendPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/service_fabric_managed_cluster#probe_protocol ServiceFabricManagedCluster#probe_protocol}.
	ProbeProtocol *string `field:"required" json:"probeProtocol" yaml:"probeProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/service_fabric_managed_cluster#protocol ServiceFabricManagedCluster#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/service_fabric_managed_cluster#probe_request_path ServiceFabricManagedCluster#probe_request_path}.
	ProbeRequestPath *string `field:"optional" json:"probeRequestPath" yaml:"probeRequestPath"`
}

