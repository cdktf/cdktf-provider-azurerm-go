// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprometheusrulegroup


type MonitorAlertPrometheusRuleGroupRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#expression MonitorAlertPrometheusRuleGroup#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#action MonitorAlertPrometheusRuleGroup#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#alert MonitorAlertPrometheusRuleGroup#alert}.
	Alert *string `field:"optional" json:"alert" yaml:"alert"`
	// alert_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#alert_resolution MonitorAlertPrometheusRuleGroup#alert_resolution}
	AlertResolution *MonitorAlertPrometheusRuleGroupRuleAlertResolution `field:"optional" json:"alertResolution" yaml:"alertResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#annotations MonitorAlertPrometheusRuleGroup#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#enabled MonitorAlertPrometheusRuleGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#for MonitorAlertPrometheusRuleGroup#for}.
	For *string `field:"optional" json:"for" yaml:"for"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#labels MonitorAlertPrometheusRuleGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#record MonitorAlertPrometheusRuleGroup#record}.
	Record *string `field:"optional" json:"record" yaml:"record"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/monitor_alert_prometheus_rule_group#severity MonitorAlertPrometheusRuleGroup#severity}.
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
}

