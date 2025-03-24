// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chaosstudioexperiment


type ChaosStudioExperimentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/chaos_studio_experiment#create ChaosStudioExperiment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/chaos_studio_experiment#delete ChaosStudioExperiment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/chaos_studio_experiment#read ChaosStudioExperiment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/chaos_studio_experiment#update ChaosStudioExperiment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

