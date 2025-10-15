// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arcmachine


type ArcMachineIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/arc_machine#type ArcMachine#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

