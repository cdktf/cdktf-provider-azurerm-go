// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingendpoint


type MediaStreamingEndpointAccessControlAkamaiSignatureHeaderAuthenticationKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/media_streaming_endpoint#base64_key MediaStreamingEndpoint#base64_key}.
	Base64Key *string `field:"optional" json:"base64Key" yaml:"base64Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/media_streaming_endpoint#expiration MediaStreamingEndpoint#expiration}.
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/media_streaming_endpoint#identifier MediaStreamingEndpoint#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

