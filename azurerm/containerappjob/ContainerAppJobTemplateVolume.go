// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateVolume struct {
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the `AzureFile` storage. Required when `storage_type` is `AzureFile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/container_app_job#storage_name ContainerAppJob#storage_name}
	StorageName *string `field:"optional" json:"storageName" yaml:"storageName"`
	// The type of storage volume. Possible values include `AzureFile` and `EmptyDir`. Defaults to `EmptyDir`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/container_app_job#storage_type ContainerAppJob#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

