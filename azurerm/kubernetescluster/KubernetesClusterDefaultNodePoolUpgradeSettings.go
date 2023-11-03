// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kubernetes_cluster#max_surge KubernetesCluster#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

