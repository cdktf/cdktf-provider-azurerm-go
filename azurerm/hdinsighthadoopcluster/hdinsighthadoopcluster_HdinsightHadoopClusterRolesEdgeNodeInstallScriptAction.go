package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesEdgeNodeInstallScriptAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#uri HdinsightHadoopCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#parameters HdinsightHadoopCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

