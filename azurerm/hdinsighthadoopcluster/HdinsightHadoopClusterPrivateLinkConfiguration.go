// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterPrivateLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_hadoop_cluster#group_id HdinsightHadoopCluster#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_hadoop_cluster#ip_configuration HdinsightHadoopCluster#ip_configuration}
	IpConfiguration *HdinsightHadoopClusterPrivateLinkConfigurationIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

