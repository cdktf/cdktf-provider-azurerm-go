package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#max_instance_count HdinsightInteractiveQueryCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#min_instance_count HdinsightInteractiveQueryCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

