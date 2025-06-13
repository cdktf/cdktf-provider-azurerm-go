// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#api_management_name ApiManagementSubscription#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#display_name ApiManagementSubscription#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#resource_group_name ApiManagementSubscription#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#allow_tracing ApiManagementSubscription#allow_tracing}.
	AllowTracing interface{} `field:"optional" json:"allowTracing" yaml:"allowTracing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#api_id ApiManagementSubscription#api_id}.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#id ApiManagementSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#primary_key ApiManagementSubscription#primary_key}.
	PrimaryKey *string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#product_id ApiManagementSubscription#product_id}.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#secondary_key ApiManagementSubscription#secondary_key}.
	SecondaryKey *string `field:"optional" json:"secondaryKey" yaml:"secondaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#state ApiManagementSubscription#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#subscription_id ApiManagementSubscription#subscription_id}.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#timeouts ApiManagementSubscription#timeouts}
	Timeouts *ApiManagementSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_subscription#user_id ApiManagementSubscription#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

