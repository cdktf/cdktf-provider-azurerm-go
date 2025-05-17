// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolNodeNetworkProfile struct {
	// allowed_host_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kubernetes_cluster_node_pool#allowed_host_ports KubernetesClusterNodePool#allowed_host_ports}
	AllowedHostPorts interface{} `field:"optional" json:"allowedHostPorts" yaml:"allowedHostPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kubernetes_cluster_node_pool#application_security_group_ids KubernetesClusterNodePool#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/kubernetes_cluster_node_pool#node_public_ip_tags KubernetesClusterNodePool#node_public_ip_tags}.
	NodePublicIpTags *map[string]*string `field:"optional" json:"nodePublicIpTags" yaml:"nodePublicIpTags"`
}

