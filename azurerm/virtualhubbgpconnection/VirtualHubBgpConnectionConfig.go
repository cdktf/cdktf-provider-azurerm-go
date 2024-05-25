// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualhubbgpconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualHubBgpConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#name VirtualHubBgpConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#peer_asn VirtualHubBgpConnection#peer_asn}.
	PeerAsn *float64 `field:"required" json:"peerAsn" yaml:"peerAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#peer_ip VirtualHubBgpConnection#peer_ip}.
	PeerIp *string `field:"required" json:"peerIp" yaml:"peerIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#virtual_hub_id VirtualHubBgpConnection#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#id VirtualHubBgpConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#timeouts VirtualHubBgpConnection#timeouts}
	Timeouts *VirtualHubBgpConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/virtual_hub_bgp_connection#virtual_network_connection_id VirtualHubBgpConnection#virtual_network_connection_id}.
	VirtualNetworkConnectionId *string `field:"optional" json:"virtualNetworkConnectionId" yaml:"virtualNetworkConnectionId"`
}

