// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConsumptionBudgetManagementGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#amount ConsumptionBudgetManagementGroup#amount}.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#management_group_id ConsumptionBudgetManagementGroup#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#name ConsumptionBudgetManagementGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#notification ConsumptionBudgetManagementGroup#notification}
	Notification interface{} `field:"required" json:"notification" yaml:"notification"`
	// time_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#time_period ConsumptionBudgetManagementGroup#time_period}
	TimePeriod *ConsumptionBudgetManagementGroupTimePeriod `field:"required" json:"timePeriod" yaml:"timePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#etag ConsumptionBudgetManagementGroup#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#filter ConsumptionBudgetManagementGroup#filter}
	Filter *ConsumptionBudgetManagementGroupFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#id ConsumptionBudgetManagementGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#time_grain ConsumptionBudgetManagementGroup#time_grain}.
	TimeGrain *string `field:"optional" json:"timeGrain" yaml:"timeGrain"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/consumption_budget_management_group#timeouts ConsumptionBudgetManagementGroup#timeouts}
	Timeouts *ConsumptionBudgetManagementGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

