// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccountsas

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermStorageAccountSasConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#connection_string DataAzurermStorageAccountSas#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#expiry DataAzurermStorageAccountSas#expiry}.
	Expiry *string `field:"required" json:"expiry" yaml:"expiry"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#permissions DataAzurermStorageAccountSas#permissions}
	Permissions *DataAzurermStorageAccountSasPermissions `field:"required" json:"permissions" yaml:"permissions"`
	// resource_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#resource_types DataAzurermStorageAccountSas#resource_types}
	ResourceTypes *DataAzurermStorageAccountSasResourceTypes `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#services DataAzurermStorageAccountSas#services}
	Services *DataAzurermStorageAccountSasServices `field:"required" json:"services" yaml:"services"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#start DataAzurermStorageAccountSas#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#https_only DataAzurermStorageAccountSas#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#id DataAzurermStorageAccountSas#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#ip_addresses DataAzurermStorageAccountSas#ip_addresses}.
	IpAddresses *string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#signed_version DataAzurermStorageAccountSas#signed_version}.
	SignedVersion *string `field:"optional" json:"signedVersion" yaml:"signedVersion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/data-sources/storage_account_sas#timeouts DataAzurermStorageAccountSas#timeouts}
	Timeouts *DataAzurermStorageAccountSasTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

