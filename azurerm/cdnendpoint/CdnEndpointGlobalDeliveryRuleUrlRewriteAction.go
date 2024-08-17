// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint


type CdnEndpointGlobalDeliveryRuleUrlRewriteAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/cdn_endpoint#destination CdnEndpoint#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/cdn_endpoint#source_pattern CdnEndpoint#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/cdn_endpoint#preserve_unmatched_path CdnEndpoint#preserve_unmatched_path}.
	PreserveUnmatchedPath interface{} `field:"optional" json:"preserveUnmatchedPath" yaml:"preserveUnmatchedPath"`
}

