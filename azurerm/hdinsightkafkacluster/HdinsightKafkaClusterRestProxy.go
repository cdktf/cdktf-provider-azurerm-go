package hdinsightkafkacluster


type HdinsightKafkaClusterRestProxy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#security_group_id HdinsightKafkaCluster#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_kafka_cluster#security_group_name HdinsightKafkaCluster#security_group_name}.
	SecurityGroupName *string `field:"required" json:"securityGroupName" yaml:"securityGroupName"`
}
