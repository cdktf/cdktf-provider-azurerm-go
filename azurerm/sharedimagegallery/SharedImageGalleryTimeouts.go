// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimagegallery


type SharedImageGalleryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_gallery#create SharedImageGallery#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_gallery#delete SharedImageGallery#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_gallery#read SharedImageGallery#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/shared_image_gallery#update SharedImageGallery#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

