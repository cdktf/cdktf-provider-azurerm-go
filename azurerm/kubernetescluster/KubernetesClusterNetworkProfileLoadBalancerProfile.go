// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterNetworkProfileLoadBalancerProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#idle_timeout_in_minutes KubernetesCluster#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#managed_outbound_ip_count KubernetesCluster#managed_outbound_ip_count}.
	ManagedOutboundIpCount *float64 `field:"optional" json:"managedOutboundIpCount" yaml:"managedOutboundIpCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#managed_outbound_ipv6_count KubernetesCluster#managed_outbound_ipv6_count}.
	ManagedOutboundIpv6Count *float64 `field:"optional" json:"managedOutboundIpv6Count" yaml:"managedOutboundIpv6Count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#outbound_ip_address_ids KubernetesCluster#outbound_ip_address_ids}.
	OutboundIpAddressIds *[]*string `field:"optional" json:"outboundIpAddressIds" yaml:"outboundIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#outbound_ip_prefix_ids KubernetesCluster#outbound_ip_prefix_ids}.
	OutboundIpPrefixIds *[]*string `field:"optional" json:"outboundIpPrefixIds" yaml:"outboundIpPrefixIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/kubernetes_cluster#outbound_ports_allocated KubernetesCluster#outbound_ports_allocated}.
	OutboundPortsAllocated *float64 `field:"optional" json:"outboundPortsAllocated" yaml:"outboundPortsAllocated"`
}

