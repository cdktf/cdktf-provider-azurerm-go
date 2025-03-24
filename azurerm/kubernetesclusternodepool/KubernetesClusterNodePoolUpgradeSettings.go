// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster_node_pool#max_surge KubernetesClusterNodePool#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster_node_pool#drain_timeout_in_minutes KubernetesClusterNodePool#drain_timeout_in_minutes}.
	DrainTimeoutInMinutes *float64 `field:"optional" json:"drainTimeoutInMinutes" yaml:"drainTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/kubernetes_cluster_node_pool#node_soak_duration_in_minutes KubernetesClusterNodePool#node_soak_duration_in_minutes}.
	NodeSoakDurationInMinutes *float64 `field:"optional" json:"nodeSoakDurationInMinutes" yaml:"nodeSoakDurationInMinutes"`
}

