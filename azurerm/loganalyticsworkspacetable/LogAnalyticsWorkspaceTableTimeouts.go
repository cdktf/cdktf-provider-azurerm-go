// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetable


type LogAnalyticsWorkspaceTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/log_analytics_workspace_table#create LogAnalyticsWorkspaceTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/log_analytics_workspace_table#delete LogAnalyticsWorkspaceTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/log_analytics_workspace_table#read LogAnalyticsWorkspaceTable#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/log_analytics_workspace_table#update LogAnalyticsWorkspaceTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

