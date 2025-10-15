// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fabriccapacity


type FabricCapacityTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/fabric_capacity#create FabricCapacity#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/fabric_capacity#delete FabricCapacity#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/fabric_capacity#read FabricCapacity#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/fabric_capacity#update FabricCapacity#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

