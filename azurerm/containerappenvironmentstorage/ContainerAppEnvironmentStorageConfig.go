// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppEnvironmentStorageConfig struct {
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
	// The Storage Account Access Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#access_key ContainerAppEnvironmentStorage#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// The access mode to connect this storage to the Container App. Possible values include `ReadOnly` and `ReadWrite`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#access_mode ContainerAppEnvironmentStorage#access_mode}
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// The Azure Storage Account in which the Share to be used is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#account_name ContainerAppEnvironmentStorage#account_name}
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// The ID of the Container App Environment to which this storage belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#container_app_environment_id ContainerAppEnvironmentStorage#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name for this Storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#name ContainerAppEnvironmentStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Azure Storage Share to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#share_name ContainerAppEnvironmentStorage#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#id ContainerAppEnvironmentStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/container_app_environment_storage#timeouts ContainerAppEnvironmentStorage#timeouts}
	Timeouts *ContainerAppEnvironmentStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

