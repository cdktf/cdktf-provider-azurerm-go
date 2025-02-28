// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsanalyticsitem


type ApplicationInsightsAnalyticsItemTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/application_insights_analytics_item#create ApplicationInsightsAnalyticsItem#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/application_insights_analytics_item#delete ApplicationInsightsAnalyticsItem#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/application_insights_analytics_item#read ApplicationInsightsAnalyticsItem#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/application_insights_analytics_item#update ApplicationInsightsAnalyticsItem#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

