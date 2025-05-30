// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsstorageinsights


type LogAnalyticsStorageInsightsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_storage_insights#create LogAnalyticsStorageInsights#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_storage_insights#delete LogAnalyticsStorageInsights#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_storage_insights#read LogAnalyticsStorageInsights#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/log_analytics_storage_insights#update LogAnalyticsStorageInsights#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

