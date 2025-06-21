// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkGatewayConfig struct {
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
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#ip_configuration VirtualNetworkGateway#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#location VirtualNetworkGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#resource_group_name VirtualNetworkGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#sku VirtualNetworkGateway#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#type VirtualNetworkGateway#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#active_active VirtualNetworkGateway#active_active}.
	ActiveActive interface{} `field:"optional" json:"activeActive" yaml:"activeActive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#bgp_route_translation_for_nat_enabled VirtualNetworkGateway#bgp_route_translation_for_nat_enabled}.
	BgpRouteTranslationForNatEnabled interface{} `field:"optional" json:"bgpRouteTranslationForNatEnabled" yaml:"bgpRouteTranslationForNatEnabled"`
	// bgp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#bgp_settings VirtualNetworkGateway#bgp_settings}
	BgpSettings *VirtualNetworkGatewayBgpSettings `field:"optional" json:"bgpSettings" yaml:"bgpSettings"`
	// custom_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#custom_route VirtualNetworkGateway#custom_route}
	CustomRoute *VirtualNetworkGatewayCustomRoute `field:"optional" json:"customRoute" yaml:"customRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#default_local_network_gateway_id VirtualNetworkGateway#default_local_network_gateway_id}.
	DefaultLocalNetworkGatewayId *string `field:"optional" json:"defaultLocalNetworkGatewayId" yaml:"defaultLocalNetworkGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#dns_forwarding_enabled VirtualNetworkGateway#dns_forwarding_enabled}.
	DnsForwardingEnabled interface{} `field:"optional" json:"dnsForwardingEnabled" yaml:"dnsForwardingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#edge_zone VirtualNetworkGateway#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#enable_bgp VirtualNetworkGateway#enable_bgp}.
	EnableBgp interface{} `field:"optional" json:"enableBgp" yaml:"enableBgp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#generation VirtualNetworkGateway#generation}.
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#id VirtualNetworkGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#ip_sec_replay_protection_enabled VirtualNetworkGateway#ip_sec_replay_protection_enabled}.
	IpSecReplayProtectionEnabled interface{} `field:"optional" json:"ipSecReplayProtectionEnabled" yaml:"ipSecReplayProtectionEnabled"`
	// policy_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#policy_group VirtualNetworkGateway#policy_group}
	PolicyGroup interface{} `field:"optional" json:"policyGroup" yaml:"policyGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#private_ip_address_enabled VirtualNetworkGateway#private_ip_address_enabled}.
	PrivateIpAddressEnabled interface{} `field:"optional" json:"privateIpAddressEnabled" yaml:"privateIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#remote_vnet_traffic_enabled VirtualNetworkGateway#remote_vnet_traffic_enabled}.
	RemoteVnetTrafficEnabled interface{} `field:"optional" json:"remoteVnetTrafficEnabled" yaml:"remoteVnetTrafficEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#tags VirtualNetworkGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#timeouts VirtualNetworkGateway#timeouts}
	Timeouts *VirtualNetworkGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#virtual_wan_traffic_enabled VirtualNetworkGateway#virtual_wan_traffic_enabled}.
	VirtualWanTrafficEnabled interface{} `field:"optional" json:"virtualWanTrafficEnabled" yaml:"virtualWanTrafficEnabled"`
	// vpn_client_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#vpn_client_configuration VirtualNetworkGateway#vpn_client_configuration}
	VpnClientConfiguration *VirtualNetworkGatewayVpnClientConfiguration `field:"optional" json:"vpnClientConfiguration" yaml:"vpnClientConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/virtual_network_gateway#vpn_type VirtualNetworkGateway#vpn_type}.
	VpnType *string `field:"optional" json:"vpnType" yaml:"vpnType"`
}

