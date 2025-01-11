// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsubnetworkacl


type WebPubsubNetworkAclPrivateEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_pubsub_network_acl#id WebPubsubNetworkAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_pubsub_network_acl#allowed_request_types WebPubsubNetworkAcl#allowed_request_types}.
	AllowedRequestTypes *[]*string `field:"optional" json:"allowedRequestTypes" yaml:"allowedRequestTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/web_pubsub_network_acl#denied_request_types WebPubsubNetworkAcl#denied_request_types}.
	DeniedRequestTypes *[]*string `field:"optional" json:"deniedRequestTypes" yaml:"deniedRequestTypes"`
}

