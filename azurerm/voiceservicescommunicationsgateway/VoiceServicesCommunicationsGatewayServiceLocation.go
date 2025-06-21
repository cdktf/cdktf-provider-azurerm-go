// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package voiceservicescommunicationsgateway


type VoiceServicesCommunicationsGatewayServiceLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/voice_services_communications_gateway#location VoiceServicesCommunicationsGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/voice_services_communications_gateway#operator_addresses VoiceServicesCommunicationsGateway#operator_addresses}.
	OperatorAddresses *[]*string `field:"required" json:"operatorAddresses" yaml:"operatorAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/voice_services_communications_gateway#allowed_media_source_address_prefixes VoiceServicesCommunicationsGateway#allowed_media_source_address_prefixes}.
	AllowedMediaSourceAddressPrefixes *[]*string `field:"optional" json:"allowedMediaSourceAddressPrefixes" yaml:"allowedMediaSourceAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/voice_services_communications_gateway#allowed_signaling_source_address_prefixes VoiceServicesCommunicationsGateway#allowed_signaling_source_address_prefixes}.
	AllowedSignalingSourceAddressPrefixes *[]*string `field:"optional" json:"allowedSignalingSourceAddressPrefixes" yaml:"allowedSignalingSourceAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/voice_services_communications_gateway#esrp_addresses VoiceServicesCommunicationsGateway#esrp_addresses}.
	EsrpAddresses *[]*string `field:"optional" json:"esrpAddresses" yaml:"esrpAddresses"`
}

