// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedlustrefilesystem


type ManagedLustreFileSystemRootSquash struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/managed_lustre_file_system#mode ManagedLustreFileSystem#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/managed_lustre_file_system#no_squash_nids ManagedLustreFileSystem#no_squash_nids}.
	NoSquashNids *string `field:"required" json:"noSquashNids" yaml:"noSquashNids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/managed_lustre_file_system#squash_gid ManagedLustreFileSystem#squash_gid}.
	SquashGid *float64 `field:"optional" json:"squashGid" yaml:"squashGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/managed_lustre_file_system#squash_uid ManagedLustreFileSystem#squash_uid}.
	SquashUid *float64 `field:"optional" json:"squashUid" yaml:"squashUid"`
}

