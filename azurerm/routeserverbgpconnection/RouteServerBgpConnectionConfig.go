// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routeserverbgpconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RouteServerBgpConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#name RouteServerBgpConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#peer_asn RouteServerBgpConnection#peer_asn}.
	PeerAsn *float64 `field:"required" json:"peerAsn" yaml:"peerAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#peer_ip RouteServerBgpConnection#peer_ip}.
	PeerIp *string `field:"required" json:"peerIp" yaml:"peerIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#route_server_id RouteServerBgpConnection#route_server_id}.
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#id RouteServerBgpConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/route_server_bgp_connection#timeouts RouteServerBgpConnection#timeouts}
	Timeouts *RouteServerBgpConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

