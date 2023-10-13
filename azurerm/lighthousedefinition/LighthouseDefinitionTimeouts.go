// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lighthousedefinition


type LighthouseDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/lighthouse_definition#create LighthouseDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/lighthouse_definition#delete LighthouseDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/lighthouse_definition#read LighthouseDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/lighthouse_definition#update LighthouseDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

