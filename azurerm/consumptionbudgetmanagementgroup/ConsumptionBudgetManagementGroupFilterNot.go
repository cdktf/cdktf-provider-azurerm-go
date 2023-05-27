package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/consumption_budget_management_group#dimension ConsumptionBudgetManagementGroup#dimension}
	Dimension *ConsumptionBudgetManagementGroupFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/consumption_budget_management_group#tag ConsumptionBudgetManagementGroup#tag}
	Tag *ConsumptionBudgetManagementGroupFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

