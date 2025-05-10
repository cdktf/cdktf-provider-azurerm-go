// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsubsocketio


type WebPubsubSocketioIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/web_pubsub_socketio#type WebPubsubSocketio#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/web_pubsub_socketio#identity_ids WebPubsubSocketio#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

