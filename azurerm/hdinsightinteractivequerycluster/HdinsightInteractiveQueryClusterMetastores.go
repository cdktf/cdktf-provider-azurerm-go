package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ambari HdinsightInteractiveQueryCluster#ambari}
	Ambari *HdinsightInteractiveQueryClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#hive HdinsightInteractiveQueryCluster#hive}
	Hive *HdinsightInteractiveQueryClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#oozie HdinsightInteractiveQueryCluster#oozie}
	Oozie *HdinsightInteractiveQueryClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}
