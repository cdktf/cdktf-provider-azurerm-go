// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountQueuePropertiesLogging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account#delete StorageAccount#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account#read StorageAccount#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account#version StorageAccount#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account#write StorageAccount#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/storage_account#retention_policy_days StorageAccount#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

