// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package capacityreservation


type CapacityReservationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/capacity_reservation#create CapacityReservation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/capacity_reservation#delete CapacityReservation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/capacity_reservation#read CapacityReservation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/capacity_reservation#update CapacityReservation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

