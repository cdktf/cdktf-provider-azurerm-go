package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupFilterNotDimension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_resource_group#name ConsumptionBudgetResourceGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_resource_group#values ConsumptionBudgetResourceGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/consumption_budget_resource_group#operator ConsumptionBudgetResourceGroup#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}
