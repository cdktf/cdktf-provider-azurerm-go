// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrontdoorConfig struct {
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
	// backend_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#backend_pool Frontdoor#backend_pool}
	BackendPool interface{} `field:"required" json:"backendPool" yaml:"backendPool"`
	// backend_pool_health_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#backend_pool_health_probe Frontdoor#backend_pool_health_probe}
	BackendPoolHealthProbe interface{} `field:"required" json:"backendPoolHealthProbe" yaml:"backendPoolHealthProbe"`
	// backend_pool_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#backend_pool_load_balancing Frontdoor#backend_pool_load_balancing}
	BackendPoolLoadBalancing interface{} `field:"required" json:"backendPoolLoadBalancing" yaml:"backendPoolLoadBalancing"`
	// frontend_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#frontend_endpoint Frontdoor#frontend_endpoint}
	FrontendEndpoint interface{} `field:"required" json:"frontendEndpoint" yaml:"frontendEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#resource_group_name Frontdoor#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#routing_rule Frontdoor#routing_rule}
	RoutingRule interface{} `field:"required" json:"routingRule" yaml:"routingRule"`
	// backend_pool_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#backend_pool_settings Frontdoor#backend_pool_settings}
	BackendPoolSettings interface{} `field:"optional" json:"backendPoolSettings" yaml:"backendPoolSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#friendly_name Frontdoor#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#id Frontdoor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#load_balancer_enabled Frontdoor#load_balancer_enabled}.
	LoadBalancerEnabled interface{} `field:"optional" json:"loadBalancerEnabled" yaml:"loadBalancerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#tags Frontdoor#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.2/docs/resources/frontdoor#timeouts Frontdoor#timeouts}
	Timeouts *FrontdoorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

