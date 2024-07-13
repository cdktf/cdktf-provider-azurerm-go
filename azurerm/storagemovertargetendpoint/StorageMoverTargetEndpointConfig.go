// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemovertargetendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageMoverTargetEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#name StorageMoverTargetEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#storage_account_id StorageMoverTargetEndpoint#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#storage_container_name StorageMoverTargetEndpoint#storage_container_name}.
	StorageContainerName *string `field:"required" json:"storageContainerName" yaml:"storageContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#storage_mover_id StorageMoverTargetEndpoint#storage_mover_id}.
	StorageMoverId *string `field:"required" json:"storageMoverId" yaml:"storageMoverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#description StorageMoverTargetEndpoint#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#id StorageMoverTargetEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/storage_mover_target_endpoint#timeouts StorageMoverTargetEndpoint#timeouts}
	Timeouts *StorageMoverTargetEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

