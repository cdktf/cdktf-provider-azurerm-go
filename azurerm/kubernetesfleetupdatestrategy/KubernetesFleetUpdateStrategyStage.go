// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdatestrategy


type KubernetesFleetUpdateStrategyStage struct {
	// group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/kubernetes_fleet_update_strategy#group KubernetesFleetUpdateStrategy#group}
	Group interface{} `field:"required" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/kubernetes_fleet_update_strategy#name KubernetesFleetUpdateStrategy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/kubernetes_fleet_update_strategy#after_stage_wait_in_seconds KubernetesFleetUpdateStrategy#after_stage_wait_in_seconds}.
	AfterStageWaitInSeconds *float64 `field:"optional" json:"afterStageWaitInSeconds" yaml:"afterStageWaitInSeconds"`
}

