// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicapptriggerhttprequest


type LogicAppTriggerHttpRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_trigger_http_request#create LogicAppTriggerHttpRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_trigger_http_request#delete LogicAppTriggerHttpRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_trigger_http_request#read LogicAppTriggerHttpRequest#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/logic_app_trigger_http_request#update LogicAppTriggerHttpRequest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

