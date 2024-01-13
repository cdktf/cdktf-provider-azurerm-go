// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralertprocessingruleactiongroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorAlertProcessingRuleActionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#add_action_group_ids MonitorAlertProcessingRuleActionGroup#add_action_group_ids}.
	AddActionGroupIds *[]*string `field:"required" json:"addActionGroupIds" yaml:"addActionGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#name MonitorAlertProcessingRuleActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#resource_group_name MonitorAlertProcessingRuleActionGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#scopes MonitorAlertProcessingRuleActionGroup#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#condition MonitorAlertProcessingRuleActionGroup#condition}
	Condition *MonitorAlertProcessingRuleActionGroupCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#description MonitorAlertProcessingRuleActionGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#enabled MonitorAlertProcessingRuleActionGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#id MonitorAlertProcessingRuleActionGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#schedule MonitorAlertProcessingRuleActionGroup#schedule}
	Schedule *MonitorAlertProcessingRuleActionGroupSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#tags MonitorAlertProcessingRuleActionGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/monitor_alert_processing_rule_action_group#timeouts MonitorAlertProcessingRuleActionGroup#timeouts}
	Timeouts *MonitorAlertProcessingRuleActionGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

