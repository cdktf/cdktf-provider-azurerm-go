// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesCognitiveAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}.
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
}

