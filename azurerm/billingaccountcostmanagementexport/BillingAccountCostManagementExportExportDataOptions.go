package billingaccountcostmanagementexport


type BillingAccountCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/billing_account_cost_management_export#time_frame BillingAccountCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/billing_account_cost_management_export#type BillingAccountCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}
