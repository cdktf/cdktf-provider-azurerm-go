// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterNetworkProfileNatGatewayProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/kubernetes_cluster#idle_timeout_in_minutes KubernetesCluster#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/kubernetes_cluster#managed_outbound_ip_count KubernetesCluster#managed_outbound_ip_count}.
	ManagedOutboundIpCount *float64 `field:"optional" json:"managedOutboundIpCount" yaml:"managedOutboundIpCount"`
}

