// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package newrelictagrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NewRelicTagRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#monitor_id NewRelicTagRule#monitor_id}.
	MonitorId *string `field:"required" json:"monitorId" yaml:"monitorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#activity_log_enabled NewRelicTagRule#activity_log_enabled}.
	ActivityLogEnabled interface{} `field:"optional" json:"activityLogEnabled" yaml:"activityLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#azure_active_directory_log_enabled NewRelicTagRule#azure_active_directory_log_enabled}.
	AzureActiveDirectoryLogEnabled interface{} `field:"optional" json:"azureActiveDirectoryLogEnabled" yaml:"azureActiveDirectoryLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#id NewRelicTagRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#log_tag_filter NewRelicTagRule#log_tag_filter}
	LogTagFilter interface{} `field:"optional" json:"logTagFilter" yaml:"logTagFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#metric_enabled NewRelicTagRule#metric_enabled}.
	MetricEnabled interface{} `field:"optional" json:"metricEnabled" yaml:"metricEnabled"`
	// metric_tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#metric_tag_filter NewRelicTagRule#metric_tag_filter}
	MetricTagFilter interface{} `field:"optional" json:"metricTagFilter" yaml:"metricTagFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#subscription_log_enabled NewRelicTagRule#subscription_log_enabled}.
	SubscriptionLogEnabled interface{} `field:"optional" json:"subscriptionLogEnabled" yaml:"subscriptionLogEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/new_relic_tag_rule#timeouts NewRelicTagRule#timeouts}
	Timeouts *NewRelicTagRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

