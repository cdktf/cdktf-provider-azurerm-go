package hdinsightkafkacluster


type HdinsightKafkaClusterRolesHeadNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_kafka_cluster#name HdinsightKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_kafka_cluster#uri HdinsightKafkaCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/hdinsight_kafka_cluster#parameters HdinsightKafkaCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

