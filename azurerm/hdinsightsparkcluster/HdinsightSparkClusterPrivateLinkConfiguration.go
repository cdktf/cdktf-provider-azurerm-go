// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterPrivateLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/hdinsight_spark_cluster#group_id HdinsightSparkCluster#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/hdinsight_spark_cluster#ip_configuration HdinsightSparkCluster#ip_configuration}
	IpConfiguration *HdinsightSparkClusterPrivateLinkConfigurationIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/hdinsight_spark_cluster#name HdinsightSparkCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

