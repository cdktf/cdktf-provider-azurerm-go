// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightsparkcluster


type HdinsightSparkClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/hdinsight_spark_cluster#spark HdinsightSparkCluster#spark}.
	Spark *string `field:"required" json:"spark" yaml:"spark"`
}

