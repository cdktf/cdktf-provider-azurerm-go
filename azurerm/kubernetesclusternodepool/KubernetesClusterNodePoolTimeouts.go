// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_node_pool#create KubernetesClusterNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_node_pool#delete KubernetesClusterNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_node_pool#read KubernetesClusterNodePool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_node_pool#update KubernetesClusterNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

