// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingendpoint


type MediaStreamingEndpointAccessControl struct {
	// akamai_signature_header_authentication_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_endpoint#akamai_signature_header_authentication_key MediaStreamingEndpoint#akamai_signature_header_authentication_key}
	AkamaiSignatureHeaderAuthenticationKey interface{} `field:"optional" json:"akamaiSignatureHeaderAuthenticationKey" yaml:"akamaiSignatureHeaderAuthenticationKey"`
	// ip_allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_streaming_endpoint#ip_allow MediaStreamingEndpoint#ip_allow}
	IpAllow interface{} `field:"optional" json:"ipAllow" yaml:"ipAllow"`
}

