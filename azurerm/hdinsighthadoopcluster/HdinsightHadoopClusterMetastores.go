// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/hdinsight_hadoop_cluster#ambari HdinsightHadoopCluster#ambari}
	Ambari *HdinsightHadoopClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/hdinsight_hadoop_cluster#hive HdinsightHadoopCluster#hive}
	Hive *HdinsightHadoopClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/hdinsight_hadoop_cluster#oozie HdinsightHadoopCluster#oozie}
	Oozie *HdinsightHadoopClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

