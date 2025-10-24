// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetSkuProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/orchestrated_virtual_machine_scale_set#allocation_strategy OrchestratedVirtualMachineScaleSet#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/orchestrated_virtual_machine_scale_set#vm_sizes OrchestratedVirtualMachineScaleSet#vm_sizes}.
	VmSizes *[]*string `field:"required" json:"vmSizes" yaml:"vmSizes"`
}

