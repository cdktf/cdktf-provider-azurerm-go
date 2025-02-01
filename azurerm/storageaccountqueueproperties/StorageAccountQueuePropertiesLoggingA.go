// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountqueueproperties


type StorageAccountQueuePropertiesLoggingA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account_queue_properties#delete StorageAccountQueuePropertiesA#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account_queue_properties#read StorageAccountQueuePropertiesA#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account_queue_properties#version StorageAccountQueuePropertiesA#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account_queue_properties#write StorageAccountQueuePropertiesA#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/storage_account_queue_properties#retention_policy_days StorageAccountQueuePropertiesA#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

