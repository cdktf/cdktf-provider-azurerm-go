// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptioncostmanagementview


type SubscriptionCostManagementViewDataset struct {
	// aggregation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_cost_management_view#aggregation SubscriptionCostManagementView#aggregation}
	Aggregation interface{} `field:"required" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_cost_management_view#granularity SubscriptionCostManagementView#granularity}.
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_cost_management_view#grouping SubscriptionCostManagementView#grouping}
	Grouping interface{} `field:"optional" json:"grouping" yaml:"grouping"`
	// sorting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_cost_management_view#sorting SubscriptionCostManagementView#sorting}
	Sorting interface{} `field:"optional" json:"sorting" yaml:"sorting"`
}

