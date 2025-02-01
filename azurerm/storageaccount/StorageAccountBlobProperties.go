// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountBlobProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#change_feed_enabled StorageAccount#change_feed_enabled}.
	ChangeFeedEnabled interface{} `field:"optional" json:"changeFeedEnabled" yaml:"changeFeedEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#change_feed_retention_in_days StorageAccount#change_feed_retention_in_days}.
	ChangeFeedRetentionInDays *float64 `field:"optional" json:"changeFeedRetentionInDays" yaml:"changeFeedRetentionInDays"`
	// container_delete_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#container_delete_retention_policy StorageAccount#container_delete_retention_policy}
	ContainerDeleteRetentionPolicy *StorageAccountBlobPropertiesContainerDeleteRetentionPolicy `field:"optional" json:"containerDeleteRetentionPolicy" yaml:"containerDeleteRetentionPolicy"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#default_service_version StorageAccount#default_service_version}.
	DefaultServiceVersion *string `field:"optional" json:"defaultServiceVersion" yaml:"defaultServiceVersion"`
	// delete_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#delete_retention_policy StorageAccount#delete_retention_policy}
	DeleteRetentionPolicy *StorageAccountBlobPropertiesDeleteRetentionPolicy `field:"optional" json:"deleteRetentionPolicy" yaml:"deleteRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#last_access_time_enabled StorageAccount#last_access_time_enabled}.
	LastAccessTimeEnabled interface{} `field:"optional" json:"lastAccessTimeEnabled" yaml:"lastAccessTimeEnabled"`
	// restore_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#restore_policy StorageAccount#restore_policy}
	RestorePolicy *StorageAccountBlobPropertiesRestorePolicy `field:"optional" json:"restorePolicy" yaml:"restorePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account#versioning_enabled StorageAccount#versioning_enabled}.
	VersioningEnabled interface{} `field:"optional" json:"versioningEnabled" yaml:"versioningEnabled"`
}

