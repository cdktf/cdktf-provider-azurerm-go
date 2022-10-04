package kustocluster


type KustoClusterSku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#name KustoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#capacity KustoCluster#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

