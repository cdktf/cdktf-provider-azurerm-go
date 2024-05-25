// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapphybridconnection


type WebAppHybridConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/web_app_hybrid_connection#create WebAppHybridConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/web_app_hybrid_connection#delete WebAppHybridConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/web_app_hybrid_connection#read WebAppHybridConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/web_app_hybrid_connection#update WebAppHybridConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

