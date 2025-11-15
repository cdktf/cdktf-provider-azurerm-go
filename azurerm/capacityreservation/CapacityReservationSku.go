// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package capacityreservation


type CapacityReservationSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/capacity_reservation#capacity CapacityReservation#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/capacity_reservation#name CapacityReservation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

