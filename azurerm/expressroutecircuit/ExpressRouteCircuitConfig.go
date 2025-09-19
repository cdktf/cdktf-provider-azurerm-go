// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package expressroutecircuit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRouteCircuitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#location ExpressRouteCircuit#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#name ExpressRouteCircuit#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#resource_group_name ExpressRouteCircuit#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#sku ExpressRouteCircuit#sku}
	Sku *ExpressRouteCircuitSku `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#allow_classic_operations ExpressRouteCircuit#allow_classic_operations}.
	AllowClassicOperations interface{} `field:"optional" json:"allowClassicOperations" yaml:"allowClassicOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#authorization_key ExpressRouteCircuit#authorization_key}.
	AuthorizationKey *string `field:"optional" json:"authorizationKey" yaml:"authorizationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#bandwidth_in_gbps ExpressRouteCircuit#bandwidth_in_gbps}.
	BandwidthInGbps *float64 `field:"optional" json:"bandwidthInGbps" yaml:"bandwidthInGbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#bandwidth_in_mbps ExpressRouteCircuit#bandwidth_in_mbps}.
	BandwidthInMbps *float64 `field:"optional" json:"bandwidthInMbps" yaml:"bandwidthInMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#express_route_port_id ExpressRouteCircuit#express_route_port_id}.
	ExpressRoutePortId *string `field:"optional" json:"expressRoutePortId" yaml:"expressRoutePortId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#id ExpressRouteCircuit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#peering_location ExpressRouteCircuit#peering_location}.
	PeeringLocation *string `field:"optional" json:"peeringLocation" yaml:"peeringLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#rate_limiting_enabled ExpressRouteCircuit#rate_limiting_enabled}.
	RateLimitingEnabled interface{} `field:"optional" json:"rateLimitingEnabled" yaml:"rateLimitingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#service_provider_name ExpressRouteCircuit#service_provider_name}.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#tags ExpressRouteCircuit#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/express_route_circuit#timeouts ExpressRouteCircuit#timeouts}
	Timeouts *ExpressRouteCircuitTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

