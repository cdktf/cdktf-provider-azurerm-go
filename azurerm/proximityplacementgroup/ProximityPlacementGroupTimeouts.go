// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proximityplacementgroup


type ProximityPlacementGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/proximity_placement_group#create ProximityPlacementGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/proximity_placement_group#delete ProximityPlacementGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/proximity_placement_group#read ProximityPlacementGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/proximity_placement_group#update ProximityPlacementGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

