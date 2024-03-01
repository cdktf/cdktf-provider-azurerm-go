// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolNodeNetworkProfile struct {
	// allowed_host_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#allowed_host_ports KubernetesCluster#allowed_host_ports}
	AllowedHostPorts interface{} `field:"optional" json:"allowedHostPorts" yaml:"allowedHostPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#application_security_group_ids KubernetesCluster#application_security_group_ids}.
	ApplicationSecurityGroupIds *[]*string `field:"optional" json:"applicationSecurityGroupIds" yaml:"applicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/kubernetes_cluster#node_public_ip_tags KubernetesCluster#node_public_ip_tags}.
	NodePublicIpTags *map[string]*string `field:"optional" json:"nodePublicIpTags" yaml:"nodePublicIpTags"`
}

