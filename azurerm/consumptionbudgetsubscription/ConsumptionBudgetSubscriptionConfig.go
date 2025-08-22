// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConsumptionBudgetSubscriptionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#amount ConsumptionBudgetSubscription#amount}.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#name ConsumptionBudgetSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#notification ConsumptionBudgetSubscription#notification}
	Notification interface{} `field:"required" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#subscription_id ConsumptionBudgetSubscription#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// time_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#time_period ConsumptionBudgetSubscription#time_period}
	TimePeriod *ConsumptionBudgetSubscriptionTimePeriod `field:"required" json:"timePeriod" yaml:"timePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#etag ConsumptionBudgetSubscription#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#filter ConsumptionBudgetSubscription#filter}
	Filter *ConsumptionBudgetSubscriptionFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#id ConsumptionBudgetSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#time_grain ConsumptionBudgetSubscription#time_grain}.
	TimeGrain *string `field:"optional" json:"timeGrain" yaml:"timeGrain"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/consumption_budget_subscription#timeouts ConsumptionBudgetSubscription#timeouts}
	Timeouts *ConsumptionBudgetSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

