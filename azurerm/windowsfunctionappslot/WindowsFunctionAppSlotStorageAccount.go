// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#access_key WindowsFunctionAppSlot#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#account_name WindowsFunctionAppSlot#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#name WindowsFunctionAppSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#share_name WindowsFunctionAppSlot#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#type WindowsFunctionAppSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/windows_function_app_slot#mount_path WindowsFunctionAppSlot#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

