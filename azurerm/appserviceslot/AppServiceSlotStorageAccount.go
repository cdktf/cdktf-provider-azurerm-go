// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceslot


type AppServiceSlotStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#access_key AppServiceSlot#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#account_name AppServiceSlot#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#name AppServiceSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#share_name AppServiceSlot#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#type AppServiceSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/app_service_slot#mount_path AppServiceSlot#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

