// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chaosstudioexperiment


type ChaosStudioExperimentStepsBranchActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/chaos_studio_experiment#action_type ChaosStudioExperiment#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/chaos_studio_experiment#duration ChaosStudioExperiment#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/chaos_studio_experiment#parameters ChaosStudioExperiment#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/chaos_studio_experiment#selector_name ChaosStudioExperiment#selector_name}.
	SelectorName *string `field:"optional" json:"selectorName" yaml:"selectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/chaos_studio_experiment#urn ChaosStudioExperiment#urn}.
	Urn *string `field:"optional" json:"urn" yaml:"urn"`
}

