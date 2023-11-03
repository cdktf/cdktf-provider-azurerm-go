// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/hdinsight_kafka_cluster#head_node HdinsightKafkaCluster#head_node}
	HeadNode *HdinsightKafkaClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/hdinsight_kafka_cluster#worker_node HdinsightKafkaCluster#worker_node}
	WorkerNode *HdinsightKafkaClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/hdinsight_kafka_cluster#zookeeper_node HdinsightKafkaCluster#zookeeper_node}
	ZookeeperNode *HdinsightKafkaClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
	// kafka_management_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/hdinsight_kafka_cluster#kafka_management_node HdinsightKafkaCluster#kafka_management_node}
	KafkaManagementNode *HdinsightKafkaClusterRolesKafkaManagementNode `field:"optional" json:"kafkaManagementNode" yaml:"kafkaManagementNode"`
}

