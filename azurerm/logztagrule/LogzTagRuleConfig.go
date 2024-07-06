// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logztagrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogzTagRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#logz_monitor_id LogzTagRule#logz_monitor_id}.
	LogzMonitorId *string `field:"required" json:"logzMonitorId" yaml:"logzMonitorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#id LogzTagRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#send_aad_logs LogzTagRule#send_aad_logs}.
	SendAadLogs interface{} `field:"optional" json:"sendAadLogs" yaml:"sendAadLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#send_activity_logs LogzTagRule#send_activity_logs}.
	SendActivityLogs interface{} `field:"optional" json:"sendActivityLogs" yaml:"sendActivityLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#send_subscription_logs LogzTagRule#send_subscription_logs}.
	SendSubscriptionLogs interface{} `field:"optional" json:"sendSubscriptionLogs" yaml:"sendSubscriptionLogs"`
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#tag_filter LogzTagRule#tag_filter}
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/logz_tag_rule#timeouts LogzTagRule#timeouts}
	Timeouts *LogzTagRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

