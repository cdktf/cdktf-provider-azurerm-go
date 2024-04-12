// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionappslot


type LinuxFunctionAppSlotStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#access_key LinuxFunctionAppSlot#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#account_name LinuxFunctionAppSlot#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#name LinuxFunctionAppSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#share_name LinuxFunctionAppSlot#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#type LinuxFunctionAppSlot#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/linux_function_app_slot#mount_path LinuxFunctionAppSlot#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

