package hdinsighthadoopcluster


type HdinsightHadoopClusterGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}
