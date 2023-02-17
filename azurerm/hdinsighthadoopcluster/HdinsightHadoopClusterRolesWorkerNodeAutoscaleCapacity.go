package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#max_instance_count HdinsightHadoopCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#min_instance_count HdinsightHadoopCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

