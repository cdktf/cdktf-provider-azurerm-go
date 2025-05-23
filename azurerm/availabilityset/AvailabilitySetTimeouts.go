// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package availabilityset


type AvailabilitySetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/availability_set#create AvailabilitySet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/availability_set#delete AvailabilitySet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/availability_set#read AvailabilitySet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/availability_set#update AvailabilitySet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

