// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policysetdefinition


type PolicySetDefinitionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/policy_set_definition#create PolicySetDefinition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/policy_set_definition#delete PolicySetDefinition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/policy_set_definition#read PolicySetDefinition#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/policy_set_definition#update PolicySetDefinition#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

