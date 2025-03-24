// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateVolume struct {
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_job#name ContainerAppJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Mount options used while mounting the AzureFile. Must be a comma-separated string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_job#mount_options ContainerAppJob#mount_options}
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The name of the `AzureFile` storage. Required when `storage_type` is `AzureFile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_job#storage_name ContainerAppJob#storage_name}
	StorageName *string `field:"optional" json:"storageName" yaml:"storageName"`
	// The type of storage volume. Possible values include `AzureFile` and `EmptyDir`. Defaults to `EmptyDir`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/container_app_job#storage_type ContainerAppJob#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

