// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceproviderregistration


type ResourceProviderRegistrationFeature struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/resource_provider_registration#name ResourceProviderRegistration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/resource_provider_registration#registered ResourceProviderRegistration#registered}.
	Registered interface{} `field:"required" json:"registered" yaml:"registered"`
}

