// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogAnalyticsWorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#location LogAnalyticsWorkspace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#name LogAnalyticsWorkspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#resource_group_name LogAnalyticsWorkspace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#allow_resource_only_permissions LogAnalyticsWorkspace#allow_resource_only_permissions}.
	AllowResourceOnlyPermissions interface{} `field:"optional" json:"allowResourceOnlyPermissions" yaml:"allowResourceOnlyPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#cmk_for_query_forced LogAnalyticsWorkspace#cmk_for_query_forced}.
	CmkForQueryForced interface{} `field:"optional" json:"cmkForQueryForced" yaml:"cmkForQueryForced"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#daily_quota_gb LogAnalyticsWorkspace#daily_quota_gb}.
	DailyQuotaGb *float64 `field:"optional" json:"dailyQuotaGb" yaml:"dailyQuotaGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#data_collection_rule_id LogAnalyticsWorkspace#data_collection_rule_id}.
	DataCollectionRuleId *string `field:"optional" json:"dataCollectionRuleId" yaml:"dataCollectionRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#id LogAnalyticsWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#identity LogAnalyticsWorkspace#identity}
	Identity *LogAnalyticsWorkspaceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#immediate_data_purge_on_30_days_enabled LogAnalyticsWorkspace#immediate_data_purge_on_30_days_enabled}.
	ImmediateDataPurgeOn30DaysEnabled interface{} `field:"optional" json:"immediateDataPurgeOn30DaysEnabled" yaml:"immediateDataPurgeOn30DaysEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#internet_ingestion_enabled LogAnalyticsWorkspace#internet_ingestion_enabled}.
	InternetIngestionEnabled interface{} `field:"optional" json:"internetIngestionEnabled" yaml:"internetIngestionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#internet_query_enabled LogAnalyticsWorkspace#internet_query_enabled}.
	InternetQueryEnabled interface{} `field:"optional" json:"internetQueryEnabled" yaml:"internetQueryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#local_authentication_disabled LogAnalyticsWorkspace#local_authentication_disabled}.
	LocalAuthenticationDisabled interface{} `field:"optional" json:"localAuthenticationDisabled" yaml:"localAuthenticationDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#reservation_capacity_in_gb_per_day LogAnalyticsWorkspace#reservation_capacity_in_gb_per_day}.
	ReservationCapacityInGbPerDay *float64 `field:"optional" json:"reservationCapacityInGbPerDay" yaml:"reservationCapacityInGbPerDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#retention_in_days LogAnalyticsWorkspace#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#sku LogAnalyticsWorkspace#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#tags LogAnalyticsWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/log_analytics_workspace#timeouts LogAnalyticsWorkspace#timeouts}
	Timeouts *LogAnalyticsWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

