package hdinsightsparkcluster


type HdinsightSparkClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#spark HdinsightSparkCluster#spark}.
	Spark *string `field:"required" json:"spark" yaml:"spark"`
}

