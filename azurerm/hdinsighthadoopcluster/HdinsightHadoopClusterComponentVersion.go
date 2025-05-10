// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/hdinsight_hadoop_cluster#hadoop HdinsightHadoopCluster#hadoop}.
	Hadoop *string `field:"required" json:"hadoop" yaml:"hadoop"`
}

