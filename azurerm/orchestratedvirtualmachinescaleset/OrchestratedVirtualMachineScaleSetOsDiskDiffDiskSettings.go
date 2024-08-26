// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#option OrchestratedVirtualMachineScaleSet#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/orchestrated_virtual_machine_scale_set#placement OrchestratedVirtualMachineScaleSet#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}

