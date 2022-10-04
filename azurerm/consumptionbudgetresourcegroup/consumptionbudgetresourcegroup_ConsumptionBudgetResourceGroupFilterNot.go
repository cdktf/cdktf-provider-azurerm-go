package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_resource_group#dimension ConsumptionBudgetResourceGroup#dimension}
	Dimension *ConsumptionBudgetResourceGroupFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_resource_group#tag ConsumptionBudgetResourceGroup#tag}
	Tag *ConsumptionBudgetResourceGroupFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

