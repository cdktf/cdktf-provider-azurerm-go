// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extendedlocationcustomlocation


type ExtendedLocationCustomLocationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/extended_location_custom_location#create ExtendedLocationCustomLocation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/extended_location_custom_location#delete ExtendedLocationCustomLocation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/extended_location_custom_location#read ExtendedLocationCustomLocation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/extended_location_custom_location#update ExtendedLocationCustomLocation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

