// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprometheusrulegroup


type MonitorAlertPrometheusRuleGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/monitor_alert_prometheus_rule_group#create MonitorAlertPrometheusRuleGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/monitor_alert_prometheus_rule_group#delete MonitorAlertPrometheusRuleGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/monitor_alert_prometheus_rule_group#read MonitorAlertPrometheusRuleGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/monitor_alert_prometheus_rule_group#update MonitorAlertPrometheusRuleGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

