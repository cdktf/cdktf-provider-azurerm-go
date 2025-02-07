// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetGalleryApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_virtual_machine_scale_set#version_id LinuxVirtualMachineScaleSet#version_id}.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_virtual_machine_scale_set#configuration_blob_uri LinuxVirtualMachineScaleSet#configuration_blob_uri}.
	ConfigurationBlobUri *string `field:"optional" json:"configurationBlobUri" yaml:"configurationBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_virtual_machine_scale_set#order LinuxVirtualMachineScaleSet#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/linux_virtual_machine_scale_set#tag LinuxVirtualMachineScaleSet#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

