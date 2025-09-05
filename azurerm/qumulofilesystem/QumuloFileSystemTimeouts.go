// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qumulofilesystem


type QumuloFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/qumulo_file_system#create QumuloFileSystem#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/qumulo_file_system#delete QumuloFileSystem#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/qumulo_file_system#read QumuloFileSystem#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/qumulo_file_system#update QumuloFileSystem#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

