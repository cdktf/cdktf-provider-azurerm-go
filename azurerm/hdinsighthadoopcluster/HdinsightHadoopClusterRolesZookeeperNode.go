// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesZookeeperNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#vm_size HdinsightHadoopCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// script_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#script_actions HdinsightHadoopCluster#script_actions}
	ScriptActions interface{} `field:"optional" json:"scriptActions" yaml:"scriptActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#ssh_keys HdinsightHadoopCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#subnet_id HdinsightHadoopCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/hdinsight_hadoop_cluster#virtual_network_id HdinsightHadoopCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

