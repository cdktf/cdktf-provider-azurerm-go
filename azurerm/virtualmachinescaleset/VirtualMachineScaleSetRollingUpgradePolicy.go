// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetRollingUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#max_batch_instance_percent VirtualMachineScaleSet#max_batch_instance_percent}.
	MaxBatchInstancePercent *float64 `field:"optional" json:"maxBatchInstancePercent" yaml:"maxBatchInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#max_unhealthy_instance_percent VirtualMachineScaleSet#max_unhealthy_instance_percent}.
	MaxUnhealthyInstancePercent *float64 `field:"optional" json:"maxUnhealthyInstancePercent" yaml:"maxUnhealthyInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#max_unhealthy_upgraded_instance_percent VirtualMachineScaleSet#max_unhealthy_upgraded_instance_percent}.
	MaxUnhealthyUpgradedInstancePercent *float64 `field:"optional" json:"maxUnhealthyUpgradedInstancePercent" yaml:"maxUnhealthyUpgradedInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/virtual_machine_scale_set#pause_time_between_batches VirtualMachineScaleSet#pause_time_between_batches}.
	PauseTimeBetweenBatches *string `field:"optional" json:"pauseTimeBetweenBatches" yaml:"pauseTimeBetweenBatches"`
}

