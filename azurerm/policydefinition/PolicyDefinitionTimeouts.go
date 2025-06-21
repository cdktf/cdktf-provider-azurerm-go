// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policydefinition


type PolicyDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/policy_definition#create PolicyDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/policy_definition#delete PolicyDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/policy_definition#read PolicyDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/policy_definition#update PolicyDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

