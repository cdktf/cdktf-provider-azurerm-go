// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#head_node HdinsightHadoopCluster#head_node}
	HeadNode *HdinsightHadoopClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#worker_node HdinsightHadoopCluster#worker_node}
	WorkerNode *HdinsightHadoopClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#zookeeper_node HdinsightHadoopCluster#zookeeper_node}
	ZookeeperNode *HdinsightHadoopClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
	// edge_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#edge_node HdinsightHadoopCluster#edge_node}
	EdgeNode *HdinsightHadoopClusterRolesEdgeNode `field:"optional" json:"edgeNode" yaml:"edgeNode"`
}

