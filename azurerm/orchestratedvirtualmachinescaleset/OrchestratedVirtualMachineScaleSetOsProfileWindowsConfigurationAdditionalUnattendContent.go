// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationAdditionalUnattendContent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/orchestrated_virtual_machine_scale_set#content OrchestratedVirtualMachineScaleSet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/orchestrated_virtual_machine_scale_set#setting OrchestratedVirtualMachineScaleSet#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
}

