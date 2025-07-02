// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracetagrules


type DynatraceTagRulesLogRule struct {
	// filtering_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/dynatrace_tag_rules#filtering_tag DynatraceTagRules#filtering_tag}
	FilteringTag interface{} `field:"required" json:"filteringTag" yaml:"filteringTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/dynatrace_tag_rules#send_activity_logs_enabled DynatraceTagRules#send_activity_logs_enabled}.
	SendActivityLogsEnabled interface{} `field:"optional" json:"sendActivityLogsEnabled" yaml:"sendActivityLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/dynatrace_tag_rules#send_azure_active_directory_logs_enabled DynatraceTagRules#send_azure_active_directory_logs_enabled}.
	SendAzureActiveDirectoryLogsEnabled interface{} `field:"optional" json:"sendAzureActiveDirectoryLogsEnabled" yaml:"sendAzureActiveDirectoryLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/dynatrace_tag_rules#send_subscription_logs_enabled DynatraceTagRules#send_subscription_logs_enabled}.
	SendSubscriptionLogsEnabled interface{} `field:"optional" json:"sendSubscriptionLogsEnabled" yaml:"sendSubscriptionLogsEnabled"`
}

