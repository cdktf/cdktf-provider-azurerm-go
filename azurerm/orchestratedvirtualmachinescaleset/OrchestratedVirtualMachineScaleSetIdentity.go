// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/orchestrated_virtual_machine_scale_set#identity_ids OrchestratedVirtualMachineScaleSet#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/orchestrated_virtual_machine_scale_set#type OrchestratedVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

