// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/hdinsight_hadoop_cluster#max_instance_count HdinsightHadoopCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/hdinsight_hadoop_cluster#min_instance_count HdinsightHadoopCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

