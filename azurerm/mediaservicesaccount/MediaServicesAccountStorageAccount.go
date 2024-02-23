// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccount


type MediaServicesAccountStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_services_account#id MediaServicesAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_services_account#is_primary MediaServicesAccount#is_primary}.
	IsPrimary interface{} `field:"optional" json:"isPrimary" yaml:"isPrimary"`
	// managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/media_services_account#managed_identity MediaServicesAccount#managed_identity}
	ManagedIdentity *MediaServicesAccountStorageAccountManagedIdentity `field:"optional" json:"managedIdentity" yaml:"managedIdentity"`
}

