package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#log_analytics_workspace_id HdinsightInteractiveQueryCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#primary_key HdinsightInteractiveQueryCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

