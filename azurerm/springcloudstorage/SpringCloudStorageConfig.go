// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudStorageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#name SpringCloudStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#spring_cloud_service_id SpringCloudStorage#spring_cloud_service_id}.
	SpringCloudServiceId *string `field:"required" json:"springCloudServiceId" yaml:"springCloudServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#storage_account_key SpringCloudStorage#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#storage_account_name SpringCloudStorage#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#id SpringCloudStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/spring_cloud_storage#timeouts SpringCloudStorage#timeouts}
	Timeouts *SpringCloudStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

