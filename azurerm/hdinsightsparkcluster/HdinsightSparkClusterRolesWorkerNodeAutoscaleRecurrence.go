package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/hdinsight_spark_cluster#schedule HdinsightSparkCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/hdinsight_spark_cluster#timezone HdinsightSparkCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

