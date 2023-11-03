// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingaccountcostmanagementexport


type BillingAccountCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/billing_account_cost_management_export#time_frame BillingAccountCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/billing_account_cost_management_export#type BillingAccountCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

