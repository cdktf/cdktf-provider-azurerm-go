// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractivitylogalert


type MonitorActivityLogAlertCriteriaServiceHealth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/monitor_activity_log_alert#events MonitorActivityLogAlert#events}.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/monitor_activity_log_alert#locations MonitorActivityLogAlert#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/monitor_activity_log_alert#services MonitorActivityLogAlert#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

