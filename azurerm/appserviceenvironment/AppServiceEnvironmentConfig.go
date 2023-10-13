// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#name AppServiceEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#resource_group_name AppServiceEnvironment#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#subnet_id AppServiceEnvironment#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#allowed_user_ip_cidrs AppServiceEnvironment#allowed_user_ip_cidrs}.
	AllowedUserIpCidrs *[]*string `field:"optional" json:"allowedUserIpCidrs" yaml:"allowedUserIpCidrs"`
	// cluster_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#cluster_setting AppServiceEnvironment#cluster_setting}
	ClusterSetting interface{} `field:"optional" json:"clusterSetting" yaml:"clusterSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#front_end_scale_factor AppServiceEnvironment#front_end_scale_factor}.
	FrontEndScaleFactor *float64 `field:"optional" json:"frontEndScaleFactor" yaml:"frontEndScaleFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#id AppServiceEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#internal_load_balancing_mode AppServiceEnvironment#internal_load_balancing_mode}.
	InternalLoadBalancingMode *string `field:"optional" json:"internalLoadBalancingMode" yaml:"internalLoadBalancingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#pricing_tier AppServiceEnvironment#pricing_tier}.
	PricingTier *string `field:"optional" json:"pricingTier" yaml:"pricingTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#tags AppServiceEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/app_service_environment#timeouts AppServiceEnvironment#timeouts}
	Timeouts *AppServiceEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

