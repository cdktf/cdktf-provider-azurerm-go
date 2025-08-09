// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customprovider


type CustomProviderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/custom_provider#create CustomProvider#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/custom_provider#delete CustomProvider#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/custom_provider#read CustomProvider#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/custom_provider#update CustomProvider#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

