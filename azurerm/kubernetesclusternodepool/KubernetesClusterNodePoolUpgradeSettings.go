// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/kubernetes_cluster_node_pool#max_surge KubernetesClusterNodePool#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

