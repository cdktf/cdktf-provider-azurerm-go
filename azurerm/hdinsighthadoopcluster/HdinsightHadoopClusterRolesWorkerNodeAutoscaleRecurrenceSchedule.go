// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#days HdinsightHadoopCluster#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#target_instance_count HdinsightHadoopCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/hdinsight_hadoop_cluster#time HdinsightHadoopCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

