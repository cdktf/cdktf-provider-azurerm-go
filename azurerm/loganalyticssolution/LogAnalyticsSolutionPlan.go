// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticssolution


type LogAnalyticsSolutionPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/log_analytics_solution#product LogAnalyticsSolution#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/log_analytics_solution#publisher LogAnalyticsSolution#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/log_analytics_solution#promotion_code LogAnalyticsSolution#promotion_code}.
	PromotionCode *string `field:"optional" json:"promotionCode" yaml:"promotionCode"`
}

