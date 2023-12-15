// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccount


type MediaServicesAccountKeyDeliveryAccessControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/media_services_account#default_action MediaServicesAccount#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/media_services_account#ip_allow_list MediaServicesAccount#ip_allow_list}.
	IpAllowList *[]*string `field:"optional" json:"ipAllowList" yaml:"ipAllowList"`
}

