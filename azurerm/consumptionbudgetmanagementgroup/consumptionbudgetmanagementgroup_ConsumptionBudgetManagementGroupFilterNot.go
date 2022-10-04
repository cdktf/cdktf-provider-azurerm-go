package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_management_group#dimension ConsumptionBudgetManagementGroup#dimension}
	Dimension *ConsumptionBudgetManagementGroupFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_management_group#tag ConsumptionBudgetManagementGroup#tag}
	Tag *ConsumptionBudgetManagementGroupFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

