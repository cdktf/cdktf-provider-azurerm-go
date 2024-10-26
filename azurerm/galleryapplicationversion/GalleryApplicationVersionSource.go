// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package galleryapplicationversion


type GalleryApplicationVersionSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/gallery_application_version#media_link GalleryApplicationVersion#media_link}.
	MediaLink *string `field:"required" json:"mediaLink" yaml:"mediaLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/gallery_application_version#default_configuration_link GalleryApplicationVersion#default_configuration_link}.
	DefaultConfigurationLink *string `field:"optional" json:"defaultConfigurationLink" yaml:"defaultConfigurationLink"`
}

