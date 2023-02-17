package hdinsightsparkcluster


type HdinsightSparkClusterGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#password HdinsightSparkCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#username HdinsightSparkCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

