// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetGalleryApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/windows_virtual_machine_scale_set#package_reference_id WindowsVirtualMachineScaleSet#package_reference_id}.
	PackageReferenceId *string `field:"required" json:"packageReferenceId" yaml:"packageReferenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/windows_virtual_machine_scale_set#configuration_reference_blob_uri WindowsVirtualMachineScaleSet#configuration_reference_blob_uri}.
	ConfigurationReferenceBlobUri *string `field:"optional" json:"configurationReferenceBlobUri" yaml:"configurationReferenceBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/windows_virtual_machine_scale_set#order WindowsVirtualMachineScaleSet#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/windows_virtual_machine_scale_set#tag WindowsVirtualMachineScaleSet#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

