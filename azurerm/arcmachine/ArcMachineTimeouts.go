// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arcmachine


type ArcMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/arc_machine#create ArcMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/arc_machine#delete ArcMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/arc_machine#read ArcMachine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/arc_machine#update ArcMachine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

