// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chaosstudioexperiment


type ChaosStudioExperimentSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/chaos_studio_experiment#chaos_studio_target_ids ChaosStudioExperiment#chaos_studio_target_ids}.
	ChaosStudioTargetIds *[]*string `field:"required" json:"chaosStudioTargetIds" yaml:"chaosStudioTargetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/chaos_studio_experiment#name ChaosStudioExperiment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

