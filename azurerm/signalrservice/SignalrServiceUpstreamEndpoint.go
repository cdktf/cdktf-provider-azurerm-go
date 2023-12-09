// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservice


type SignalrServiceUpstreamEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/signalr_service#category_pattern SignalrService#category_pattern}.
	CategoryPattern *[]*string `field:"required" json:"categoryPattern" yaml:"categoryPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/signalr_service#event_pattern SignalrService#event_pattern}.
	EventPattern *[]*string `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/signalr_service#hub_pattern SignalrService#hub_pattern}.
	HubPattern *[]*string `field:"required" json:"hubPattern" yaml:"hubPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/signalr_service#url_template SignalrService#url_template}.
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/signalr_service#user_assigned_identity_id SignalrService#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

