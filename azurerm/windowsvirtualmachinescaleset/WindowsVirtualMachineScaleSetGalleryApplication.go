// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetGalleryApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/windows_virtual_machine_scale_set#version_id WindowsVirtualMachineScaleSet#version_id}.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/windows_virtual_machine_scale_set#configuration_blob_uri WindowsVirtualMachineScaleSet#configuration_blob_uri}.
	ConfigurationBlobUri *string `field:"optional" json:"configurationBlobUri" yaml:"configurationBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/windows_virtual_machine_scale_set#order WindowsVirtualMachineScaleSet#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/windows_virtual_machine_scale_set#tag WindowsVirtualMachineScaleSet#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

