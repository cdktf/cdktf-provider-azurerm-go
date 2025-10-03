// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterPrivateLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/hdinsight_interactive_query_cluster#group_id HdinsightInteractiveQueryCluster#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/hdinsight_interactive_query_cluster#ip_configuration HdinsightInteractiveQueryCluster#ip_configuration}
	IpConfiguration *HdinsightInteractiveQueryClusterPrivateLinkConfigurationIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/hdinsight_interactive_query_cluster#name HdinsightInteractiveQueryCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

