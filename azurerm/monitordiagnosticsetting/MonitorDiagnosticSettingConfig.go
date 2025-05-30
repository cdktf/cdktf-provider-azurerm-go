// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordiagnosticsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDiagnosticSettingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#name MonitorDiagnosticSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#target_resource_id MonitorDiagnosticSetting#target_resource_id}.
	TargetResourceId *string `field:"required" json:"targetResourceId" yaml:"targetResourceId"`
	// enabled_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#enabled_log MonitorDiagnosticSetting#enabled_log}
	EnabledLog interface{} `field:"optional" json:"enabledLog" yaml:"enabledLog"`
	// enabled_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#enabled_metric MonitorDiagnosticSetting#enabled_metric}
	EnabledMetric interface{} `field:"optional" json:"enabledMetric" yaml:"enabledMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#eventhub_authorization_rule_id MonitorDiagnosticSetting#eventhub_authorization_rule_id}.
	EventhubAuthorizationRuleId *string `field:"optional" json:"eventhubAuthorizationRuleId" yaml:"eventhubAuthorizationRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#eventhub_name MonitorDiagnosticSetting#eventhub_name}.
	EventhubName *string `field:"optional" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#id MonitorDiagnosticSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#log_analytics_destination_type MonitorDiagnosticSetting#log_analytics_destination_type}.
	LogAnalyticsDestinationType *string `field:"optional" json:"logAnalyticsDestinationType" yaml:"logAnalyticsDestinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#log_analytics_workspace_id MonitorDiagnosticSetting#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"optional" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#metric MonitorDiagnosticSetting#metric}
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#partner_solution_id MonitorDiagnosticSetting#partner_solution_id}.
	PartnerSolutionId *string `field:"optional" json:"partnerSolutionId" yaml:"partnerSolutionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#storage_account_id MonitorDiagnosticSetting#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/monitor_diagnostic_setting#timeouts MonitorDiagnosticSetting#timeouts}
	Timeouts *MonitorDiagnosticSettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

