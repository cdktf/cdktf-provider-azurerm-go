// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork


type VirtualNetworkSubnet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#address_prefixes VirtualNetwork#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#default_outbound_access_enabled VirtualNetwork#default_outbound_access_enabled}.
	DefaultOutboundAccessEnabled interface{} `field:"optional" json:"defaultOutboundAccessEnabled" yaml:"defaultOutboundAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#delegation VirtualNetwork#delegation}.
	Delegation interface{} `field:"optional" json:"delegation" yaml:"delegation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#id VirtualNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#name VirtualNetwork#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#private_endpoint_network_policies VirtualNetwork#private_endpoint_network_policies}.
	PrivateEndpointNetworkPolicies *string `field:"optional" json:"privateEndpointNetworkPolicies" yaml:"privateEndpointNetworkPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#private_link_service_network_policies_enabled VirtualNetwork#private_link_service_network_policies_enabled}.
	PrivateLinkServiceNetworkPoliciesEnabled interface{} `field:"optional" json:"privateLinkServiceNetworkPoliciesEnabled" yaml:"privateLinkServiceNetworkPoliciesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#route_table_id VirtualNetwork#route_table_id}.
	RouteTableId *string `field:"optional" json:"routeTableId" yaml:"routeTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#security_group VirtualNetwork#security_group}.
	SecurityGroup *string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#service_endpoint_policy_ids VirtualNetwork#service_endpoint_policy_ids}.
	ServiceEndpointPolicyIds *[]*string `field:"optional" json:"serviceEndpointPolicyIds" yaml:"serviceEndpointPolicyIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/virtual_network#service_endpoints VirtualNetwork#service_endpoints}.
	ServiceEndpoints *[]*string `field:"optional" json:"serviceEndpoints" yaml:"serviceEndpoints"`
}

