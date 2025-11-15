// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsubsocketio


type WebPubsubSocketioSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/web_pubsub_socketio#name WebPubsubSocketio#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/web_pubsub_socketio#capacity WebPubsubSocketio#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

