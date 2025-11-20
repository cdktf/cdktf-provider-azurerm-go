// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewayconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkGatewayConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#location VirtualNetworkGatewayConnection#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#name VirtualNetworkGatewayConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#resource_group_name VirtualNetworkGatewayConnection#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#type VirtualNetworkGatewayConnection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#virtual_network_gateway_id VirtualNetworkGatewayConnection#virtual_network_gateway_id}.
	VirtualNetworkGatewayId *string `field:"required" json:"virtualNetworkGatewayId" yaml:"virtualNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#authorization_key VirtualNetworkGatewayConnection#authorization_key}.
	AuthorizationKey *string `field:"optional" json:"authorizationKey" yaml:"authorizationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#connection_mode VirtualNetworkGatewayConnection#connection_mode}.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#connection_protocol VirtualNetworkGatewayConnection#connection_protocol}.
	ConnectionProtocol *string `field:"optional" json:"connectionProtocol" yaml:"connectionProtocol"`
	// custom_bgp_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#custom_bgp_addresses VirtualNetworkGatewayConnection#custom_bgp_addresses}
	CustomBgpAddresses *VirtualNetworkGatewayConnectionCustomBgpAddresses `field:"optional" json:"customBgpAddresses" yaml:"customBgpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#dpd_timeout_seconds VirtualNetworkGatewayConnection#dpd_timeout_seconds}.
	DpdTimeoutSeconds *float64 `field:"optional" json:"dpdTimeoutSeconds" yaml:"dpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#egress_nat_rule_ids VirtualNetworkGatewayConnection#egress_nat_rule_ids}.
	EgressNatRuleIds *[]*string `field:"optional" json:"egressNatRuleIds" yaml:"egressNatRuleIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#enable_bgp VirtualNetworkGatewayConnection#enable_bgp}.
	EnableBgp interface{} `field:"optional" json:"enableBgp" yaml:"enableBgp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#express_route_circuit_id VirtualNetworkGatewayConnection#express_route_circuit_id}.
	ExpressRouteCircuitId *string `field:"optional" json:"expressRouteCircuitId" yaml:"expressRouteCircuitId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#express_route_gateway_bypass VirtualNetworkGatewayConnection#express_route_gateway_bypass}.
	ExpressRouteGatewayBypass interface{} `field:"optional" json:"expressRouteGatewayBypass" yaml:"expressRouteGatewayBypass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#id VirtualNetworkGatewayConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#ingress_nat_rule_ids VirtualNetworkGatewayConnection#ingress_nat_rule_ids}.
	IngressNatRuleIds *[]*string `field:"optional" json:"ingressNatRuleIds" yaml:"ingressNatRuleIds"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#ipsec_policy VirtualNetworkGatewayConnection#ipsec_policy}
	IpsecPolicy *VirtualNetworkGatewayConnectionIpsecPolicy `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#local_azure_ip_address_enabled VirtualNetworkGatewayConnection#local_azure_ip_address_enabled}.
	LocalAzureIpAddressEnabled interface{} `field:"optional" json:"localAzureIpAddressEnabled" yaml:"localAzureIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#local_network_gateway_id VirtualNetworkGatewayConnection#local_network_gateway_id}.
	LocalNetworkGatewayId *string `field:"optional" json:"localNetworkGatewayId" yaml:"localNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#peer_virtual_network_gateway_id VirtualNetworkGatewayConnection#peer_virtual_network_gateway_id}.
	PeerVirtualNetworkGatewayId *string `field:"optional" json:"peerVirtualNetworkGatewayId" yaml:"peerVirtualNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#private_link_fast_path_enabled VirtualNetworkGatewayConnection#private_link_fast_path_enabled}.
	PrivateLinkFastPathEnabled interface{} `field:"optional" json:"privateLinkFastPathEnabled" yaml:"privateLinkFastPathEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#routing_weight VirtualNetworkGatewayConnection#routing_weight}.
	RoutingWeight *float64 `field:"optional" json:"routingWeight" yaml:"routingWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#shared_key VirtualNetworkGatewayConnection#shared_key}.
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#tags VirtualNetworkGatewayConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#timeouts VirtualNetworkGatewayConnection#timeouts}
	Timeouts *VirtualNetworkGatewayConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_selector_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#traffic_selector_policy VirtualNetworkGatewayConnection#traffic_selector_policy}
	TrafficSelectorPolicy interface{} `field:"optional" json:"trafficSelectorPolicy" yaml:"trafficSelectorPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_network_gateway_connection#use_policy_based_traffic_selectors VirtualNetworkGatewayConnection#use_policy_based_traffic_selectors}.
	UsePolicyBasedTrafficSelectors interface{} `field:"optional" json:"usePolicyBasedTrafficSelectors" yaml:"usePolicyBasedTrafficSelectors"`
}

