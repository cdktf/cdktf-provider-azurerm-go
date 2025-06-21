// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/hdinsight_spark_cluster#capacity HdinsightSparkCluster#capacity}
	Capacity *HdinsightSparkClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/hdinsight_spark_cluster#recurrence HdinsightSparkCluster#recurrence}
	Recurrence *HdinsightSparkClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

