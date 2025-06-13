// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetRollingUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#max_batch_instance_percent OrchestratedVirtualMachineScaleSet#max_batch_instance_percent}.
	MaxBatchInstancePercent *float64 `field:"required" json:"maxBatchInstancePercent" yaml:"maxBatchInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#max_unhealthy_instance_percent OrchestratedVirtualMachineScaleSet#max_unhealthy_instance_percent}.
	MaxUnhealthyInstancePercent *float64 `field:"required" json:"maxUnhealthyInstancePercent" yaml:"maxUnhealthyInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#max_unhealthy_upgraded_instance_percent OrchestratedVirtualMachineScaleSet#max_unhealthy_upgraded_instance_percent}.
	MaxUnhealthyUpgradedInstancePercent *float64 `field:"required" json:"maxUnhealthyUpgradedInstancePercent" yaml:"maxUnhealthyUpgradedInstancePercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#pause_time_between_batches OrchestratedVirtualMachineScaleSet#pause_time_between_batches}.
	PauseTimeBetweenBatches *string `field:"required" json:"pauseTimeBetweenBatches" yaml:"pauseTimeBetweenBatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#cross_zone_upgrades_enabled OrchestratedVirtualMachineScaleSet#cross_zone_upgrades_enabled}.
	CrossZoneUpgradesEnabled interface{} `field:"optional" json:"crossZoneUpgradesEnabled" yaml:"crossZoneUpgradesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#maximum_surge_instances_enabled OrchestratedVirtualMachineScaleSet#maximum_surge_instances_enabled}.
	MaximumSurgeInstancesEnabled interface{} `field:"optional" json:"maximumSurgeInstancesEnabled" yaml:"maximumSurgeInstancesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/orchestrated_virtual_machine_scale_set#prioritize_unhealthy_instances_enabled OrchestratedVirtualMachineScaleSet#prioritize_unhealthy_instances_enabled}.
	PrioritizeUnhealthyInstancesEnabled interface{} `field:"optional" json:"prioritizeUnhealthyInstancesEnabled" yaml:"prioritizeUnhealthyInstancesEnabled"`
}

