package hdinsightkafkacluster


type HdinsightKafkaClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_kafka_cluster#create HdinsightKafkaCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_kafka_cluster#delete HdinsightKafkaCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_kafka_cluster#read HdinsightKafkaCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/hdinsight_kafka_cluster#update HdinsightKafkaCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

