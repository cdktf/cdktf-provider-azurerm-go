package hdinsighthadoopcluster


type HdinsightHadoopClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#hadoop HdinsightHadoopCluster#hadoop}.
	Hadoop *string `field:"required" json:"hadoop" yaml:"hadoop"`
}

