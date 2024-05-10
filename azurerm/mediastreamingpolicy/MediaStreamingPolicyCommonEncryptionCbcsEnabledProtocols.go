// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_streaming_policy#dash MediaStreamingPolicy#dash}.
	Dash interface{} `field:"optional" json:"dash" yaml:"dash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_streaming_policy#download MediaStreamingPolicy#download}.
	Download interface{} `field:"optional" json:"download" yaml:"download"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_streaming_policy#hls MediaStreamingPolicy#hls}.
	Hls interface{} `field:"optional" json:"hls" yaml:"hls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_streaming_policy#smooth_streaming MediaStreamingPolicy#smooth_streaming}.
	SmoothStreaming interface{} `field:"optional" json:"smoothStreaming" yaml:"smoothStreaming"`
}

