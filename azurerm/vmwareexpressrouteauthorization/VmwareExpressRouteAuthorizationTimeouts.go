// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareexpressrouteauthorization


type VmwareExpressRouteAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vmware_express_route_authorization#create VmwareExpressRouteAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vmware_express_route_authorization#delete VmwareExpressRouteAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/vmware_express_route_authorization#read VmwareExpressRouteAuthorization#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

