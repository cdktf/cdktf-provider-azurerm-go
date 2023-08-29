// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#access_key LinuxFunctionApp#access_key}.
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#account_name LinuxFunctionApp#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#name LinuxFunctionApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#share_name LinuxFunctionApp#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#type LinuxFunctionApp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/linux_function_app#mount_path LinuxFunctionApp#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

