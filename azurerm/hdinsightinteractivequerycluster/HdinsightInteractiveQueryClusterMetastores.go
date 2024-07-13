// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_interactive_query_cluster#ambari HdinsightInteractiveQueryCluster#ambari}
	Ambari *HdinsightInteractiveQueryClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_interactive_query_cluster#hive HdinsightInteractiveQueryCluster#hive}
	Hive *HdinsightInteractiveQueryClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_interactive_query_cluster#oozie HdinsightInteractiveQueryCluster#oozie}
	Oozie *HdinsightInteractiveQueryClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

