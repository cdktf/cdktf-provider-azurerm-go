// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesLogAnalyticsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.12.0/docs#permanently_delete_on_destroy AzurermProvider#permanently_delete_on_destroy}.
	PermanentlyDeleteOnDestroy interface{} `field:"optional" json:"permanentlyDeleteOnDestroy" yaml:"permanentlyDeleteOnDestroy"`
}

