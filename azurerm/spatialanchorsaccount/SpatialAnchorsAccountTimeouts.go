// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spatialanchorsaccount


type SpatialAnchorsAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/spatial_anchors_account#create SpatialAnchorsAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/spatial_anchors_account#delete SpatialAnchorsAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/spatial_anchors_account#read SpatialAnchorsAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/spatial_anchors_account#update SpatialAnchorsAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

