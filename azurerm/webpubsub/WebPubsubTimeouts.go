// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webpubsub


type WebPubsubTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/web_pubsub#create WebPubsub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/web_pubsub#delete WebPubsub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/web_pubsub#read WebPubsub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/web_pubsub#update WebPubsub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

