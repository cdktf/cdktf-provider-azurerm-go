// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimagegallery


type SharedImageGallerySharing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/shared_image_gallery#permission SharedImageGallery#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// community_gallery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/shared_image_gallery#community_gallery SharedImageGallery#community_gallery}
	CommunityGallery *SharedImageGallerySharingCommunityGallery `field:"optional" json:"communityGallery" yaml:"communityGallery"`
}

