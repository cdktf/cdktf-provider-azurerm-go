// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetAdditionalUnattendContent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#content WindowsVirtualMachineScaleSet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/windows_virtual_machine_scale_set#setting WindowsVirtualMachineScaleSet#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
}

