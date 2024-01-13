// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterComputeIsolation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/hdinsight_spark_cluster#compute_isolation_enabled HdinsightSparkCluster#compute_isolation_enabled}.
	ComputeIsolationEnabled interface{} `field:"optional" json:"computeIsolationEnabled" yaml:"computeIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/hdinsight_spark_cluster#host_sku HdinsightSparkCluster#host_sku}.
	HostSku *string `field:"optional" json:"hostSku" yaml:"hostSku"`
}

