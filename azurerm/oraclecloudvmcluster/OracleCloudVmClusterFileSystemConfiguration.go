// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oraclecloudvmcluster


type OracleCloudVmClusterFileSystemConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_cloud_vm_cluster#mount_point OracleCloudVmCluster#mount_point}.
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/oracle_cloud_vm_cluster#size_in_gb OracleCloudVmCluster#size_in_gb}.
	SizeInGb *float64 `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
}

