// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package expressrouteport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRoutePortConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#bandwidth_in_gbps ExpressRoutePort#bandwidth_in_gbps}.
	BandwidthInGbps *float64 `field:"required" json:"bandwidthInGbps" yaml:"bandwidthInGbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#encapsulation ExpressRoutePort#encapsulation}.
	Encapsulation *string `field:"required" json:"encapsulation" yaml:"encapsulation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#location ExpressRoutePort#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#name ExpressRoutePort#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#peering_location ExpressRoutePort#peering_location}.
	PeeringLocation *string `field:"required" json:"peeringLocation" yaml:"peeringLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#resource_group_name ExpressRoutePort#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#billing_type ExpressRoutePort#billing_type}.
	BillingType *string `field:"optional" json:"billingType" yaml:"billingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#id ExpressRoutePort#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#identity ExpressRoutePort#identity}
	Identity *ExpressRoutePortIdentity `field:"optional" json:"identity" yaml:"identity"`
	// link1 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#link1 ExpressRoutePort#link1}
	Link1 *ExpressRoutePortLink1 `field:"optional" json:"link1" yaml:"link1"`
	// link2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#link2 ExpressRoutePort#link2}
	Link2 *ExpressRoutePortLink2 `field:"optional" json:"link2" yaml:"link2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#tags ExpressRoutePort#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/express_route_port#timeouts ExpressRoutePort#timeouts}
	Timeouts *ExpressRoutePortTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

