// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#head_node HdinsightSparkCluster#head_node}
	HeadNode *HdinsightSparkClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#worker_node HdinsightSparkCluster#worker_node}
	WorkerNode *HdinsightSparkClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/hdinsight_spark_cluster#zookeeper_node HdinsightSparkCluster#zookeeper_node}
	ZookeeperNode *HdinsightSparkClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

