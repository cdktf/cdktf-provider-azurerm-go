package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/consumption_budget_resource_group#dimension ConsumptionBudgetResourceGroup#dimension}
	Dimension *ConsumptionBudgetResourceGroupFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/consumption_budget_resource_group#tag ConsumptionBudgetResourceGroup#tag}
	Tag *ConsumptionBudgetResourceGroupFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

