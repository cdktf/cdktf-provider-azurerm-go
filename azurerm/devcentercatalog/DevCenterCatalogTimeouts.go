// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcentercatalog


type DevCenterCatalogTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/dev_center_catalog#create DevCenterCatalog#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/dev_center_catalog#delete DevCenterCatalog#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/dev_center_catalog#read DevCenterCatalog#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/dev_center_catalog#update DevCenterCatalog#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

