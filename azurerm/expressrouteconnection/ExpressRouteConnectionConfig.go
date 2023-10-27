// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package expressrouteconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRouteConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#express_route_circuit_peering_id ExpressRouteConnection#express_route_circuit_peering_id}.
	ExpressRouteCircuitPeeringId *string `field:"required" json:"expressRouteCircuitPeeringId" yaml:"expressRouteCircuitPeeringId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#express_route_gateway_id ExpressRouteConnection#express_route_gateway_id}.
	ExpressRouteGatewayId *string `field:"required" json:"expressRouteGatewayId" yaml:"expressRouteGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#name ExpressRouteConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#authorization_key ExpressRouteConnection#authorization_key}.
	AuthorizationKey *string `field:"optional" json:"authorizationKey" yaml:"authorizationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#enable_internet_security ExpressRouteConnection#enable_internet_security}.
	EnableInternetSecurity interface{} `field:"optional" json:"enableInternetSecurity" yaml:"enableInternetSecurity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#express_route_gateway_bypass_enabled ExpressRouteConnection#express_route_gateway_bypass_enabled}.
	ExpressRouteGatewayBypassEnabled interface{} `field:"optional" json:"expressRouteGatewayBypassEnabled" yaml:"expressRouteGatewayBypassEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#id ExpressRouteConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#routing ExpressRouteConnection#routing}
	Routing *ExpressRouteConnectionRouting `field:"optional" json:"routing" yaml:"routing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#routing_weight ExpressRouteConnection#routing_weight}.
	RoutingWeight *float64 `field:"optional" json:"routingWeight" yaml:"routingWeight"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/express_route_connection#timeouts ExpressRouteConnection#timeouts}
	Timeouts *ExpressRouteConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

