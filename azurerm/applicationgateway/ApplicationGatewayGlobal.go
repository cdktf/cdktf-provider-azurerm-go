// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayGlobal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/application_gateway#request_buffering_enabled ApplicationGateway#request_buffering_enabled}.
	RequestBufferingEnabled interface{} `field:"required" json:"requestBufferingEnabled" yaml:"requestBufferingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/application_gateway#response_buffering_enabled ApplicationGateway#response_buffering_enabled}.
	ResponseBufferingEnabled interface{} `field:"required" json:"responseBufferingEnabled" yaml:"responseBufferingEnabled"`
}

