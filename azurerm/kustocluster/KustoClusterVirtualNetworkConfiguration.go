package kustocluster


type KustoClusterVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_cluster#data_management_public_ip_id KustoCluster#data_management_public_ip_id}.
	DataManagementPublicIpId *string `field:"required" json:"dataManagementPublicIpId" yaml:"dataManagementPublicIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_cluster#engine_public_ip_id KustoCluster#engine_public_ip_id}.
	EnginePublicIpId *string `field:"required" json:"enginePublicIpId" yaml:"enginePublicIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/kusto_cluster#subnet_id KustoCluster#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

