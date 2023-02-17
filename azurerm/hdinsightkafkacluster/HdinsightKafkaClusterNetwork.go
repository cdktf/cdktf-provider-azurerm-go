package hdinsightkafkacluster


type HdinsightKafkaClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#connection_direction HdinsightKafkaCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#private_link_enabled HdinsightKafkaCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}

