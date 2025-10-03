// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesAppConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}.
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs#recover_soft_deleted AzurermProvider#recover_soft_deleted}.
	RecoverSoftDeleted interface{} `field:"optional" json:"recoverSoftDeleted" yaml:"recoverSoftDeleted"`
}

