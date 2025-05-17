// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappactionhttp


type LogicAppActionHttpTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/logic_app_action_http#create LogicAppActionHttp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/logic_app_action_http#delete LogicAppActionHttp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/logic_app_action_http#read LogicAppActionHttp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/logic_app_action_http#update LogicAppActionHttp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

