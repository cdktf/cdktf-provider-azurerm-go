// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterComputeIsolation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_kafka_cluster#compute_isolation_enabled HdinsightKafkaCluster#compute_isolation_enabled}.
	ComputeIsolationEnabled interface{} `field:"optional" json:"computeIsolationEnabled" yaml:"computeIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/hdinsight_kafka_cluster#host_sku HdinsightKafkaCluster#host_sku}.
	HostSku *string `field:"optional" json:"hostSku" yaml:"hostSku"`
}

