// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package restorepointcollection


type RestorePointCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/restore_point_collection#create RestorePointCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/restore_point_collection#delete RestorePointCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/restore_point_collection#read RestorePointCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/restore_point_collection#update RestorePointCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

