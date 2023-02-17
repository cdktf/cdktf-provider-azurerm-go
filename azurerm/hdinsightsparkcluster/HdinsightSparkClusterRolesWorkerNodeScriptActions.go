package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#name HdinsightSparkCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#uri HdinsightSparkCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#parameters HdinsightSparkCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

