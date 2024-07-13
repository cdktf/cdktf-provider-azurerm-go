// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractiongroup


type MonitorActionGroupArmRoleReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_group#role_id MonitorActionGroup#role_id}.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/monitor_action_group#use_common_alert_schema MonitorActionGroup#use_common_alert_schema}.
	UseCommonAlertSchema interface{} `field:"optional" json:"useCommonAlertSchema" yaml:"useCommonAlertSchema"`
}

