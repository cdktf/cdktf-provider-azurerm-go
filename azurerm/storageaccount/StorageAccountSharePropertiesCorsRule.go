// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountSharePropertiesCorsRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account#allowed_headers StorageAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account#allowed_methods StorageAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account#allowed_origins StorageAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account#exposed_headers StorageAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/storage_account#max_age_in_seconds StorageAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

