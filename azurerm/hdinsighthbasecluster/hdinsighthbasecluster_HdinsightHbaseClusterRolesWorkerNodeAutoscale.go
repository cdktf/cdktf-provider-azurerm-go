package hdinsighthbasecluster


type HdinsightHbaseClusterRolesWorkerNodeAutoscale struct {
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#recurrence HdinsightHbaseCluster#recurrence}
	Recurrence *HdinsightHbaseClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

