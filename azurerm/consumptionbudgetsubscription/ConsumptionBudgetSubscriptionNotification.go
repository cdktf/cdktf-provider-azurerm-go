package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#operator ConsumptionBudgetSubscription#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#threshold ConsumptionBudgetSubscription#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#contact_emails ConsumptionBudgetSubscription#contact_emails}.
	ContactEmails *[]*string `field:"optional" json:"contactEmails" yaml:"contactEmails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#contact_groups ConsumptionBudgetSubscription#contact_groups}.
	ContactGroups *[]*string `field:"optional" json:"contactGroups" yaml:"contactGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#contact_roles ConsumptionBudgetSubscription#contact_roles}.
	ContactRoles *[]*string `field:"optional" json:"contactRoles" yaml:"contactRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#enabled ConsumptionBudgetSubscription#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/consumption_budget_subscription#threshold_type ConsumptionBudgetSubscription#threshold_type}.
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
}

