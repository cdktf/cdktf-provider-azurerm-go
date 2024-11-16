// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthbasecluster


type HdinsightHbaseClusterRolesWorkerNodeAutoscale struct {
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/hdinsight_hbase_cluster#recurrence HdinsightHbaseCluster#recurrence}
	Recurrence *HdinsightHbaseClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

