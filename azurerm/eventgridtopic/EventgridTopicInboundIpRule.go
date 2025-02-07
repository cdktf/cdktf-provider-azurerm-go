// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridtopic


type EventgridTopicInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/eventgrid_topic#action EventgridTopic#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/eventgrid_topic#ip_mask EventgridTopic#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

