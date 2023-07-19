package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilterNotDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/consumption_budget_subscription#name ConsumptionBudgetSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/consumption_budget_subscription#values ConsumptionBudgetSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/consumption_budget_subscription#operator ConsumptionBudgetSubscription#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

