package hdinsighthbasecluster


type HdinsightHbaseClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#schedule HdinsightHbaseCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#timezone HdinsightHbaseCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

