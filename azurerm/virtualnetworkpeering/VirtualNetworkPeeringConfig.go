// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkPeeringConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#name VirtualNetworkPeering#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#remote_virtual_network_id VirtualNetworkPeering#remote_virtual_network_id}.
	RemoteVirtualNetworkId *string `field:"required" json:"remoteVirtualNetworkId" yaml:"remoteVirtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#resource_group_name VirtualNetworkPeering#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#virtual_network_name VirtualNetworkPeering#virtual_network_name}.
	VirtualNetworkName *string `field:"required" json:"virtualNetworkName" yaml:"virtualNetworkName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#allow_forwarded_traffic VirtualNetworkPeering#allow_forwarded_traffic}.
	AllowForwardedTraffic interface{} `field:"optional" json:"allowForwardedTraffic" yaml:"allowForwardedTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#allow_gateway_transit VirtualNetworkPeering#allow_gateway_transit}.
	AllowGatewayTransit interface{} `field:"optional" json:"allowGatewayTransit" yaml:"allowGatewayTransit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#allow_virtual_network_access VirtualNetworkPeering#allow_virtual_network_access}.
	AllowVirtualNetworkAccess interface{} `field:"optional" json:"allowVirtualNetworkAccess" yaml:"allowVirtualNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#id VirtualNetworkPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#local_subnet_names VirtualNetworkPeering#local_subnet_names}.
	LocalSubnetNames *[]*string `field:"optional" json:"localSubnetNames" yaml:"localSubnetNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#only_ipv6_peering_enabled VirtualNetworkPeering#only_ipv6_peering_enabled}.
	OnlyIpv6PeeringEnabled interface{} `field:"optional" json:"onlyIpv6PeeringEnabled" yaml:"onlyIpv6PeeringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#peer_complete_virtual_networks_enabled VirtualNetworkPeering#peer_complete_virtual_networks_enabled}.
	PeerCompleteVirtualNetworksEnabled interface{} `field:"optional" json:"peerCompleteVirtualNetworksEnabled" yaml:"peerCompleteVirtualNetworksEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#remote_subnet_names VirtualNetworkPeering#remote_subnet_names}.
	RemoteSubnetNames *[]*string `field:"optional" json:"remoteSubnetNames" yaml:"remoteSubnetNames"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#timeouts VirtualNetworkPeering#timeouts}
	Timeouts *VirtualNetworkPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#triggers VirtualNetworkPeering#triggers}.
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_network_peering#use_remote_gateways VirtualNetworkPeering#use_remote_gateways}.
	UseRemoteGateways interface{} `field:"optional" json:"useRemoteGateways" yaml:"useRemoteGateways"`
}

