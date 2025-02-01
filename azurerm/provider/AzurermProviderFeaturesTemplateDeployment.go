// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesTemplateDeployment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs#delete_nested_items_during_deletion AzurermProvider#delete_nested_items_during_deletion}.
	DeleteNestedItemsDuringDeletion interface{} `field:"required" json:"deleteNestedItemsDuringDeletion" yaml:"deleteNestedItemsDuringDeletion"`
}

