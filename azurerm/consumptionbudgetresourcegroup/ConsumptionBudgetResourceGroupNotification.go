// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#operator ConsumptionBudgetResourceGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#threshold ConsumptionBudgetResourceGroup#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#contact_emails ConsumptionBudgetResourceGroup#contact_emails}.
	ContactEmails *[]*string `field:"optional" json:"contactEmails" yaml:"contactEmails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#contact_groups ConsumptionBudgetResourceGroup#contact_groups}.
	ContactGroups *[]*string `field:"optional" json:"contactGroups" yaml:"contactGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#contact_roles ConsumptionBudgetResourceGroup#contact_roles}.
	ContactRoles *[]*string `field:"optional" json:"contactRoles" yaml:"contactRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#enabled ConsumptionBudgetResourceGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/consumption_budget_resource_group#threshold_type ConsumptionBudgetResourceGroup#threshold_type}.
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
}

