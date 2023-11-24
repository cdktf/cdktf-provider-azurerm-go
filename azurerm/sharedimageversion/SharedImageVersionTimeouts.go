// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion


type SharedImageVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_version#create SharedImageVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_version#delete SharedImageVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_version#read SharedImageVersion#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_version#update SharedImageVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

