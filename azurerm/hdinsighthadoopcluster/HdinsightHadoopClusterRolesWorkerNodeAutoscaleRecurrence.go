package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#schedule HdinsightHadoopCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#timezone HdinsightHadoopCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

