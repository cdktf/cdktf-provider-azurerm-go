// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsdatasourcewindowsevent


type LogAnalyticsDatasourceWindowsEventTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/log_analytics_datasource_windows_event#create LogAnalyticsDatasourceWindowsEvent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/log_analytics_datasource_windows_event#delete LogAnalyticsDatasourceWindowsEvent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/log_analytics_datasource_windows_event#read LogAnalyticsDatasourceWindowsEvent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/log_analytics_datasource_windows_event#update LogAnalyticsDatasourceWindowsEvent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

