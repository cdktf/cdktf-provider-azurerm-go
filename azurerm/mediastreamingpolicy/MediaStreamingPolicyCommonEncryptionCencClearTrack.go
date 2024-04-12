// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCencClearTrack struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/media_streaming_policy#condition MediaStreamingPolicy#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
}

