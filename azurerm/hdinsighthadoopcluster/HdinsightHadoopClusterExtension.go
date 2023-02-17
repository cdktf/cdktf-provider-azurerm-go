package hdinsighthadoopcluster


type HdinsightHadoopClusterExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#log_analytics_workspace_id HdinsightHadoopCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#primary_key HdinsightHadoopCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

