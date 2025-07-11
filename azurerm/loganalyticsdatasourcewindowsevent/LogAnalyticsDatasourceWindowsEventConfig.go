// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsdatasourcewindowsevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogAnalyticsDatasourceWindowsEventConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#event_log_name LogAnalyticsDatasourceWindowsEvent#event_log_name}.
	EventLogName *string `field:"required" json:"eventLogName" yaml:"eventLogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#event_types LogAnalyticsDatasourceWindowsEvent#event_types}.
	EventTypes *[]*string `field:"required" json:"eventTypes" yaml:"eventTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#name LogAnalyticsDatasourceWindowsEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#resource_group_name LogAnalyticsDatasourceWindowsEvent#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#workspace_name LogAnalyticsDatasourceWindowsEvent#workspace_name}.
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#id LogAnalyticsDatasourceWindowsEvent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/log_analytics_datasource_windows_event#timeouts LogAnalyticsDatasourceWindowsEvent#timeouts}
	Timeouts *LogAnalyticsDatasourceWindowsEventTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

