// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppLogsApplicationLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/linux_web_app#file_system_level LinuxWebApp#file_system_level}.
	FileSystemLevel *string `field:"required" json:"fileSystemLevel" yaml:"fileSystemLevel"`
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/linux_web_app#azure_blob_storage LinuxWebApp#azure_blob_storage}
	AzureBlobStorage *LinuxWebAppLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
}

