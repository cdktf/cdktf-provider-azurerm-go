// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitivedeployment


type CognitiveDeploymentModel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/cognitive_deployment#format CognitiveDeployment#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/cognitive_deployment#name CognitiveDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/cognitive_deployment#version CognitiveDeployment#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

