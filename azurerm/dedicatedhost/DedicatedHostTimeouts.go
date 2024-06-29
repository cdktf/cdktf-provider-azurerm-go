// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dedicatedhost


type DedicatedHostTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/dedicated_host#create DedicatedHost#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/dedicated_host#delete DedicatedHost#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/dedicated_host#read DedicatedHost#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/dedicated_host#update DedicatedHost#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

