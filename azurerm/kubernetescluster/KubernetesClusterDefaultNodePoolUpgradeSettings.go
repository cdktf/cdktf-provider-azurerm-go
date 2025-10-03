// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/kubernetes_cluster#max_surge KubernetesCluster#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/kubernetes_cluster#drain_timeout_in_minutes KubernetesCluster#drain_timeout_in_minutes}.
	DrainTimeoutInMinutes *float64 `field:"optional" json:"drainTimeoutInMinutes" yaml:"drainTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/kubernetes_cluster#node_soak_duration_in_minutes KubernetesCluster#node_soak_duration_in_minutes}.
	NodeSoakDurationInMinutes *float64 `field:"optional" json:"nodeSoakDurationInMinutes" yaml:"nodeSoakDurationInMinutes"`
}

