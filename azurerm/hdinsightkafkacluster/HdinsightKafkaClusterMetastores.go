// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/hdinsight_kafka_cluster#ambari HdinsightKafkaCluster#ambari}
	Ambari *HdinsightKafkaClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/hdinsight_kafka_cluster#hive HdinsightKafkaCluster#hive}
	Hive *HdinsightKafkaClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/hdinsight_kafka_cluster#oozie HdinsightKafkaCluster#oozie}
	Oozie *HdinsightKafkaClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

