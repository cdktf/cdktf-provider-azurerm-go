package datasharedatasetkustocluster


type DataShareDatasetKustoClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_kusto_cluster#create DataShareDatasetKustoCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_kusto_cluster#delete DataShareDatasetKustoCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share_dataset_kusto_cluster#read DataShareDatasetKustoCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
