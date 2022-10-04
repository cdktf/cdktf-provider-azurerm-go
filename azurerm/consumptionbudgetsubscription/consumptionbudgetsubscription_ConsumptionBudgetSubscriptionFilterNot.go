package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_subscription#dimension ConsumptionBudgetSubscription#dimension}
	Dimension *ConsumptionBudgetSubscriptionFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_subscription#tag ConsumptionBudgetSubscription#tag}
	Tag *ConsumptionBudgetSubscriptionFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

