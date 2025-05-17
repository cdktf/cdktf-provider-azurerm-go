// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelloganalyticsworkspaceonboarding


type SentinelLogAnalyticsWorkspaceOnboardingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_log_analytics_workspace_onboarding#create SentinelLogAnalyticsWorkspaceOnboarding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_log_analytics_workspace_onboarding#delete SentinelLogAnalyticsWorkspaceOnboarding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/sentinel_log_analytics_workspace_onboarding#read SentinelLogAnalyticsWorkspaceOnboarding#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

