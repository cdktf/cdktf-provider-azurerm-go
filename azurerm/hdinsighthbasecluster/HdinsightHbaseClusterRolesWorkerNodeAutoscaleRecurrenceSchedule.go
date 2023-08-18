package hdinsighthbasecluster


type HdinsightHbaseClusterRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_hbase_cluster#days HdinsightHbaseCluster#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_hbase_cluster#target_instance_count HdinsightHbaseCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_hbase_cluster#time HdinsightHbaseCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

