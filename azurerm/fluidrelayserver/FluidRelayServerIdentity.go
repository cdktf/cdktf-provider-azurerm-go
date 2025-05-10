// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fluidrelayserver


type FluidRelayServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/fluid_relay_server#type FluidRelayServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/fluid_relay_server#identity_ids FluidRelayServer#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

