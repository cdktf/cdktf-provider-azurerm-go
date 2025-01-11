// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileLinuxConfigSshKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#path VirtualMachineScaleSet#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#key_data VirtualMachineScaleSet#key_data}.
	KeyData *string `field:"optional" json:"keyData" yaml:"keyData"`
}

