// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package galleryapplication


type GalleryApplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/gallery_application#create GalleryApplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/gallery_application#delete GalleryApplication#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/gallery_application#read GalleryApplication#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/gallery_application#update GalleryApplication#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

