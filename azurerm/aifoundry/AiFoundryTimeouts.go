// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aifoundry


type AiFoundryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/ai_foundry#create AiFoundry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/ai_foundry#delete AiFoundry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/ai_foundry#read AiFoundry#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/ai_foundry#update AiFoundry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

