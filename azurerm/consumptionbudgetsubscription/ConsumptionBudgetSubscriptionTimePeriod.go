package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionTimePeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/consumption_budget_subscription#start_date ConsumptionBudgetSubscription#start_date}.
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/consumption_budget_subscription#end_date ConsumptionBudgetSubscription#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
}

