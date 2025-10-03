// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package galleryapplicationversion


type GalleryApplicationVersionManageAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/gallery_application_version#install GalleryApplicationVersion#install}.
	Install *string `field:"required" json:"install" yaml:"install"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/gallery_application_version#remove GalleryApplicationVersion#remove}.
	Remove *string `field:"required" json:"remove" yaml:"remove"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/gallery_application_version#update GalleryApplicationVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

