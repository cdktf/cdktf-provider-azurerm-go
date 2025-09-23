// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateVolume struct {
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Mount options used while mounting the AzureFile. Must be a comma-separated string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/container_app#mount_options ContainerApp#mount_options}
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The name of the `AzureFile` storage. Required when `storage_type` is `AzureFile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/container_app#storage_name ContainerApp#storage_name}
	StorageName *string `field:"optional" json:"storageName" yaml:"storageName"`
	// The type of storage volume. Possible values include `AzureFile` and `EmptyDir`. Defaults to `EmptyDir`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/container_app#storage_type ContainerApp#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

