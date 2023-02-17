package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#schedule HdinsightInteractiveQueryCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#timezone HdinsightInteractiveQueryCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

