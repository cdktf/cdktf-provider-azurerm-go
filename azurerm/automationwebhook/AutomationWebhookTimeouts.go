// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationwebhook


type AutomationWebhookTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/automation_webhook#create AutomationWebhook#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/automation_webhook#delete AutomationWebhook#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/automation_webhook#read AutomationWebhook#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/automation_webhook#update AutomationWebhook#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

