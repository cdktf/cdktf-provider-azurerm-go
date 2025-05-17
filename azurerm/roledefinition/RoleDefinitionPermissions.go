// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package roledefinition


type RoleDefinitionPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/role_definition#actions RoleDefinition#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/role_definition#data_actions RoleDefinition#data_actions}.
	DataActions *[]*string `field:"optional" json:"dataActions" yaml:"dataActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/role_definition#not_actions RoleDefinition#not_actions}.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/role_definition#not_data_actions RoleDefinition#not_data_actions}.
	NotDataActions *[]*string `field:"optional" json:"notDataActions" yaml:"notDataActions"`
}

