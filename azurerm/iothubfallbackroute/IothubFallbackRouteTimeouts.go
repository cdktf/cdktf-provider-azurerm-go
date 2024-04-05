// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubfallbackroute


type IothubFallbackRouteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/iothub_fallback_route#create IothubFallbackRouteA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/iothub_fallback_route#delete IothubFallbackRouteA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/iothub_fallback_route#read IothubFallbackRouteA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/iothub_fallback_route#update IothubFallbackRouteA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

