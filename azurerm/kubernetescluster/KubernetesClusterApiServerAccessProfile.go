// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterApiServerAccessProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/kubernetes_cluster#authorized_ip_ranges KubernetesCluster#authorized_ip_ranges}.
	AuthorizedIpRanges *[]*string `field:"optional" json:"authorizedIpRanges" yaml:"authorizedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/kubernetes_cluster#subnet_id KubernetesCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/kubernetes_cluster#virtual_network_integration_enabled KubernetesCluster#virtual_network_integration_enabled}.
	VirtualNetworkIntegrationEnabled interface{} `field:"optional" json:"virtualNetworkIntegrationEnabled" yaml:"virtualNetworkIntegrationEnabled"`
}

