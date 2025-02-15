// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprometheusrulegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertPrometheusRuleGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#location MonitorAlertPrometheusRuleGroup#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#name MonitorAlertPrometheusRuleGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#resource_group_name MonitorAlertPrometheusRuleGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#rule MonitorAlertPrometheusRuleGroup#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#scopes MonitorAlertPrometheusRuleGroup#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#cluster_name MonitorAlertPrometheusRuleGroup#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#description MonitorAlertPrometheusRuleGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#id MonitorAlertPrometheusRuleGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#interval MonitorAlertPrometheusRuleGroup#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#rule_group_enabled MonitorAlertPrometheusRuleGroup#rule_group_enabled}.
	RuleGroupEnabled interface{} `field:"optional" json:"ruleGroupEnabled" yaml:"ruleGroupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#tags MonitorAlertPrometheusRuleGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/monitor_alert_prometheus_rule_group#timeouts MonitorAlertPrometheusRuleGroup#timeouts}
	Timeouts *MonitorAlertPrometheusRuleGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

