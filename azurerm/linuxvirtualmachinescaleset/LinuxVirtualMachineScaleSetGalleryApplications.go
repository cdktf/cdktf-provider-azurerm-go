// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetGalleryApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/linux_virtual_machine_scale_set#package_reference_id LinuxVirtualMachineScaleSet#package_reference_id}.
	PackageReferenceId *string `field:"required" json:"packageReferenceId" yaml:"packageReferenceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/linux_virtual_machine_scale_set#configuration_reference_blob_uri LinuxVirtualMachineScaleSet#configuration_reference_blob_uri}.
	ConfigurationReferenceBlobUri *string `field:"optional" json:"configurationReferenceBlobUri" yaml:"configurationReferenceBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/linux_virtual_machine_scale_set#order LinuxVirtualMachineScaleSet#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/linux_virtual_machine_scale_set#tag LinuxVirtualMachineScaleSet#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

