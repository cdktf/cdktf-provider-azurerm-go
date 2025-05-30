// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoororigin


type CdnFrontdoorOriginPrivateLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_origin#location CdnFrontdoorOrigin#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_origin#private_link_target_id CdnFrontdoorOrigin#private_link_target_id}.
	PrivateLinkTargetId *string `field:"required" json:"privateLinkTargetId" yaml:"privateLinkTargetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_origin#request_message CdnFrontdoorOrigin#request_message}.
	RequestMessage *string `field:"optional" json:"requestMessage" yaml:"requestMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/cdn_frontdoor_origin#target_type CdnFrontdoorOrigin#target_type}.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

