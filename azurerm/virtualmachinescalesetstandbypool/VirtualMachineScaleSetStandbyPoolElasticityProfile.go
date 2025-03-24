// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetstandbypool


type VirtualMachineScaleSetStandbyPoolElasticityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_scale_set_standby_pool#max_ready_capacity VirtualMachineScaleSetStandbyPool#max_ready_capacity}.
	MaxReadyCapacity *float64 `field:"required" json:"maxReadyCapacity" yaml:"maxReadyCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_scale_set_standby_pool#min_ready_capacity VirtualMachineScaleSetStandbyPool#min_ready_capacity}.
	MinReadyCapacity *float64 `field:"required" json:"minReadyCapacity" yaml:"minReadyCapacity"`
}

