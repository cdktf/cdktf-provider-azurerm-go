// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent


type MediaLiveEventCrossSiteAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event#client_access_policy MediaLiveEvent#client_access_policy}.
	ClientAccessPolicy *string `field:"optional" json:"clientAccessPolicy" yaml:"clientAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/media_live_event#cross_domain_policy MediaLiveEvent#cross_domain_policy}.
	CrossDomainPolicy *string `field:"optional" json:"crossDomainPolicy" yaml:"crossDomainPolicy"`
}

