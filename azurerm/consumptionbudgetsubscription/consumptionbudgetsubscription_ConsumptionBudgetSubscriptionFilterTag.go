package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilterTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_subscription#name ConsumptionBudgetSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_subscription#values ConsumptionBudgetSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_subscription#operator ConsumptionBudgetSubscription#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}
