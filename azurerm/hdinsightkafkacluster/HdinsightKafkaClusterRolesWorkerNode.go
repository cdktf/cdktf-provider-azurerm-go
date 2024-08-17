// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster


type HdinsightKafkaClusterRolesWorkerNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#number_of_disks_per_node HdinsightKafkaCluster#number_of_disks_per_node}.
	NumberOfDisksPerNode *float64 `field:"required" json:"numberOfDisksPerNode" yaml:"numberOfDisksPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#target_instance_count HdinsightKafkaCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#username HdinsightKafkaCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#vm_size HdinsightKafkaCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#password HdinsightKafkaCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// script_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#script_actions HdinsightKafkaCluster#script_actions}
	ScriptActions interface{} `field:"optional" json:"scriptActions" yaml:"scriptActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#ssh_keys HdinsightKafkaCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#subnet_id HdinsightKafkaCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/hdinsight_kafka_cluster#virtual_network_id HdinsightKafkaCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

