// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageblobinventorypolicy


type StorageBlobInventoryPolicyRulesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#blob_types StorageBlobInventoryPolicy#blob_types}.
	BlobTypes *[]*string `field:"required" json:"blobTypes" yaml:"blobTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#exclude_prefixes StorageBlobInventoryPolicy#exclude_prefixes}.
	ExcludePrefixes *[]*string `field:"optional" json:"excludePrefixes" yaml:"excludePrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#include_blob_versions StorageBlobInventoryPolicy#include_blob_versions}.
	IncludeBlobVersions interface{} `field:"optional" json:"includeBlobVersions" yaml:"includeBlobVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#include_deleted StorageBlobInventoryPolicy#include_deleted}.
	IncludeDeleted interface{} `field:"optional" json:"includeDeleted" yaml:"includeDeleted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#include_snapshots StorageBlobInventoryPolicy#include_snapshots}.
	IncludeSnapshots interface{} `field:"optional" json:"includeSnapshots" yaml:"includeSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/storage_blob_inventory_policy#prefix_match StorageBlobInventoryPolicy#prefix_match}.
	PrefixMatch *[]*string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
}

