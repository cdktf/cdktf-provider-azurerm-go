// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountImmutabilityPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/storage_account#allow_protected_append_writes StorageAccount#allow_protected_append_writes}.
	AllowProtectedAppendWrites interface{} `field:"required" json:"allowProtectedAppendWrites" yaml:"allowProtectedAppendWrites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/storage_account#period_since_creation_in_days StorageAccount#period_since_creation_in_days}.
	PeriodSinceCreationInDays *float64 `field:"required" json:"periodSinceCreationInDays" yaml:"periodSinceCreationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/storage_account#state StorageAccount#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

