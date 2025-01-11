// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package galleryapplicationversion


type GalleryApplicationVersionTargetRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/gallery_application_version#name GalleryApplicationVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/gallery_application_version#regional_replica_count GalleryApplicationVersion#regional_replica_count}.
	RegionalReplicaCount *float64 `field:"required" json:"regionalReplicaCount" yaml:"regionalReplicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/gallery_application_version#exclude_from_latest GalleryApplicationVersion#exclude_from_latest}.
	ExcludeFromLatest interface{} `field:"optional" json:"excludeFromLatest" yaml:"excludeFromLatest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/gallery_application_version#storage_account_type GalleryApplicationVersion#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
}

