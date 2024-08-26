// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfleetupdaterun


type KubernetesFleetUpdateRunStage struct {
	// group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_fleet_update_run#group KubernetesFleetUpdateRun#group}
	Group interface{} `field:"required" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_fleet_update_run#name KubernetesFleetUpdateRun#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/kubernetes_fleet_update_run#after_stage_wait_in_seconds KubernetesFleetUpdateRun#after_stage_wait_in_seconds}.
	AfterStageWaitInSeconds *float64 `field:"optional" json:"afterStageWaitInSeconds" yaml:"afterStageWaitInSeconds"`
}

