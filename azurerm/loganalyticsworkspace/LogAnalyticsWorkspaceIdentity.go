// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspace


type LogAnalyticsWorkspaceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/log_analytics_workspace#type LogAnalyticsWorkspace#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/log_analytics_workspace#identity_ids LogAnalyticsWorkspace#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

