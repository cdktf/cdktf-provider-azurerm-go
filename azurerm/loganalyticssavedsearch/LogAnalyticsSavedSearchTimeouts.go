// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticssavedsearch


type LogAnalyticsSavedSearchTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/log_analytics_saved_search#create LogAnalyticsSavedSearch#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/log_analytics_saved_search#delete LogAnalyticsSavedSearch#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/log_analytics_saved_search#read LogAnalyticsSavedSearch#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

