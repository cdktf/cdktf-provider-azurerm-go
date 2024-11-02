// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package relayhybridconnection


type RelayHybridConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/relay_hybrid_connection#create RelayHybridConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/relay_hybrid_connection#delete RelayHybridConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/relay_hybrid_connection#read RelayHybridConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/relay_hybrid_connection#update RelayHybridConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

