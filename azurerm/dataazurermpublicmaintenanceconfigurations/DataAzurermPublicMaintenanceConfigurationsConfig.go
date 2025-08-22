// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermpublicmaintenanceconfigurations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermPublicMaintenanceConfigurationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/data-sources/public_maintenance_configurations#id DataAzurermPublicMaintenanceConfigurations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/data-sources/public_maintenance_configurations#location DataAzurermPublicMaintenanceConfigurations#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/data-sources/public_maintenance_configurations#recur_every DataAzurermPublicMaintenanceConfigurations#recur_every}.
	RecurEvery *string `field:"optional" json:"recurEvery" yaml:"recurEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/data-sources/public_maintenance_configurations#scope DataAzurermPublicMaintenanceConfigurations#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/data-sources/public_maintenance_configurations#timeouts DataAzurermPublicMaintenanceConfigurations#timeouts}
	Timeouts *DataAzurermPublicMaintenanceConfigurationsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

