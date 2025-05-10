// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterPrivateLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/hdinsight_kafka_cluster#group_id HdinsightKafkaCluster#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/hdinsight_kafka_cluster#ip_configuration HdinsightKafkaCluster#ip_configuration}
	IpConfiguration *HdinsightKafkaClusterPrivateLinkConfigurationIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/hdinsight_kafka_cluster#name HdinsightKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

