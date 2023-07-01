package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/hdinsight_spark_cluster#capacity HdinsightSparkCluster#capacity}
	Capacity *HdinsightSparkClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/hdinsight_spark_cluster#recurrence HdinsightSparkCluster#recurrence}
	Recurrence *HdinsightSparkClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

