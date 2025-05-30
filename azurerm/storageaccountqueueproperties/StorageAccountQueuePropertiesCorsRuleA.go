// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountqueueproperties


type StorageAccountQueuePropertiesCorsRuleA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/storage_account_queue_properties#allowed_headers StorageAccountQueuePropertiesA#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/storage_account_queue_properties#allowed_methods StorageAccountQueuePropertiesA#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/storage_account_queue_properties#allowed_origins StorageAccountQueuePropertiesA#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/storage_account_queue_properties#exposed_headers StorageAccountQueuePropertiesA#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/storage_account_queue_properties#max_age_in_seconds StorageAccountQueuePropertiesA#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

