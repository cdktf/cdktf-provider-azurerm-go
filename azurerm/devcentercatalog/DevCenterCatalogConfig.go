// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcentercatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevCenterCatalogConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#dev_center_id DevCenterCatalog#dev_center_id}.
	DevCenterId *string `field:"required" json:"devCenterId" yaml:"devCenterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#name DevCenterCatalog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#resource_group_name DevCenterCatalog#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// catalog_adogit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#catalog_adogit DevCenterCatalog#catalog_adogit}
	CatalogAdogit *DevCenterCatalogCatalogAdogit `field:"optional" json:"catalogAdogit" yaml:"catalogAdogit"`
	// catalog_github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#catalog_github DevCenterCatalog#catalog_github}
	CatalogGithub *DevCenterCatalogCatalogGithub `field:"optional" json:"catalogGithub" yaml:"catalogGithub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#id DevCenterCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/dev_center_catalog#timeouts DevCenterCatalog#timeouts}
	Timeouts *DevCenterCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

