// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolNodeNetworkProfileAllowedHostPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/kubernetes_cluster#port_end KubernetesCluster#port_end}.
	PortEnd *float64 `field:"optional" json:"portEnd" yaml:"portEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/kubernetes_cluster#port_start KubernetesCluster#port_start}.
	PortStart *float64 `field:"optional" json:"portStart" yaml:"portStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/kubernetes_cluster#protocol KubernetesCluster#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

