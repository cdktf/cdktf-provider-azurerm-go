// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routeserverbgpconnection


type RouteServerBgpConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/route_server_bgp_connection#create RouteServerBgpConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/route_server_bgp_connection#delete RouteServerBgpConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/route_server_bgp_connection#read RouteServerBgpConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

