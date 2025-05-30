// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracemonitor


type DynatraceMonitorPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dynatrace_monitor#plan DynatraceMonitor#plan}.
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dynatrace_monitor#billing_cycle DynatraceMonitor#billing_cycle}.
	BillingCycle *string `field:"optional" json:"billingCycle" yaml:"billingCycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/dynatrace_monitor#usage_type DynatraceMonitor#usage_type}.
	UsageType *string `field:"optional" json:"usageType" yaml:"usageType"`
}

