// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimagegallery


type SharedImageGallerySharingCommunityGallery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/shared_image_gallery#eula SharedImageGallery#eula}.
	Eula *string `field:"required" json:"eula" yaml:"eula"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/shared_image_gallery#prefix SharedImageGallery#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/shared_image_gallery#publisher_email SharedImageGallery#publisher_email}.
	PublisherEmail *string `field:"required" json:"publisherEmail" yaml:"publisherEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/shared_image_gallery#publisher_uri SharedImageGallery#publisher_uri}.
	PublisherUri *string `field:"required" json:"publisherUri" yaml:"publisherUri"`
}

