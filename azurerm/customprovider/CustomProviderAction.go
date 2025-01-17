// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customprovider


type CustomProviderAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/custom_provider#endpoint CustomProvider#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/custom_provider#name CustomProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

