// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingaccountcostmanagementexport


type BillingAccountCostManagementExportExportDataStorageLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/billing_account_cost_management_export#container_id BillingAccountCostManagementExport#container_id}.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/billing_account_cost_management_export#root_folder_path BillingAccountCostManagementExport#root_folder_path}.
	RootFolderPath *string `field:"required" json:"rootFolderPath" yaml:"rootFolderPath"`
}

