// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridnamespace


type EventgridNamespaceInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/eventgrid_namespace#ip_mask EventgridNamespace#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/eventgrid_namespace#action EventgridNamespace#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

