// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oraclecloudvmcluster


type OracleCloudVmClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/oracle_cloud_vm_cluster#create OracleCloudVmCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/oracle_cloud_vm_cluster#delete OracleCloudVmCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/oracle_cloud_vm_cluster#read OracleCloudVmCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.32.0/docs/resources/oracle_cloud_vm_cluster#update OracleCloudVmCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

