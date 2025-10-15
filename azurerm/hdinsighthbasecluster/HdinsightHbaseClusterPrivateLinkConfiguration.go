// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster


type HdinsightHbaseClusterPrivateLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/hdinsight_hbase_cluster#group_id HdinsightHbaseCluster#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/hdinsight_hbase_cluster#ip_configuration HdinsightHbaseCluster#ip_configuration}
	IpConfiguration *HdinsightHbaseClusterPrivateLinkConfigurationIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/hdinsight_hbase_cluster#name HdinsightHbaseCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

