package hdinsighthbasecluster


type HdinsightHbaseClusterRolesZookeeperNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#name HdinsightHbaseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#uri HdinsightHbaseCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#parameters HdinsightHbaseCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

