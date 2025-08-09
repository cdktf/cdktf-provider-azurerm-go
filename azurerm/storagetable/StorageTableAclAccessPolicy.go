// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetable


type StorageTableAclAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/storage_table#expiry StorageTable#expiry}.
	Expiry *string `field:"required" json:"expiry" yaml:"expiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/storage_table#permissions StorageTable#permissions}.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/storage_table#start StorageTable#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

