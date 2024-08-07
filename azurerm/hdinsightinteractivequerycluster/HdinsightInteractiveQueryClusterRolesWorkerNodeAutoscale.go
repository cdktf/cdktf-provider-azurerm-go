// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/hdinsight_interactive_query_cluster#capacity HdinsightInteractiveQueryCluster#capacity}
	Capacity *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/hdinsight_interactive_query_cluster#recurrence HdinsightInteractiveQueryCluster#recurrence}
	Recurrence *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

