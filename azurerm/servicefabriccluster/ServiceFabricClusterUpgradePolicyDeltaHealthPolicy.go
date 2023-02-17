package servicefabriccluster


type ServiceFabricClusterUpgradePolicyDeltaHealthPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_delta_unhealthy_applications_percent ServiceFabricCluster#max_delta_unhealthy_applications_percent}.
	MaxDeltaUnhealthyApplicationsPercent *float64 `field:"optional" json:"maxDeltaUnhealthyApplicationsPercent" yaml:"maxDeltaUnhealthyApplicationsPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_delta_unhealthy_nodes_percent ServiceFabricCluster#max_delta_unhealthy_nodes_percent}.
	MaxDeltaUnhealthyNodesPercent *float64 `field:"optional" json:"maxDeltaUnhealthyNodesPercent" yaml:"maxDeltaUnhealthyNodesPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_upgrade_domain_delta_unhealthy_nodes_percent ServiceFabricCluster#max_upgrade_domain_delta_unhealthy_nodes_percent}.
	MaxUpgradeDomainDeltaUnhealthyNodesPercent *float64 `field:"optional" json:"maxUpgradeDomainDeltaUnhealthyNodesPercent" yaml:"maxUpgradeDomainDeltaUnhealthyNodesPercent"`
}

