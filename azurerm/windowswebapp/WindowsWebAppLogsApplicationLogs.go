// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp


type WindowsWebAppLogsApplicationLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#file_system_level WindowsWebApp#file_system_level}.
	FileSystemLevel *string `field:"required" json:"fileSystemLevel" yaml:"fileSystemLevel"`
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/windows_web_app#azure_blob_storage WindowsWebApp#azure_blob_storage}
	AzureBlobStorage *WindowsWebAppLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
}

