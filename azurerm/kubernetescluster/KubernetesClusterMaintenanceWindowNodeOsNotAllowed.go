// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterMaintenanceWindowNodeOsNotAllowed struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/kubernetes_cluster#end KubernetesCluster#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/kubernetes_cluster#start KubernetesCluster#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

