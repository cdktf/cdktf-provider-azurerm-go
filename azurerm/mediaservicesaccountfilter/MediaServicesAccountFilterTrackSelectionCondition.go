// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccountfilter


type MediaServicesAccountFilterTrackSelectionCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/media_services_account_filter#operation MediaServicesAccountFilter#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/media_services_account_filter#property MediaServicesAccountFilter#property}.
	Property *string `field:"required" json:"property" yaml:"property"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/media_services_account_filter#value MediaServicesAccountFilter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

