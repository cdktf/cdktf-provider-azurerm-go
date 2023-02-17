package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilterNotDimension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_management_group#name ConsumptionBudgetManagementGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_management_group#values ConsumptionBudgetManagementGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_management_group#operator ConsumptionBudgetManagementGroup#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

