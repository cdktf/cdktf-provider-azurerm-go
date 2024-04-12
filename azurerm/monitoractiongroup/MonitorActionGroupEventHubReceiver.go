// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractiongroup


type MonitorActionGroupEventHubReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#event_hub_id MonitorActionGroup#event_hub_id}.
	EventHubId *string `field:"optional" json:"eventHubId" yaml:"eventHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#event_hub_name MonitorActionGroup#event_hub_name}.
	EventHubName *string `field:"optional" json:"eventHubName" yaml:"eventHubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#event_hub_namespace MonitorActionGroup#event_hub_namespace}.
	EventHubNamespace *string `field:"optional" json:"eventHubNamespace" yaml:"eventHubNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#subscription_id MonitorActionGroup#subscription_id}.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#tenant_id MonitorActionGroup#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_action_group#use_common_alert_schema MonitorActionGroup#use_common_alert_schema}.
	UseCommonAlertSchema interface{} `field:"optional" json:"useCommonAlertSchema" yaml:"useCommonAlertSchema"`
}

