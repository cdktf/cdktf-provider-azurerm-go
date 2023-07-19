package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_spark_cluster#days HdinsightSparkCluster#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_spark_cluster#target_instance_count HdinsightSparkCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_spark_cluster#time HdinsightSparkCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

