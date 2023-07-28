package subscriptioncostmanagementview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscriptionCostManagementViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#accumulated SubscriptionCostManagementView#accumulated}.
	Accumulated interface{} `field:"required" json:"accumulated" yaml:"accumulated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#chart_type SubscriptionCostManagementView#chart_type}.
	ChartType *string `field:"required" json:"chartType" yaml:"chartType"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#dataset SubscriptionCostManagementView#dataset}
	Dataset *SubscriptionCostManagementViewDataset `field:"required" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#display_name SubscriptionCostManagementView#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#name SubscriptionCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#report_type SubscriptionCostManagementView#report_type}.
	ReportType *string `field:"required" json:"reportType" yaml:"reportType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#subscription_id SubscriptionCostManagementView#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#timeframe SubscriptionCostManagementView#timeframe}.
	Timeframe *string `field:"required" json:"timeframe" yaml:"timeframe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#id SubscriptionCostManagementView#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kpi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#kpi SubscriptionCostManagementView#kpi}
	Kpi interface{} `field:"optional" json:"kpi" yaml:"kpi"`
	// pivot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#pivot SubscriptionCostManagementView#pivot}
	Pivot interface{} `field:"optional" json:"pivot" yaml:"pivot"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#timeouts SubscriptionCostManagementView#timeouts}
	Timeouts *SubscriptionCostManagementViewTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

