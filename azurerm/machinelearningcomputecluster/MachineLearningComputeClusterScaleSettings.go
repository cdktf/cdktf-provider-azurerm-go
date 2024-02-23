// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningcomputecluster


type MachineLearningComputeClusterScaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/machine_learning_compute_cluster#max_node_count MachineLearningComputeCluster#max_node_count}.
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/machine_learning_compute_cluster#min_node_count MachineLearningComputeCluster#min_node_count}.
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/machine_learning_compute_cluster#scale_down_nodes_after_idle_duration MachineLearningComputeCluster#scale_down_nodes_after_idle_duration}.
	ScaleDownNodesAfterIdleDuration *string `field:"required" json:"scaleDownNodesAfterIdleDuration" yaml:"scaleDownNodesAfterIdleDuration"`
}

