// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package newrelicmonitor


type NewRelicMonitorPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/new_relic_monitor#effective_date NewRelicMonitor#effective_date}.
	EffectiveDate *string `field:"required" json:"effectiveDate" yaml:"effectiveDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/new_relic_monitor#billing_cycle NewRelicMonitor#billing_cycle}.
	BillingCycle *string `field:"optional" json:"billingCycle" yaml:"billingCycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/new_relic_monitor#plan_id NewRelicMonitor#plan_id}.
	PlanId *string `field:"optional" json:"planId" yaml:"planId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/new_relic_monitor#usage_type NewRelicMonitor#usage_type}.
	UsageType *string `field:"optional" json:"usageType" yaml:"usageType"`
}

