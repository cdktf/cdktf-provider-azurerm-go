package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterRolesZookeeperNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#name HdinsightInteractiveQueryCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#uri HdinsightInteractiveQueryCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#parameters HdinsightInteractiveQueryCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}
