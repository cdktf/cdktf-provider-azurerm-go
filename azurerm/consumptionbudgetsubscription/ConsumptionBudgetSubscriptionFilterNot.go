package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/consumption_budget_subscription#dimension ConsumptionBudgetSubscription#dimension}
	Dimension *ConsumptionBudgetSubscriptionFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/consumption_budget_subscription#tag ConsumptionBudgetSubscription#tag}
	Tag *ConsumptionBudgetSubscriptionFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

