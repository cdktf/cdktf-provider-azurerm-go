// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptioncostmanagementexport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscriptionCostManagementExportConfig struct {
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
	// export_data_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#export_data_options SubscriptionCostManagementExport#export_data_options}
	ExportDataOptions *SubscriptionCostManagementExportExportDataOptions `field:"required" json:"exportDataOptions" yaml:"exportDataOptions"`
	// export_data_storage_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#export_data_storage_location SubscriptionCostManagementExport#export_data_storage_location}
	ExportDataStorageLocation *SubscriptionCostManagementExportExportDataStorageLocation `field:"required" json:"exportDataStorageLocation" yaml:"exportDataStorageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#name SubscriptionCostManagementExport#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#recurrence_period_end_date SubscriptionCostManagementExport#recurrence_period_end_date}.
	RecurrencePeriodEndDate *string `field:"required" json:"recurrencePeriodEndDate" yaml:"recurrencePeriodEndDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#recurrence_period_start_date SubscriptionCostManagementExport#recurrence_period_start_date}.
	RecurrencePeriodStartDate *string `field:"required" json:"recurrencePeriodStartDate" yaml:"recurrencePeriodStartDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#recurrence_type SubscriptionCostManagementExport#recurrence_type}.
	RecurrenceType *string `field:"required" json:"recurrenceType" yaml:"recurrenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#subscription_id SubscriptionCostManagementExport#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#active SubscriptionCostManagementExport#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#id SubscriptionCostManagementExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/subscription_cost_management_export#timeouts SubscriptionCostManagementExport#timeouts}
	Timeouts *SubscriptionCostManagementExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

