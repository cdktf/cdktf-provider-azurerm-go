// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chaosstudioexperiment


type ChaosStudioExperimentSteps struct {
	// branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/chaos_studio_experiment#branch ChaosStudioExperiment#branch}
	Branch interface{} `field:"required" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/chaos_studio_experiment#name ChaosStudioExperiment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

