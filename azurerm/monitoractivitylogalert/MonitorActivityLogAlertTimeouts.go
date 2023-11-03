// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractivitylogalert


type MonitorActivityLogAlertTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_activity_log_alert#create MonitorActivityLogAlert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_activity_log_alert#delete MonitorActivityLogAlert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_activity_log_alert#read MonitorActivityLogAlert#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/monitor_activity_log_alert#update MonitorActivityLogAlert#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

