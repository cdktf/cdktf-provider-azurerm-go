// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudapp


type SpringCloudAppCustomPersistentDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/spring_cloud_app#mount_path SpringCloudApp#mount_path}.
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/spring_cloud_app#share_name SpringCloudApp#share_name}.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/spring_cloud_app#storage_name SpringCloudApp#storage_name}.
	StorageName *string `field:"required" json:"storageName" yaml:"storageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/spring_cloud_app#mount_options SpringCloudApp#mount_options}.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/spring_cloud_app#read_only_enabled SpringCloudApp#read_only_enabled}.
	ReadOnlyEnabled interface{} `field:"optional" json:"readOnlyEnabled" yaml:"readOnlyEnabled"`
}

