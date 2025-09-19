// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/hdinsight_interactive_query_cluster#head_node HdinsightInteractiveQueryCluster#head_node}
	HeadNode *HdinsightInteractiveQueryClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/hdinsight_interactive_query_cluster#worker_node HdinsightInteractiveQueryCluster#worker_node}
	WorkerNode *HdinsightInteractiveQueryClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/hdinsight_interactive_query_cluster#zookeeper_node HdinsightInteractiveQueryCluster#zookeeper_node}
	ZookeeperNode *HdinsightInteractiveQueryClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

