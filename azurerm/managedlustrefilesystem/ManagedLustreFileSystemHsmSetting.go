// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedlustrefilesystem


type ManagedLustreFileSystemHsmSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/managed_lustre_file_system#container_id ManagedLustreFileSystem#container_id}.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/managed_lustre_file_system#logging_container_id ManagedLustreFileSystem#logging_container_id}.
	LoggingContainerId *string `field:"required" json:"loggingContainerId" yaml:"loggingContainerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/managed_lustre_file_system#import_prefix ManagedLustreFileSystem#import_prefix}.
	ImportPrefix *string `field:"optional" json:"importPrefix" yaml:"importPrefix"`
}

