// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedlustrefilesystem


type ManagedLustreFileSystemIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/managed_lustre_file_system#identity_ids ManagedLustreFileSystem#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/managed_lustre_file_system#type ManagedLustreFileSystem#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

