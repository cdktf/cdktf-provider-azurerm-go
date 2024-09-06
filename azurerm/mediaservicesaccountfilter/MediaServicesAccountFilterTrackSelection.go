// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccountfilter


type MediaServicesAccountFilterTrackSelection struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/media_services_account_filter#condition MediaServicesAccountFilter#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
}

