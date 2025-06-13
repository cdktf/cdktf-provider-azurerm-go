// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aadb2cdirectory


type Aadb2CDirectoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/aadb2c_directory#create Aadb2CDirectory#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/aadb2c_directory#delete Aadb2CDirectory#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/aadb2c_directory#read Aadb2CDirectory#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/aadb2c_directory#update Aadb2CDirectory#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

