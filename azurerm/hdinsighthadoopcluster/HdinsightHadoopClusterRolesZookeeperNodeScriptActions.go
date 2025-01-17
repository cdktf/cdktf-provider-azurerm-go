// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesZookeeperNodeScriptActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hdinsight_hadoop_cluster#uri HdinsightHadoopCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hdinsight_hadoop_cluster#parameters HdinsightHadoopCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

