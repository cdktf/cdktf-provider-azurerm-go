package hdinsightsparkcluster


type HdinsightSparkClusterRolesZookeeperNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/hdinsight_spark_cluster#name HdinsightSparkCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/hdinsight_spark_cluster#uri HdinsightSparkCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/hdinsight_spark_cluster#parameters HdinsightSparkCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

