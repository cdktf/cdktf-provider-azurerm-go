package hdinsightkafkacluster


type HdinsightKafkaClusterComputeIsolation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#compute_isolation_enabled HdinsightKafkaCluster#compute_isolation_enabled}.
	ComputeIsolationEnabled interface{} `field:"optional" json:"computeIsolationEnabled" yaml:"computeIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#host_sku HdinsightKafkaCluster#host_sku}.
	HostSku *string `field:"optional" json:"hostSku" yaml:"hostSku"`
}
