// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolWindowsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/kubernetes_cluster_node_pool#outbound_nat_enabled KubernetesClusterNodePool#outbound_nat_enabled}.
	OutboundNatEnabled interface{} `field:"optional" json:"outboundNatEnabled" yaml:"outboundNatEnabled"`
}

