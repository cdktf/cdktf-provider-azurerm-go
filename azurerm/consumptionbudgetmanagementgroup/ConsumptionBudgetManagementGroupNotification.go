package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/consumption_budget_management_group#contact_emails ConsumptionBudgetManagementGroup#contact_emails}.
	ContactEmails *[]*string `field:"required" json:"contactEmails" yaml:"contactEmails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/consumption_budget_management_group#operator ConsumptionBudgetManagementGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/consumption_budget_management_group#threshold ConsumptionBudgetManagementGroup#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/consumption_budget_management_group#enabled ConsumptionBudgetManagementGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/consumption_budget_management_group#threshold_type ConsumptionBudgetManagementGroup#threshold_type}.
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
}

