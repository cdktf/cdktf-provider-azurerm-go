// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprometheusrulegroup


type MonitorAlertPrometheusRuleGroupRuleAlertResolution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_alert_prometheus_rule_group#auto_resolved MonitorAlertPrometheusRuleGroup#auto_resolved}.
	AutoResolved interface{} `field:"optional" json:"autoResolved" yaml:"autoResolved"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/monitor_alert_prometheus_rule_group#time_to_resolve MonitorAlertPrometheusRuleGroup#time_to_resolve}.
	TimeToResolve *string `field:"optional" json:"timeToResolve" yaml:"timeToResolve"`
}

