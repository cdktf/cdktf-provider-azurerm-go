// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagementpolicy


type StorageManagementPolicyRuleActionsBaseBlob struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#auto_tier_to_hot_from_cool_enabled StorageManagementPolicy#auto_tier_to_hot_from_cool_enabled}.
	AutoTierToHotFromCoolEnabled interface{} `field:"optional" json:"autoTierToHotFromCoolEnabled" yaml:"autoTierToHotFromCoolEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#delete_after_days_since_creation_greater_than StorageManagementPolicy#delete_after_days_since_creation_greater_than}.
	DeleteAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceCreationGreaterThan" yaml:"deleteAfterDaysSinceCreationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#delete_after_days_since_last_access_time_greater_than StorageManagementPolicy#delete_after_days_since_last_access_time_greater_than}.
	DeleteAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceLastAccessTimeGreaterThan" yaml:"deleteAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#delete_after_days_since_modification_greater_than StorageManagementPolicy#delete_after_days_since_modification_greater_than}.
	DeleteAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceModificationGreaterThan" yaml:"deleteAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_creation_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_creation_greater_than}.
	TierToArchiveAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceCreationGreaterThan" yaml:"tierToArchiveAfterDaysSinceCreationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_access_time_greater_than}.
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_last_tier_change_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_tier_change_greater_than}.
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_modification_greater_than}.
	TierToArchiveAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceModificationGreaterThan" yaml:"tierToArchiveAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cold_after_days_since_creation_greater_than StorageManagementPolicy#tier_to_cold_after_days_since_creation_greater_than}.
	TierToColdAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"tierToColdAfterDaysSinceCreationGreaterThan" yaml:"tierToColdAfterDaysSinceCreationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cold_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_cold_after_days_since_last_access_time_greater_than}.
	TierToColdAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToColdAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToColdAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cold_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_cold_after_days_since_modification_greater_than}.
	TierToColdAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToColdAfterDaysSinceModificationGreaterThan" yaml:"tierToColdAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cool_after_days_since_creation_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_creation_greater_than}.
	TierToCoolAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceCreationGreaterThan" yaml:"tierToCoolAfterDaysSinceCreationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cool_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_last_access_time_greater_than}.
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/storage_management_policy#tier_to_cool_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_modification_greater_than}.
	TierToCoolAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceModificationGreaterThan" yaml:"tierToCoolAfterDaysSinceModificationGreaterThan"`
}

