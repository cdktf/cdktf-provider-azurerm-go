// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fabriccapacity


type FabricCapacitySku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/fabric_capacity#name FabricCapacity#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/fabric_capacity#tier FabricCapacity#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
}

