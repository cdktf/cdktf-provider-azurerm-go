// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterComputeIsolation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#compute_isolation_enabled HdinsightInteractiveQueryCluster#compute_isolation_enabled}.
	ComputeIsolationEnabled interface{} `field:"optional" json:"computeIsolationEnabled" yaml:"computeIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/hdinsight_interactive_query_cluster#host_sku HdinsightInteractiveQueryCluster#host_sku}.
	HostSku *string `field:"optional" json:"hostSku" yaml:"hostSku"`
}

