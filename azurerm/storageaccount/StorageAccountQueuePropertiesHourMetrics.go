// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountQueuePropertiesHourMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_account#enabled StorageAccount#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_account#version StorageAccount#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_account#include_apis StorageAccount#include_apis}.
	IncludeApis interface{} `field:"optional" json:"includeApis" yaml:"includeApis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/storage_account#retention_policy_days StorageAccount#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

