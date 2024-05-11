// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddisk


type ManagedDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/managed_disk#create ManagedDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/managed_disk#delete ManagedDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/managed_disk#read ManagedDisk#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/managed_disk#update ManagedDisk#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

