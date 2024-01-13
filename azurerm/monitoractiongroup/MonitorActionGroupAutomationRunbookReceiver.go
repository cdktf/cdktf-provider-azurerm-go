// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractiongroup


type MonitorActionGroupAutomationRunbookReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#automation_account_id MonitorActionGroup#automation_account_id}.
	AutomationAccountId *string `field:"required" json:"automationAccountId" yaml:"automationAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#is_global_runbook MonitorActionGroup#is_global_runbook}.
	IsGlobalRunbook interface{} `field:"required" json:"isGlobalRunbook" yaml:"isGlobalRunbook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#runbook_name MonitorActionGroup#runbook_name}.
	RunbookName *string `field:"required" json:"runbookName" yaml:"runbookName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#service_uri MonitorActionGroup#service_uri}.
	ServiceUri *string `field:"required" json:"serviceUri" yaml:"serviceUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#webhook_resource_id MonitorActionGroup#webhook_resource_id}.
	WebhookResourceId *string `field:"required" json:"webhookResourceId" yaml:"webhookResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_action_group#use_common_alert_schema MonitorActionGroup#use_common_alert_schema}.
	UseCommonAlertSchema interface{} `field:"optional" json:"useCommonAlertSchema" yaml:"useCommonAlertSchema"`
}

