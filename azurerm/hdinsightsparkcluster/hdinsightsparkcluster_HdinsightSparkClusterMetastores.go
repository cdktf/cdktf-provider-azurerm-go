package hdinsightsparkcluster


type HdinsightSparkClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#ambari HdinsightSparkCluster#ambari}
	Ambari *HdinsightSparkClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#hive HdinsightSparkCluster#hive}
	Hive *HdinsightSparkClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_spark_cluster#oozie HdinsightSparkCluster#oozie}
	Oozie *HdinsightSparkClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

