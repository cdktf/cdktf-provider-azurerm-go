package hdinsightsparkcluster


type HdinsightSparkClusterMetastoresHive struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hdinsight_spark_cluster#database_name HdinsightSparkCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hdinsight_spark_cluster#password HdinsightSparkCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hdinsight_spark_cluster#server HdinsightSparkCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/hdinsight_spark_cluster#username HdinsightSparkCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

